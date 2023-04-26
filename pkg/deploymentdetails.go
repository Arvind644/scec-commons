// Package pkg - DeploymentDetails defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package pkg

import "encoding/json"

// DeploymentDetails defines a deployment plus the associated log
type DeploymentDetails struct {
	Key        string     `json:"_key,omitempty"`
	NftJSON    string     `json:"_json,omitempty"`
	Deployment Deployment `json:"deployment"`
	Log        []string   `json:"log,omitempty"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *DeploymentDetails) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	data, _ := json.Marshal(&struct {
		Deployment NFT      `json:"deployment"`
		Log        []string `json:"log,omitempty"`
		ObjType    string   `json:"objtype"`
	}{
		Deployment: new(NFT).Init(obj.Deployment.MarshalNFT(cid2json)),
		Log:        obj.Log,
		ObjType:    "DeploymentDetails",
	})

	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *DeploymentDetails) UnmarshalNFT(cid2json map[string]string) {
	var deploydetails DeploymentDetails
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	err := json.Unmarshal([]byte(obj.NftJSON), &deploydetails)

	if err == nil {
		// Deep Copy
		obj.Deployment.Key = deploydetails.Deployment.Key
		obj.Deployment.UnmarshalNFT(cid2json)

		obj.Log = append(obj.Log, deploydetails.Log...)
	}
}
