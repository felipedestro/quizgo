package game

import "github.com/felipedestro/quiz/internal/model"

type State struct {
	Name      string
	Quiz      string
	Points    int
	Questions []model.Question
	Done      chan struct{}
	Timeout   chan struct{}
}

type QuizCategory struct {
	ID   int
	Text string
	File string
}
