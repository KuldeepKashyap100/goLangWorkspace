package interaction

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
)

type Round struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func PrintGreetings() {
	asciiFigure := figure.NewFigure("MONSTER SLAYER", "", true)
	asciiFigure.Print()
	fmt.Println("Starting a new Game")
	fmt.Println("Good luck!")
}

func ShowAvailableActions(isSpecialAttackAvailable bool) {
	fmt.Println("Please choose your action")
	fmt.Println("--------------------------")
	fmt.Println("1) Attack Monster")
	fmt.Println("2) Heal")

	if isSpecialAttackAvailable {
		fmt.Println("3) Special Attack")
	}
}

func PrintRoundStatistics(roundData *Round) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked Monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a strong attack against Monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "HEAL" {
		fmt.Printf("Player Healed for %v.\n", roundData.PlayerHealValue)
	}
	fmt.Printf("Monster attacked Player for %v damage.\n", roundData.MonsterAttackDmg)
	fmt.Printf("Player Health:%v.\n", roundData.PlayerHealth)
	fmt.Printf("Monster Health: %v.\n", roundData.MonsterHealth)
}

func DeclareWinner(winner string) {
	fmt.Println("--------------------------")
	asciiFigure := figure.NewColorFigure("Game Over!", "", "red", true)
	asciiFigure.Print()
	fmt.Println("--------------------------")
	fmt.Printf("%v won!\n", winner)
}

func WriteLogFile(roundData *[]Round) {
	exPath, err := os.Executable()

	if err != nil {
		fmt.Println("Writing log file failed. Exiting...")
		return
	}

	exPath = filepath.Dir(exPath)

	file, err := os.Create(exPath + "/gamelog.txt")
	
	if err != nil {
		fmt.Println("Saving log failed, Exiting...")
		return
	}
	for roundIdx, roundValue := range *roundData {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(roundIdx + 1),
			"Action":                roundValue.Action,
			"Player Attack Damage":  fmt.Sprint(roundValue.PlayerAttackDmg),
			"Player Heal value":     fmt.Sprint(roundValue.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprint(roundValue.PlayerAttackDmg),
			"Player Health":         fmt.Sprint(roundValue.PlayerHealth),
			"Monster Health":        fmt.Sprint(roundValue.MonsterHealth),
		}
		logLine := fmt.Sprintln(logEntry)
		_, err := file.WriteString(logLine)
		if err != nil {
			fmt.Println("Writing into log file failed. Exiting...")
			continue
		}
	}
	file.Close()
	fmt.Println("Wrote data to log!")
}
