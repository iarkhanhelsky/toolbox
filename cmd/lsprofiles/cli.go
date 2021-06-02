package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type cliArgs struct {
	UUIDFilter   string
	AppIDFilter  string
	TeamIDFilter string
	NameFilter   string
	DateFilter   string

	Path         string
	PrintPlist   bool
	PrintDetails bool

	ShowVersion  bool
}

func parseCLI() cliArgs  {
	var args cliArgs

	userHome, _ := os.UserHomeDir()
	defaultProvisioningDir := filepath.Join(userHome, "Library/MobileDevice/Provisioning Profiles")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "NAME\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\tlsprofiles - list installed provision profiles\n")
		fmt.Fprintf(flag.CommandLine.Output(), "SYNOPSIS\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\tlsprofiles [OPTION]... [PATH]\n")
		fmt.Fprintf(flag.CommandLine.Output(), "DESCRIPTION\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\tList information about installed provision profiles ($HOME/Library/MobileDevice/Provisioning Profiles by default)\n")
		flag.PrintDefaults()
	}

	flag.StringVar(&args.UUIDFilter, "uuid-filter", "", "Filter by UUID")
	flag.StringVar(&args.UUIDFilter, "U", "", "Filter by UUID")
	flag.StringVar(&args.AppIDFilter, "appid-filter", "", "Filter by Application ID")
	flag.StringVar(&args.AppIDFilter, "A", "", "Filter by Application ID")
	flag.StringVar(&args.TeamIDFilter, "teamid-filter", "", "Filter by Team ID")
	flag.StringVar(&args.TeamIDFilter, "T", "", "Filter by Team ID")
	flag.StringVar(&args.NameFilter, "name-filter", "", "Filter by Name")
	flag.StringVar(&args.NameFilter, "N", "", "Filter by Name")
	flag.StringVar(&args.DateFilter, "D", "", "Filter by Date")
	flag.StringVar(&args.DateFilter, "date-filter", "", "Filter by Date")

	flag.BoolVar(&args.PrintDetails, "print-details", false, "Print full information for each profile")
	flag.BoolVar(&args.PrintDetails, "d", false, "Print full information for each profile")
	flag.BoolVar(&args.PrintPlist, "print-plist", false, "Print provisioning profile plist")
	flag.BoolVar(&args.PrintPlist, "p", false, "Print provisioning profile plist")
	flag.BoolVar(&args.ShowVersion, "v", false, "Show version and exit")

	flag.Parse()

	if flag.NArg() > 0 {
		args.Path = flag.Arg(0)
		if _, err := os.Stat(args.Path); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "error: %s does not exist\n", args.Path)
			os.Exit(2)
		}
	} else {
		args.Path = defaultProvisioningDir
	}

	if args.ShowVersion {
		println("v0.0.0")
		os.Exit(0)
	}

	return args
}

func (receiver cliArgs) filter() Filter {
	var filters []Filter

	if receiver.UUIDFilter != "" {
		filter := StringContainsFilter{value: receiver.UUIDFilter,
			extractFunc: func (p provisioningProfile) string { return p.UUID }}
		filters = append(filters, filter)
	}

	if receiver.AppIDFilter != "" {
		filter := StringContainsFilter{value: receiver.AppIDFilter,
			extractFunc: func (p provisioningProfile) string { return p.appId() }}
		filters = append(filters, filter)
	}

	if receiver.TeamIDFilter != "" {
		filter :=  StringContainsFilter{value: receiver.TeamIDFilter,
			extractFunc: func (p provisioningProfile) string { return p.TeamIdentifier[0] }}
		filters = append(filters, filter)
	}

	if receiver.NameFilter != "" {
		filter := StringContainsFilter{value: receiver.NameFilter,
			extractFunc: func (p provisioningProfile) string { return p.Name }}
		filters = append(filters, filter)
	}

	if receiver.DateFilter != "" {
		filter := StringEqualsFilter{value: receiver.DateFilter,
			extractFunc: func(p provisioningProfile) string {
				return p.CreationDate.Format("2006-01-02")
			}}

		filters = append(filters, filter)
	}

	return CompoundFilter{filters: filters}
}