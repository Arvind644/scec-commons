package model

import (
	"encoding/json"
	"testing"

	"github.com/ortelius/scec-commons/database"
	"github.com/stretchr/testify/assert"
)

func TestPackages(t *testing.T) {

	jsonObj := []byte(`{
		"objtype": "Packages",
		"packages": [{
				"objtype": "Package",
				"purl": "pkg:deb/debian/libc-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8",
				"name": "libc-bin",
				"version": "2.19.18+deb8u7",
				"license_key": 23,
				"license": "GP-2.0"
			},
			{
				"objtype": "Package",
				"purl": "pkg:deb/debian/libcpp-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8",
				"name": "libcpp-bin",
				"version": "2.19.18+deb8u7",
				"license_key": 23,
				"license": "GP-2.0"
			}
		]
	}`)

	expected := "{\"objtype\":\"Packages\",\"packages\":[{\"license\":\"GP-2.0\",\"name\":\"libc-bin\",\"objtype\":\"Package\",\"purl\":\"pkg:deb/debian/libc-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8\",\"version\":\"2.19.18+deb8u7\"},{\"license\":\"GP-2.0\",\"name\":\"libcpp-bin\",\"objtype\":\"Package\",\"purl\":\"pkg:deb/debian/libcpp-bin@2.19-18+deb8u7?arch=amd64&upstream=glibc&distro=debian-8\",\"version\":\"2.19.18+deb8u7\"}]}"
	expectedCid := "bafkreicrzo2siragfgs2z6a7ptsua4uzl5qwsi6ulfr5rszyafutmkf4hm"

	// define user object to marshal into
	obj := NewPackages()

	// convert json string into the user object
	json.Unmarshal(jsonObj, obj)

	// create all cids for the json string
	cid, _ := database.MakeNFT(obj)
	// 	fmt.Println(cid)
	assert.Equal(t, expectedCid, cid, "check persisted cid with test cid")

	// convert all the cids back to json string
	jsonStr, _ := database.MakeJSON(cid)
	assert.Equal(t, expected, jsonStr, "check persisted cid json with test json string")

}
