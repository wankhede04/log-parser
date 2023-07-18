package testparsing

import (
	"testing"

	cmdargs "github.com/m/v2/cmd-args"
	parselog "github.com/m/v2/parse-log"
)

func TestParsing(t *testing.T) {
	fileInfo := cmdargs.FileInfo{
		LogFile:    "./log.log",
		OutputFile: "./output/kill-means-info.json",
	}

	parselog.ParseLogs(fileInfo)
}
