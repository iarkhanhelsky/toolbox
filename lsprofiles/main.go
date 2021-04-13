package main

import (
	"bytes"
	"fmt"
	"github.com/groob/plist"
	"go.mozilla.org/pkcs7"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type entitlements struct {
	ApsEnv string `plist:"aps-environment"`
	ApplicationIdentifier string `plist:"application-identifier"`
}

type profileInfo struct {
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

func (receiver profileInfo) appId() string {
	// Remove '%team_id%.'
	return receiver.Entitlements.ApplicationIdentifier[(len(receiver.TeamIdentifier[0]) + 1):]
}

func readProvisioningProfile(path string) (profileInfo, error) {
	pkey, _ := ioutil.ReadFile(path)
	obj, _ := pkcs7.Parse(pkey)
	var info profileInfo
	err := plist.NewXMLDecoder(bytes.NewReader(obj.Content)).Decode(&info)
	info.PlistData = string(obj.Content)
	info.FilePath = path
	return info, err
}

func printProvisioningProfilesTable(infos []profileInfo) {
	sort.Slice(infos, func(i, j int) bool {
		return strings.Compare(infos[i].Entitlements.ApplicationIdentifier, infos[j].Entitlements.ApplicationIdentifier) < 0
	})
	for _, i := range(infos) {
		printProvisioningProfilesRow(i)
	}
}

func printProvisioningProfilesRow(info profileInfo) {
	fmt.Printf("%-40s %10s %32s %15s\n", info.appId(), info.CreationDate.Format("2006-01-02"), info.UUID, info.Entitlements.ApsEnv)
}

func printProvisioningProfilesDetails(info profileInfo) {

}

func printHelp() {
	println("lsprofiles [PATH...]              print provisioning profiles table")
	println("lsprofiles PROVISIONING_PROFILE   print details about provisioning profile")
	println("lsprofiles --help                 show this message")
	println("Examples:")
}

func parseArgs() []string {
	var paths []string
	var args []string

	if len(os.Args) < 2 {
		userHome, _ := os.UserHomeDir()
		defaultProvisioningDir := filepath.Join(userHome, "Library/MobileDevice/Provisioning Profiles")
		args = []string{defaultProvisioningDir}
	} else {
		args = os.Args[1:]
	}

	if args[0] == "--help"{
		printHelp()
		os.Exit(1)
	}

	for _, p := range args {
		stat, err := os.Stat(p)
		if os.IsNotExist(err) {
			fmt.Printf("error: %s", err.Error())
			os.Exit(2)
		}

		if stat.IsDir() {
			files, err := ioutil.ReadDir(p)
			if err != nil {
				fmt.Printf("error: %s", err.Error())
			}

			for _, f := range(files) {
				if !f.IsDir() {
					paths = append(paths, filepath.Join(p, f.Name()))
				}
			}
		} else {
			paths = append(paths, p)
		}
	}

	return paths
}

func main() {
	paths := parseArgs()
	if len(paths) == 0 {
		printHelp()
		os.Exit(1)
	} else if len(paths) == 1 {
		info, err := readProvisioningProfile(paths[0])
		if err == nil {
			printProvisioningProfilesDetails(info)
		} else {
			fmt.Printf("error: %s", err.Error())
			os.Exit(2)
		}
	} else {
		var infos []profileInfo
		for _, p := range paths {
			info, err := readProvisioningProfile(p)
			if err == nil {
				infos = append(infos, info)
			} else {
				fmt.Printf("error: %s", err.Error())
				os.Exit(2)
			}
		}

		printProvisioningProfilesTable(infos)
	}
}
