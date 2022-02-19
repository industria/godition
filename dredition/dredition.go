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
