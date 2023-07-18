package parselog

import (
	"fmt"
	"sort"
	"sync"
)

// function to report grouped info in console synchronously
// have manually printed instead of using library to convert object to json to maintain synchronicity and provide better formatting to output
func ReportRank(gameData chan GameData, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("{\n")
	isFirst := true
	for game := range gameData {
		if isFirst {
			fmt.Printf("\n\t\"%s\": {\n", game.GameName)
			isFirst = false
		} else {
			fmt.Printf(",\n\t\"%s\": {\n", game.GameName)
		}
		fmt.Printf("\t\t\"total_kills\": %d,\n", game.TotalKills)
		fmt.Printf("\t\t\"kills\": {\n")
		playerCount := 0
		for player, killCount := range game.Kills {
			if len(game.Kills)-1 != playerCount {
				fmt.Printf("\t\t\t\"%s\": %d,\n", player, killCount)
				playerCount++
			} else {
				fmt.Printf("\t\t\t\"%s\": %d\n", player, killCount)
			}
		}
		fmt.Printf("\t\t}\n")
		fmt.Printf("\t\t\"ranks\": {\n")
		PrintRank(game.Kills)
		fmt.Printf("\t\t}\n")
		fmt.Printf("\t}")

	}
	fmt.Printf("\n}\n")
}

func PrintRank(playersKills map[string]int) {
	keys := make([]string, 0, len(playersKills))

	for key := range playersKills {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return playersKills[keys[i]] > playersKills[keys[j]]
	})

	rank := 1

	for index, key := range keys {
		//	print same rank if kills are same
		if index != 0 && playersKills[key] < playersKills[keys[index-1]] {
			rank++
		}

		if index != len(keys)-1 {
			fmt.Printf("\t\t\t\"%s\": %d\n", key, rank)
		} else {
			fmt.Printf("\t\t\t\"%s\": %d\n", key, rank)
		}
	}
}
