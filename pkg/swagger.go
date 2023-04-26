// Package pkg - Swagger defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package pkg

import "encoding/json"

// Swagger defines an OpenAPI or Swagger file
type Swagger struct {
	Key     string   `json:"_key,omitempty"`
	NftJSON string   `json:"_json,omitempty"`
	Content []string `json:"content"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *Swagger) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Content []string `json:"content"`
		ObjType string   `json:"objtype"`
	}{
		Content: obj.Content,
		ObjType: "Swagger",
	})

	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *Swagger) UnmarshalNFT(cid2json map[string]string) {
	var swagger Swagger // define domain object to marshal into
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	err := json.Unmarshal([]byte(obj.NftJSON), &swagger) // Convert the nft json into the domain object

	if err == nil {
		// Deep Copy
		obj.Content = append(obj.Content, swagger.Content...)
	}
}
