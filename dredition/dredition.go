package dredition // github.com/industria/godition/dredition"

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Product struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"productType"`
}

type Edition struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Data struct {
	Product Product `json:"product"`
	Edition Edition `json:"edition"`
}

type Notification struct {
	Data  Data   `json:"data"`
	Event string `json:"event"`
}

func ReadNotification(r io.Reader) (*Notification, error) {
	var notification Notification
	err := json.NewDecoder(r).Decode(&notification)
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

type BurnMetadata struct {
	ClientID      string `json:"clientId"`
	EditionID     string `json:"editionId"`
	HTMLHash      string `json:"htmlHash"`
	HTMLUpdatedAt string `json:"htmlUpdatedAt"`
	HTMLUrl       string `json:"htmlUrl"`
	CSSHash       string `json:"cssHash"`
	CSSUpdatedAt  string `json:"cssUpdatedAt"`
	CSSUrl        string `json:"cssUrl"`
	PreviewUrl    string `json:"previewUrl"`
}

func ReadBurnMetadata(r io.Reader) (*BurnMetadata, error) {
	var meta BurnMetadata
	err := json.NewDecoder(r).Decode(&meta)
	if err != nil {
		return nil, err
	}
	return &meta, nil
}

// DrEdition client for accessing services
type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{Timeout: time.Second * 10},
	}
}

func (c *Client) CSS(m *BurnMetadata) (*string, error) {
	return c.Resource(m.CSSUrl)
}

func (c *Client) HTML(m *BurnMetadata) (*string, error) {
	return c.Resource(m.HTMLUrl)
}

func (c *Client) Resource(url string) (*string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request from %s : %v", url, err)
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed request of %s : %v", url, err)
	}
	defer res.Body.Close()

	builder := &strings.Builder{}
	_, err = io.Copy(builder, res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read from %s : %v", url, err)
	}
	s := builder.String()
	return &s, nil
}

func (c *Client) Metadata(n Notification) (*BurnMetadata, error) {
	url := "https://sphynx.aptoma.no/burned/" + n.Data.Edition.Id
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ReadBurnMetadata(resp.Body)
}

// Get Metadata using
// https://sphynx.aptoma.no/burned/{editionId}
// https://sphynx.aptoma.no/burned/5d5a8cf857cd2009c74b6378
