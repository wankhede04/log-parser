package testparsing

import (
	"sync"
	"testing"

	cmdargs "github.com/m/v2/cmd-args"
	parselog "github.com/m/v2/parse-log"
)

func TestPrintingGameReport(t *testing.T) {
	gameData := make(chan parselog.GameData)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go parselog.ReportRank(gameData, &wg)
	sampleGame := parselog.GameData{
		GameName:   "TestGame",
		TotalKills: 4,
		Players:    map[string]bool{"player1": true, "player2": true},
		Kills:      map[string]int{"player1": 2, "player2": 2},
	}
	gameData <- sampleGame
	close(gameData)
	wg.Wait()
}

func TestPrintingKillReport(t *testing.T) {
	fileInfo := cmdargs.FileInfo{
		LogFile:    "./log.log",
		OutputFile: "./output/kill-means-info.json",
	}
	killData := make(chan parselog.KillData)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go parselog.GenerateKillReport(killData, fileInfo, &wg)
	sampleKillData := parselog.KillData{
		GameName: "TestGame",
		Kills:    map[string]int{"MOD_GUN_1": 2, "MOD_GUN_2": 2},
	}
	killData <- sampleKillData
	close(killData)
	wg.Wait()
}
