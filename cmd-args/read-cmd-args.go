package cmdargs

import (
	"flag"
	"fmt"
	"os"
)

type FileInfo struct {
	LogFile    string
	OutputFile string
}

// read arguments passed while executing to build or code
func ReadCMDArgs() FileInfo {
	var fileInfo = FileInfo{}

	flag.StringVar(&fileInfo.LogFile, "log-file", "", "file to read logs from")
	flag.StringVar(&fileInfo.OutputFile, "report-file", "./output/kill-means-info.json", "file to write report to")
	flag.Parse()

	if fileInfo.LogFile == "" {
		fmt.Println("Please specify a valid log file")
		os.Exit(1)
	}

	return fileInfo
}
