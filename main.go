package main

import (
	cmdargs "github.com/m/v2/cmd-args"
	parselog "github.com/m/v2/parse-log"
)

func main() {
	cmdArgs := cmdargs.ReadCMDArgs()
	parselog.ParseLogs(cmdArgs)
}
