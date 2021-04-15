package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func scanProfiles(path string) []string {
	var paths []string

	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("error: %s", err.Error())
		os.Exit(2)
	}

	if stat.IsDir() {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			fmt.Printf("error: %s", err.Error())
		}

		for _, f := range(files) {
			if !f.IsDir() && strings.HasSuffix(f.Name(), ".mobileprovision") {
				paths = append(paths, filepath.Join(path, f.Name()))
			}
		}
	} else {
		paths = append(paths, path)
	}
	return paths
}

func readAll(paths []string) ([]provisioningProfile, error) {
	var provisions []provisioningProfile
	for _, f := range(paths) {
		i, err := readProvisioningProfile(f)
		if err != nil {
			return nil, err
		}
		provisions = append(provisions, i)
	}

	return provisions, nil
}

func filter(profiles []provisioningProfile, uuidFilter string) []provisioningProfile {
	var filtered []provisioningProfile
	for _, p := range profiles {
		if strings.Contains(p.UUID, uuidFilter) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

func main() {
	args := parseCLI()
	profiles, err := readAll(scanProfiles(args.Path))
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		os.Exit(2)
	}

	profiles = filter(profiles, args.UUIDFilter)
	if args.PrintPlist || args.PrintDetails {
		for _, profile := range profiles {
			println(profile.FilePath)
			if args.PrintDetails {
				printProvisioningProfilesDetails(profile)
			}
			if args.PrintPlist {
				printProvisioningProfilePlist(profile)
			}
		}
	} else {
		printProvisioningProfilesTable(profiles)
	}
}
