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
	state.Init(quizFiles)
	state.Run()

	media := 70
	points := len(state.Questions) * 10
	approved := points > 0 && state.Points*100 >= points*media
	reproved := points > 0 && state.Points*100 < points*media

	if approved {
		fmt.Println("\n\033[32m[APROVADO]\033[0m")
		fmt.Printf("\033[32mFim do jogo, parabéns você foi aprovado com %d pontos\033[0m\n", state.Points)
	}

	if reproved {
		fmt.Println("\n\033[31m[REPROVADO]\033[0m")
		fmt.Printf("\033[31mFim do jogo, infelizmente você foi reprovado com %d pontos\033[0m\n", state.Points)
	}
}
