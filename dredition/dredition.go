package dredition // github.com/industria/godition/dredition"

import (
	"encoding/json"
	"io"
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

// Get Metadata using
// https://sphynx.aptoma.no/burned/{editionId}
// https://sphynx.aptoma.no/burned/5d5a8cf857cd2009c74b6378
