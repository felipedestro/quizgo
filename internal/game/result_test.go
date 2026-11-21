package game

import (
	"testing"

	"github.com/felipedestro/quiz/internal/model"
)

func TestIsApproved(t *testing.T) {
	tests := []struct {
		name      string
		points    int
		questions int
		minimum   int
		want      bool
	}{
		{"100% de acertos", 100, 10, 70, true},
		{"Exatamente 70%", 70, 10, 70, true},
		{"Abaixo de 70%", 60, 10, 70, false},
		{"Acima de 70%", 80, 10, 70, true},
		{"Sem perguntas", 0, 0, 70, false},
		{"Mínimo de 80%", 70, 10, 80, false},
		{"Mínimo de 50%", 60, 10, 50, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := State{
				Points:    tt.points,
				Questions: make([]model.Question, tt.questions),
			}

			got := state.IsApproved(tt.minimum)

			if got != tt.want {
				t.Errorf(
					"IsApproved(%d) = %v, esperado %v",
					tt.minimum, got, tt.want,
				)
			}
		})
	}
}
