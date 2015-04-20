package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

type buildArch struct {
	OS     string
	Arch   string
	Native bool
}

func getArchitectures() map[string]buildArch {
	a := map[string]buildArch{}

	// Add native environ
	a[fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH)] = buildArch{runtime.GOOS, runtime.GOARCH, true}

	otherArchs, err := ioutil.ReadDir(filepath.Join(runtime.GOROOT(), "bin"))
	if err != nil {
		fatal("Can't retrieve other architectures", err)
	}

	for _, file := range otherArchs {
		if file.IsDir() {
			//split that sucka
			a[file.Name()] = buildArch{OS: "", Arch: ""}
		}
	}

	return a
}
