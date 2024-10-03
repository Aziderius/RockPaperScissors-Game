package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 //Piedra. Vence a Tijeras. (tijeras + 1) % 3 = 0
	PAPER    = 1 //Papel. Vence a la piedra. (piedra + 1) % 3 = 1
	SCISSORS = 2 //Tijeras. Vence a la papel. (papel + 1) % 3 = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessage = []string{
	"¡Ganaste!",
	"Buen trabajo",
	"Pudiste leer el futuro",
}

var loseMessage = []string{
	"Buuuuuu",
	"Intentalo de nuevo",
	"No pudiste",
	"Aprende a leer el futuro",
}

var drawMessages = []string{
	"Las grandes mentes piensan igual",
	"Empate",
	"Refina tus habilidades",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	//Mensaje dependiendo de lo que eligió la computadora
	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligió PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligió PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligió TIJERA"
	}

	//generar un numero aleatorio de 0-2, que usamos para elegir un mensaje aleatorio
	messageInt := rand.Intn(3)

	//declarar una var para contener el mensaje
	var message string

	if playerValue == computerValue {
		roundResult = "Es un empate"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "El jugador GANA"
		message = winMessage[messageInt]
	} else {
		ComputerScore++
		roundResult = "La computadora GANA"
		message = loseMessage[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}

}
