// Package ortelius - Vulnerabilities defines the struct and handles marshaling/unmarshaling the struct to/from NFT Storage.
package ortelius

import "encoding/json"

// Vulnerabilities defines a list of Vulnerability
type Vulnerabilities struct {
	Key             string          `json:"_key,omitempty"`
	NftJSON         string          `json:"_json,omitempty"`
	Vulnerabilities []Vulnerability `json:"vulnerabilties,omitempty"`
}

// MarshalNFT converts the struct into a normalized JSON NFT
func (obj *Vulnerabilities) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	type VulnerabilityNFT struct {
		Vulnerabilities []NFT `json:"vulnerabilties,omitempty"`
	}
	var vulnlist VulnerabilityNFT

	for _, v := range obj.Vulnerabilities {
		nft := new(NFT).Init(v.MarshalNFT(cid2json))
		vulnlist.Vulnerabilities = append(vulnlist.Vulnerabilities, nft)
	}

	data, _ := json.Marshal(vulnlist)
	obj.NftJSON = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJSON // Add cid=json for persisting later

	return data
}

// UnmarshalNFT converts the JSON from NFT Storage to a new instance of the struct
func (obj *Vulnerabilities) UnmarshalNFT(cid2json map[string]string) {
	var pkgs Vulnerabilities // define domain object to marshal into
	var exists bool
	var NftJSON string

	// get the json from storage
	if NftJSON, exists = cid2json[obj.Key]; exists {
		obj.NftJSON = NftJSON // Set the nft json for the object
	}

	err := json.Unmarshal([]byte(obj.NftJSON), &pkgs) // Convert the nft json into the domain object

	if err == nil {
		// Deep Copy
		obj.Vulnerabilities = make([]Vulnerability, 0)

		for _, v := range pkgs.Vulnerabilities {
			var rec Vulnerability

			rec.Key = v.Key
			rec.UnmarshalNFT(cid2json)
			obj.Vulnerabilities = append(obj.Vulnerabilities, rec)
		}
	}
}