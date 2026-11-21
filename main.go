package main

import (
	"embed"
	"fmt"

	"github.com/felipedestro/quiz/internal/game"
)

//go:embed internal/data/*.csv
var quizFiles embed.FS

func main() {
	state := game.State{}

	if err := state.Init(quizFiles); err != nil {
		fmt.Println("Não foi possível iniciar o jogo", err)
		return
	}

	err := state.Run()

	close(state.Done)

	if err != nil {
		fmt.Println(err)
		return
	}

	if state.IsApproved(70) {
		fmt.Println("\n\033[32m[APROVADO]\033[0m")
		fmt.Printf("\033[32mFim do jogo, parabéns você foi aprovado com %d pontos\033[0m\n", state.Points)
	} else {
		fmt.Println("\n\033[31m[REPROVADO]\033[0m")
		fmt.Printf("\033[31mFim do jogo, infelizmente você foi reprovado com %d pontos\033[0m\n", state.Points)
	}
}
