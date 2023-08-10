// Package model - Swagger defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package model

import "encoding/json"

// Swagger defines an OpenAPI or Swagger file
type Swagger struct {
	Key     string          `json:"_key,omitempty"`
	Content json.RawMessage `json:"content"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *Swagger) MarshalNFT(cid2json map[string]string) string {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Content json.RawMessage `json:"content"`
		ObjType string          `json:"objtype"`
	}{
		Content: obj.Content,
		ObjType: "Swagger",
	})

	obj.Key = new(NFT).Init(string(data)).Key
	cid2json[obj.Key] = string(data) // Add cid=json for persisting later

	return string(data)
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *Swagger) UnmarshalNFT(cid2json map[string]string) {
	var swagger Swagger // define domain object to marshal into
	var exists bool
	var nftJSON string

	// get the json from storage
	if nftJSON, exists = cid2json[obj.Key]; exists {

		err := json.Unmarshal([]byte(nftJSON), &swagger) // Convert the nft json into the domain object

		if err == nil {
			// Deep Copy
			obj.Content = swagger.Content
		}
	}
}
