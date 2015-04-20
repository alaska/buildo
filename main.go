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
		build targets for platforms

DESCRIPTION
	tbd

FLAGS
	--list
		list available targets and os/arch combinations
	--output=<location/file>
		location to place the built binaries / compressed file
		DEFAULT: ./<current_dir>_bin or ./<current_dir>_bin.gz
	--gz (default: false)
		whether or not to make a .gz out of built binaries
	--debug (default: false)
		output debug information

EXAMPLES
(Assuming $PWD of /home/user/go/src/myproj)

	$ buildo ALL
		builds all binaries for current platform and places them in myproj/myproj_bin
	$ buildo foo
		builds 'foo' binary for current platform and places it into myproj_bin
	$ buildo ALL linux/amd64,freebsd
		builds all binaries for
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
	flag.BoolVar(&toZip, "gz", false, "")
	flag.BoolVar(&debugOut, "debug", false, "")
	flag.StringVar(&outputLoc, "output", "", "")
	flag.Parse()

	targets := getTargets()
	architectures := getArchitectures()

	if doList {
		if flag.NFlag() > 1 {
			fatal("--list can't be combined with other flags")
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
