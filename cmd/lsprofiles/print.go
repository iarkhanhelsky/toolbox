package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

func printProvisioningProfilesTable(profiles []provisioningProfile) {
	sort.Slice(profiles, func(i, j int) bool {
		return strings.Compare(profiles[i].Entitlements.ApplicationIdentifier, profiles[j].Entitlements.ApplicationIdentifier) < 0
	})

	fmt.Printf("%-10s %-3s %-10s %-40s %-36s %s\n", "Created", "Env", "Team ID", "App ID", "UUID", "File")
	for _, i := range profiles {
		printProvisioningProfilesRow(i)
	}
}

func printProvisioningProfilesRow(profile provisioningProfile) {
	appId := profile.appId()
	if len(appId) > 40 {
		appId = appId[:36] + "..."
	}

	t := ""
	if profile.Entitlements.ApsEnv == "production" {
		t = "P"
	} else if profile.Entitlements.ApsEnv == "development" {
		t = "D"
	}

	fmt.Printf("%10s %3s %10s %-40s %36s %s\n",
		profile.CreationDate.Format("2006-01-02"), t, profile.TeamIdentifier[0], appId, profile.UUID, filepath.Base(profile.FilePath))
}

func printProvisioningProfilesDetails(profile provisioningProfile) {
	fmt.Printf("Name:       %s\n", profile.Name)
	fmt.Printf("UUID:       %s\n", profile.UUID)
	fmt.Printf("Team ID:    %s\n", profile.TeamIdentifier[0])
	fmt.Printf("App ID:     %s\n", profile.appId())
	fmt.Printf("Created:    %s\n", profile.CreationDate.Format("2006-01-02"))
	fmt.Printf("Platform:   %s\n", profile.Platform[0])
	fmt.Printf("Enviroment: %v\n", profile.Entitlements.ApsEnv)
	fmt.Printf("-----------------------------------------------------------------------------\n")
}

func printProvisioningProfilePlist(profile provisioningProfile) {
	println(profile.PlistData)
}