package main

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

//var mainFinder = regexp.MustCompile(`^\s*package\s+main`)
var mainFinder = regexp.MustCompile(`^(/\*([^*]|[\r\n]|(\*+([^*/]|[\r\n])))*\*+/\n)*|(//.*\n)*package\s+main`)

func getTargets() map[string]string {
	targets := map[string]string{}
	processed := map[string]bool{}
	toResolve := map[string][]string{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if string(path[0]) == "." {
			return filepath.SkipDir
		}
		if !info.IsDir() {
			// We're looking at a file somewhere

			fileDir := filepath.Dir(path)
			if processed[fileDir] {
				// only need to process one .go file per directory
				return nil
			}
			if filepath.Ext(path) == ".go" {
				// We've found the first .go file in this directory

				// check for "package main"
				fh, err := os.Open(path)
				if err != nil {
					fatal("error opening source file:", err)
				}
				defer fh.Close()
				buf := bufio.NewReader(fh)

				if !mainFinder.MatchReader(buf) {
					// If the first .go file we find doesn't have "package main"
					// at the beginning, we're not interested in this directory
					processed[fileDir] = true
					return nil
				}

				// Found package main, this is a binary
				targetName := filepath.Base(fileDir)

				// See if there is a conflicting target name
				if _, found := targets[targetName]; found {
					debug("Conflicting target name", targetName, "found")
					// Is this the first conflict? Start the queue with the
					// original if so.
					if _, found := toResolve[targetName]; !found {
						toResolve[targetName] = []string{targets[targetName]}
					}
					// Add the conflicting target
					toResolve[targetName] = append(toResolve[targetName], fileDir)

				} else {
					// Add the target hame
					targets[targetName] = fileDir
				}
				processed[fileDir] = true
			}
		}
		return nil
	}
	err := filepath.Walk(cwd, walkFn)
	if err != nil {
		fatal("Error scanning directory:", err)
	}
	// resolve any conflicting target names
	for targetName, paths := range toResolve {
		// remove the conflicting entry from the targets list
		delete(targets, targetName)
		for _, path := range paths {
			newTargetName := strings.Replace(path, cwd, "", 1)
			newTargetName = strings.Replace(newTargetName, string(os.PathSeparator), "", 1)
			newTargetName = strings.Replace(newTargetName, string(os.PathSeparator), "_", -1)
			targets[newTargetName] = path
		}
	}
	return targets
}
