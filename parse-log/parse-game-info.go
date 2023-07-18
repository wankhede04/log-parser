package parselog

import (
	"bufio"
	"regexp"
	"strings"
)

// function to read info about single game
func ReadGameInfo(scanner *bufio.Scanner, name string) (GameData, KillData) {
	gameData := GameData{
		GameName: name,
	}
	killData := KillData{
		GameName: name,
	}

	// regex pattern to match when parsing kill info, link to visualization can be found in README.md
	var killPattern = regexp.MustCompile(`Kill: (\d+) (\d+) (\d+): (.+?) killed (.+?) by ([^\n]+)`)

	players := make(map[string]bool)
	weaponKills := make(map[string]int)
	playerKills := make(map[string]int)

	for scanner.Scan() {
		//	read line as a string to get info about a kill
		line := scanner.Text()

		//  match pattern to identify whether it is a string containing info on a kill
		if matches := killPattern.FindStringSubmatch(line); matches != nil {
			//  player that killed
			player1 := matches[4]
			//  player that got killed
			player2 := matches[5]
			//  weapon that was used
			weapon := matches[6]

			if player1 == "<world>" {
				//	reduce number of kills of a player if the player gets killed by <world>
				playerKills[player2]--
			} else {
				//	increment number of kills for the player who killed
				playerKills[player1]++
				//	register player that killed
				players[player1] = true
			}
			//	register the player that got killed
			players[player2] = true
			// 	increment kill associated to weapon
			weaponKills[weapon]++
			gameData.TotalKills++
		} else if strings.Contains(line, "--------------------------------------------------------") { //	check if this line contains "----..." string as it symbolizes end of game
			break
		}
	}

	gameData.Kills = playerKills
	killData.Kills = weaponKills
	gameData.Players = players
	return gameData, killData
}
