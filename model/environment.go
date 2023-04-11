package model

import (
	"encoding/json"
	"time"
)

type Environment struct {
	Key     string    `json:"_key,omitempty"`
	NftJson string    `json:"_json,omitempty"`
	Created time.Time `json:"created"`
	Creator User      `json:"creator"`
	Domain  Domain    `json:"domain"`
	Name    string    `json:"name"`
	Owner   User      `json:"owner"`
}

func (obj *Environment) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Created time.Time `json:"created"`
		Creator NFT       `json:"creator"`
		Domain  NFT       `json:"domain"`
		Name    string    `json:"name"`
		ObjType string    `json:"objtype"`
		Owner   NFT       `json:"owner"`
	}{
		Created: obj.Created,
		Creator: new(NFT).Init(obj.Creator.MarshalNFT(cid2json)),
		Domain:  new(NFT).Init(obj.Domain.MarshalNFT(cid2json)),
		Name:    obj.Name,
		ObjType: "Environment",
		Owner:   new(NFT).Init(obj.Owner.MarshalNFT(cid2json)),
	})

	obj.NftJson = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJson // Add cid=json for persisting later

	return data
}

func (obj *Environment) UnmarshalNFT(cid2json map[string]string) {
	var environment Environment
	var exists bool
	var NftJson string

	// get the json from storage
	if NftJson, exists = cid2json[obj.Key]; exists {
		obj.NftJson = NftJson // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJson), &environment)

	// Deep Copy
	obj.Created = environment.Created
	obj.Creator.Key = environment.Creator.Key
	obj.Creator.UnmarshalNFT(cid2json)
	obj.Domain.Key = environment.Domain.Key
	obj.Domain.UnmarshalNFT(cid2json)
	obj.Name = environment.Name
	obj.Owner.Key = environment.Owner.Key
	obj.Owner.UnmarshalNFT(cid2json)
}
