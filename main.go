package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

func longUsage() {
	fmt.Println(`NAME
	buildo -- build your go binaries

USAGE
	buildo [flags] <ALL|target,...> [ALL|platform,...]

DESCRIPTION
	tbd

FLAGS
	--list
		list available targets and os/arch combinations
	--output=<location/file>
		location to place the built binaries / compressed file
		default: ./$PWD_bin | ./$PWD_bin.tgz
	--tgz
		whether or not to make a gzipped tar out of built binaries
		default: false
	--debug 
		output debug information
		default: false
	--help
		display help info

EXAMPLES
(Assuming $PWD of /home/user/go/src/myproj)
	$ buildo
		builds the default binary for the current platform and places it in
		myproj/myproj_bin
	$ buildo ALL
		builds all binaries for current platform and places them in myproj/myproj_bin
	$ buildo --gz foo,bar windows_386,linux_386
		builds foo and bar binaries and gzips them into myproj/myproj_bin.tgz
`)
	os.Exit(1)
}

var outputLoc string
var doList bool
var toZip bool
var debugOut bool
var cwd, _ = os.Getwd()

func main() {
	flag.Usage = longUsage
	flag.BoolVar(&doList, "list", false, "")
	flag.BoolVar(&toZip, "tgz", false, "")
	flag.BoolVar(&debugOut, "debug", false, "")
	flag.StringVar(&outputLoc, "output", "", "")
	flag.Parse()

	targets := getTargets()
	architectures := getArchitectures()

	if doList {
		if flag.NFlag() >= 2 && !debugOut {
			fatal("--list can only be combined with --debug")
			longUsage()
		}

		sortedTargets := make([]string, 0, len(targets))
		for targetName, _ := range targets {
			//TODO: Default check
			sortedTargets = append(sortedTargets, targetName)
		}
		sort.Strings(sortedTargets)
		fmt.Println("Available targets:")
		for _, targetName := range sortedTargets {
			fmt.Printf("\t%s", targetName)
			if targets[targetName] == cwd || len(targets) == 1 {
				fmt.Println(" *")
			} else {
				fmt.Println()
			}
		}

		sortedArchitectures := make([]string, 0, len(architectures))
		for archName, _ := range architectures {
			sortedArchitectures = append(sortedArchitectures, archName)
		}
		sort.Strings(sortedArchitectures)
		fmt.Println("Available architectures:")
		for _, archName := range sortedArchitectures {
			fmt.Printf("\t%s", archName)
			if architectures[archName].Native {
				fmt.Println(" *")
			} else {
				fmt.Println("")
			}
		}
	}
}
