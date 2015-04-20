# buildo
```
$ ./buildo --help
NAME
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
```
