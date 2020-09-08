package cmd

import "encoding/json"

type RequestBody struct {
	Query   string `json:"gremlin"`
	Profile string `json:"profile.serializer"`
}

func NewRequestBody(query string) ([]byte, error) {
	body := &RequestBody{
		Query:   query,
		Profile: "application/vnd.gremlin-v3.0+gryo",
	}
	return json.Marshal(body)
}
