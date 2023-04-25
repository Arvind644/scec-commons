package model

import "encoding/json"

type Components struct {
	Key        string             `json:"_key,omitempty"`
	NftJSON    string             `json:"_json,omitempty"`
	Components []ComponentVersion `json:"components,omitempty"`
}

func (obj *Components) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	type ComponentsNFT struct {
		Components []NFT `json:"components,omitempty"`
	}
	var complist ComponentsNFT

	for _, v := range obj.Components {
		nft := new(NFT).Init(v.MarshalNFT(cid2json))
		complist.Components = append(complist.Components, nft)
	}

	data, _ := json.Marshal(complist)
	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

func (obj *Components) UnmarshalNFT(cid2json map[string]string) {
	var comps Components // define domain object to marshal into
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJSON), &comps) // Convert the nft json into the domain object

	// Deep Copy
	obj.Components = make([]ComponentVersion, 0)

	for _, v := range comps.Components {
		var rec ComponentVersion

		rec.Key = v.Key
		rec.UnmarshalNFT(cid2json)
		obj.Components = append(obj.Components, rec)
	}
}