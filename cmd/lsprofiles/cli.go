package main

import (
	"flag"
	"os"
	"path/filepath"
)

type cliArgs struct {
	UUIDFilter   string
	Path         string
	PrintPlist   bool
	PrintDetails bool

	ShowVersion bool
}

func parseCLI() cliArgs  {
	var args cliArgs

	userHome, _ := os.UserHomeDir()
	defaultProvisioningDir := filepath.Join(userHome, "Library/MobileDevice/Provisioning Profiles")

	flag.StringVar(&args.Path, "path", defaultProvisioningDir, "Directory path or *.mobileprovision file")
	flag.StringVar(&args.UUIDFilter, "uuid-filter", "", "Filter by UUID")
	flag.StringVar(&args.UUIDFilter, "u", "", "Filter by UUID")
	flag.BoolVar(&args.PrintDetails, "print-details", false, "Print full information for each profile")
	flag.BoolVar(&args.PrintDetails, "d", false, "Print full information for each profile")
	flag.BoolVar(&args.PrintPlist, "print-plist", false, "Print provisioning profile plist")
	flag.BoolVar(&args.PrintPlist, "p", false, "Print provisioning profile plist")
	flag.BoolVar(&args.ShowVersion, "v", false, "Show version and exit")

	flag.Parse()

	if args.ShowVersion {
		println("v0.0.0")
		os.Exit(0)
	}

	return args
}