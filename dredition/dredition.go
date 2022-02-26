package dredition // github.com/industria/godition/dredition"

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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

type BurnProcessor struct {
	httpClient *http.Client
}

func NewBurnProcessor() *BurnProcessor {
	return &BurnProcessor{
		httpClient: &http.Client{Timeout: time.Second * 10},
	}

}

func (bp *BurnProcessor) Process(n Notification) error {
	m, err := bp.metadata(n)
	if err != nil {
		return err
	}
	log.Printf("Metadata : %v", m)
	return nil
}

func (bp *BurnProcessor) metadata(n Notification) (*BurnMetadata, error) {
	url := "https://sphynx.aptoma.no/burned/" + n.Data.Edition.Id
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := bp.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	metadata, err := ReadBurnMetadata(resp.Body)
	//var metadata BurnMetadata
	//err = json.NewDecoder(resp.Body).Decode(&metadata)
	//if err != nil {
	//	return nil, err
	//}
	return metadata, err
}

// Get Metadata using
// https://sphynx.aptoma.no/burned/{editionId}
// https://sphynx.aptoma.no/burned/5d5a8cf857cd2009c74b6378
