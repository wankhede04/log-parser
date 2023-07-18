package parselog

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	cmdargs "github.com/m/v2/cmd-args"
)

// function to report grouped info to file synchronously
// have manually printed instead of using library to convert object to json to maintain synchronicity and provide better formatting to output
func GenerateKillReport(killData chan KillData, file cmdargs.FileInfo, wg *sync.WaitGroup) {
	defer wg.Done()

	err := os.MkdirAll(filepath.Dir(file.OutputFile), 0770)
	if err != nil {
		fmt.Printf("Unable to create specified directory: %s \n", err)
		os.Exit(1)
	}
	outputFile, err := os.Create(file.OutputFile)
	if err != nil {
		fmt.Printf("Unable to create file: %s \n", err)
		os.Exit(1)
	}
	fmt.Fprintf(outputFile, "{\n")
	count := 0
	for kill := range killData {
		if count == 0 {
			fmt.Fprintf(outputFile, "\t\"%s\": {\n", kill.GameName)
			count++
		} else {
			fmt.Fprintf(outputFile, ",\n\t\"%s\": {\n", kill.GameName)
		}
		fmt.Fprintf(outputFile, "\t\t\"kills_by_means\": {\n")
		weaponCount := 0
		for weapon, killCount := range kill.Kills {
			if len(kill.Kills)-1 != weaponCount {
				fmt.Fprintf(outputFile, "\t\t\t\"%s\": %d,\n", weapon, killCount)
				weaponCount++
			} else {
				fmt.Fprintf(outputFile, "\t\t\t\"%s\": %d\n", weapon, killCount)
			}
		}
		fmt.Fprintf(outputFile, "\t\t}\n")
		fmt.Fprintf(outputFile, "\t}")
	}
	fmt.Fprintf(outputFile, "\n}")
}
