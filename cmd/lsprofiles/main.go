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

		for _, f := range files {
			if !f.IsDir() && strings.HasSuffix(f.Name(), ".mobileprovision") {
				paths = append(paths, filepath.Join(path, f.Name()))
			}
		}
	} else {
		paths = append(paths, path)
	}
	return paths
}

func readAll(paths []string) []provisioningProfile {
	var provisions []provisioningProfile
	results := make(chan provisioningProfile)
	errors := make(chan error)
	defer close(results)
	defer close(errors)

	errorsCount := 0

	for _, f := range paths {
		f := f
		go func() {
			p, err := readProvisioningProfile(f)
			if err != nil {
				errors <- err
			} else {
				results <- p
			}
		}()
	}

	for len(provisions) + errorsCount < len(paths) {
		select {
		case r := <- results:
			provisions = append(provisions, r)
		case err := <- errors:
			errorsCount += 1
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		}
	}

	return provisions
}

func filter(profiles []provisioningProfile, filter Filter) []provisioningProfile {
	var filtered []provisioningProfile
	for _, p := range profiles {
		if filter.match(p) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

func main() {
	args := parseCLI()
	profiles := readAll(scanProfiles(args.Path))

	profiles = filter(profiles, args.filter())
	if args.PrintPlist || args.PrintDetails {
		for _, profile := range profiles {
			fmt.Printf("%s\n", profile.FilePath)
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
