package main

import (
	"bytes"
	"github.com/groob/plist"
	"go.mozilla.org/pkcs7"
	"io/ioutil"
	"time"
)

type entitlements struct {
	ApsEnv string `plist:"aps-environment"`
	ApplicationIdentifier string `plist:"application-identifier"`
}

type provisioningProfile struct {
	AppIDName string `plist:"AppIDName"`
	CreationDate time.Time `plist:"CreationDate"`
	Entitlements *entitlements `plist:"Entitlements"`
	Name string `plist:"Name"`
	Platform []string `plist:"Platform"`
	TeamIdentifier []string `plist:"TeamIdentifier"`
	TeamName string `plist:"TeamName"`
	UUID string `plist:"UUID"`

	PlistData string
	FilePath string
}

func (receiver provisioningProfile) appId() string {
	// Remove '%team_id%.'
	return receiver.Entitlements.ApplicationIdentifier[(len(receiver.TeamIdentifier[0]) + 1):]
}

func readProvisioningProfile(path string) (provisioningProfile, error) {
	pkey, _ := ioutil.ReadFile(path)
	obj, _ := pkcs7.Parse(pkey)
	var info provisioningProfile
	err := plist.NewXMLDecoder(bytes.NewReader(obj.Content)).Decode(&info)
	info.PlistData = string(obj.Content)
	info.FilePath = path
	return info, err
}

