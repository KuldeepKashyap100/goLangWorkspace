package main

import (
	"github.com/kuldeep/monster-slayer/actions"

	"github.com/kuldeep/monster-slayer/interaction"
)

var currentRound = 0
var gameRounds = []interaction.Round{}

func main() {
	startGame()

	winner := ""
	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreetings()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
	playerChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerAttackDmg int
	var monsterAttackDmg int
	var playerHealValue int

	if playerChoice == "ATTACK" {
		playerAttackDmg = actions.AttackMonster(false)
	} else if playerChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		playerAttackDmg = actions.AttackMonster(true)
	}

	monsterAttackDmg = actions.AttackPlayer()

	playerHealth, monsterHealth := actions.GetHealthAmounts()

	roundData := interaction.Round{
		Action:           playerChoice,
		PlayerHealth:     playerHealth,
		MonsterHealth:    monsterHealth,
		PlayerAttackDmg:  playerAttackDmg,
		MonsterAttackDmg: monsterAttackDmg,
		PlayerHealValue:  playerHealValue,
	}
	interaction.PrintRoundStatistics(&roundData)

	gameRounds = append(gameRounds, roundData)

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
