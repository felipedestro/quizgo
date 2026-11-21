package game

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"time"

	"github.com/felipedestro/quiz/internal/helper"
	"github.com/felipedestro/quiz/internal/model"
	"github.com/felipedestro/quiz/internal/quizfile"
	"github.com/felipedestro/quiz/internal/stopwatch"
)

func (g *State) selectCategory(files fs.FS) ([]model.Question, error) {
	categories := map[int]QuizCategory{
		1: {ID: 1, Text: "Conhecimentos Gerais", File: "internal/data/quiz_geral.csv"},
		2: {ID: 2, Text: "Inglês", File: "internal/data/quiz_ingles.csv"},
		3: {ID: 3, Text: "Matemática", File: "internal/data/quiz_matematica.csv"},
	}

	fmt.Println("\n=================================")
	fmt.Printf("Seja bem vindo, %s\n", g.Name)
	fmt.Println("=================================")
	fmt.Println()
	fmt.Println("Selecione um categoria para jogar: ")

	for category := range categories {
		fmt.Printf("\033[33m[%d] - %s\033[0m\n", categories[category].ID, categories[category].Text)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		read, err := helper.ReadInput(reader)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler a categoria: %w", err)
		}

		answer, err := helper.ToInt(strings.TrimSpace(read), 3)

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

func (g *State) Init(files fs.FS) error {
	fmt.Println("Seja bem vindo(a) ao quiz")
	fmt.Println("Escreva o seu nome: ")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		return fmt.Errorf("Erro ao ler o noem: %w", err)
	}

	g.Name = strings.TrimSpace(name)

	questions, err := g.selectCategory(files)

	if err != nil {
		return fmt.Errorf("\nErro ao carregar as questões: %w", err)
	}

	g.Questions = questions

	minutes := 5

	g.Done = make(chan struct{})
	g.Timeout = make(chan struct{})

	go stopwatch.StopWatch(
		time.Duration(minutes)*time.Minute,
		g.Done,
		g.Timeout,
	)

	fmt.Printf(
		"\033[32m\nVocê tem %d minuto(s) para terminar o jogo!\033[0m\n\n",
		minutes,
	)

	return nil
}

func (g *State) Run() error {
	reader := bufio.NewReader(os.Stdin)
	answers := make(chan string)

	go helper.ReadAnswer(reader, answers)

	for index, question := range g.Questions {
		fmt.Printf("\033[33m %d. %s \033[0m\n", index+1, question.Text)

		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}

		for {
			fmt.Print("Digite uma alternativa: ")

			select {
			case <-g.Timeout:
				return fmt.Errorf("Tempo Esgotado")

			case input, ok := <-answers:
				if !ok {
					return fmt.Errorf("entrada encerrada")
				}

				answer, err := helper.ToInt(input, len(question.Options))

				if err != nil {
					fmt.Println(err)
					continue
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

			break
		}
	}

	return nil
}
