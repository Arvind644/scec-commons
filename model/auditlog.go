package model

import "encoding/json"

type AuditLog struct {
	Key      string        `json:"_key,omitempty"`
	NftJson  string        `json:"_json,omitempty"`
	AuditLog []AuditRecord `json:"auditlog,omitempty"`
}

func (obj *AuditLog) MarshalNFT(cid2json map[string]string) []byte {

	// Sturct must be manually sorted alphabetically in order for consistent CID to be produced
	type AuditLogNFT struct {
		AuditLog []NFT `json:"auditlog,omitempty"`
	}
	var reclist AuditLogNFT

	for _, v := range obj.AuditLog {
		nft := new(NFT).Init(v.MarshalNFT(cid2json))
		reclist.AuditLog = append(reclist.AuditLog, nft)
	}

	data, _ := json.Marshal(reclist)
	obj.NftJson = string(data)
	obj.Key = new(NFT).Init(data).Key
	cid2json[obj.Key] = obj.NftJson // Add cid=json for persisting later

	return data
}

func (obj *AuditLog) UnmarshalNFT(cid2json map[string]string) {
	var auditlog AuditLog // define domain object to marshal into
	var exists bool
	var NftJson string

	// get the json from storage
	if NftJson, exists = cid2json[obj.Key]; exists {
		obj.NftJson = NftJson // Set the nft json for the object
	}

	json.Unmarshal([]byte(obj.NftJson), &auditlog) // Convert the nft json into the domain object

	// Deep Copy
	obj.AuditLog = make([]AuditRecord, 0)

	for _, v := range auditlog.AuditLog {
		var rec AuditRecord

		rec.Key = v.Key
		rec.UnmarshalNFT(cid2json)
		obj.AuditLog = append(obj.AuditLog, rec)
	}
}
