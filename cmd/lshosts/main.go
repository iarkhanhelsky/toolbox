package main

import (
	"fmt"
	"github.com/kevinburke/ssh_config"
	"os"
	"path/filepath"
)

type hostData struct {
	patterns []string
	hostName string
	userName string
	port string
}

func readHost(host *ssh_config.Host) hostData {
	var data hostData
	for _, p := range host.Patterns {
		data.patterns = append(data.patterns, p.String())
	}

	for _, node := range host.Nodes {
		switch n := node.(type) {
		case *ssh_config.KV:
			if n.Key == "User" {
				data.userName = n.Value
			} else if n.Key == "HostName" {
				data.hostName = n.Value
			} else if n.Key == "Port" {
				data.port = n.Value
			}
		}
	}

	return data
}

func (receiver hostData) String() string {
	var line string
	for i, p := range receiver.patterns {
		line += fmt.Sprintf("%-30s", p)
		if i < len(receiver.patterns) - 1 {
			line += "\n"
		}
	}

	if len(receiver.hostName) > 0 {
		h := receiver.hostName
		if len(receiver.userName) > 0 {
			h = receiver.userName + "@" + h
		}
		line += " " + h
	}

	if len(receiver.port) > 0 {
		line += " -p " + receiver.port
	}

	return line
}

func main() {
	f, _ := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "config"))
	cfg, _ := ssh_config.Decode(f)
	for _, host := range cfg.Hosts {
		fmt.Println(readHost(host).String())
	}
}