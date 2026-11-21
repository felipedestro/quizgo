package game

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"

	"github.com/felipedestro/quiz/internal/helper"
	"github.com/felipedestro/quiz/internal/model"
	"github.com/felipedestro/quiz/internal/quizfile"
	"github.com/felipedestro/quiz/internal/stopwatch"
)

func (g *State) selectCategory(files fs.FS) ([]model.Question, error) {
	categories := map[int]QuizCategory{
		1: {ID: 1, Text: "Conhecimentos Gerais", File: "/internal/data/quiz_geral.csv"},
		2: {ID: 2, Text: "Inglês", File: "/internal/data/quiz_ingles.csv"},
		3: {ID: 3, Text: "Matemática", File: "/internal/data/quiz_matematica.csv"},
	}

	fmt.Println("\n-------------------------------------------------")
	fmt.Printf("Seja bem vindo, %s", g.Name)
	fmt.Println("-------------------------------------------------\n")
	fmt.Println("Selecione um categoria para jogar: ")

	for category := range categories {
		fmt.Printf("\033[33m[%d] - %s\033[0m\n", categories[category].ID, categories[category].Text)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		read, _ := reader.ReadString('\n')
		answer, err := helper.ToInt(read[:len(read)-1], 3)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		g.Quiz = categories[answer].File

		questions, err := quizfile.ProcessCSV(files, g.Quiz)

		if err != nil {
			return nil, err
		}

		return questions, nil
	}
}

func (g *State) Init(files fs.FS) {
	fmt.Println("Seja bem vindo(a) ao quiz")
	fmt.Println("Escreva o seu nome: ")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Erro ao ler a string")
	}

	g.Name = name

	questions, err := g.selectCategory(files)

	if err != nil {
		fmt.Println("\nErro ao carregar as questões: ", err)
		return
	}

	g.Questions = questions

	minutes := 5
	g.Done = make(chan struct{})

	go stopwatch.StopWatch(minutes, g.Done)

	fmt.Printf("\033[32m\nVocê tem %d minuto(s) para terminar o jogo!\033[0m\n\n", minutes)
}

func (g *State) Run() {
	for index, question := range g.Questions {
		fmt.Printf("\033[33m %d. %s \033[0m\n", index+1, question.Text)

		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}

		fmt.Println("Digite uma alternativa: ")

		var answer int
		var err error

		for {
			reader := bufio.NewReader(os.Stdin)
			read, _ := reader.ReadString('\n')
			answer, err = helper.ToInt(read[:len(read)-1], 4)

			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			break
		}

		if answer == question.Answer {
			fmt.Println("\033[32m Paranbés você acertou!! \033[0m")
			fmt.Println("______________________________________")
			g.Points += 10
		} else {
			fmt.Println("\033[31m Ops, Errou! \033[0m")
			fmt.Println("______________________________________")
		}
	}
}
