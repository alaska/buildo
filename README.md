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
## Brand new
As of 4/20 12:30:00 EST, this is a brand new repository and the program is currently still in development. As of right now, the checkin is just insurance against loss of work. Sometime in the next couple of days I will add the build portion, followed directly by the parallelization of the build. You probably don't even know this project exists yet, because I haven't told anyone, but if you do happen across it, check back very soon :)
