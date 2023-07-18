package testparsing

import (
	"os"
	"testing"

	cmdargs "github.com/m/v2/cmd-args"
)

func TestCMDArgs(t *testing.T) {
	os.Args = append(os.Args, "--log-file=./t")
	cmdargs.ReadCMDArgs()
}
