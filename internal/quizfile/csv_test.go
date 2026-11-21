package quizfile

import (
	"testing"
	"testing/fstest"
)

func TestProcessCSV(t *testing.T) {
	tests := []struct {
		name    string
		csv     string
		wantErr bool
		wantLen int
	}{
		{
			name: "CSV válido",
			csv: "Pergunta,Opção 1,Opção 2,Opção 3,Opção 4,Resposta\n" +
				"Quanto é 2+2?,1,2,3,4,4\n",
			wantLen: 1,
		},
		{
			name: "Resposta fora do limite",
			csv: "Pergunta,Opção 1,Opção 2,Opção 3,Opção 4,Resposta\n" +
				"Quanto é 2+2?,1,2,3,4,5\n",
			wantErr: true,
		},
		{
			name: "Colunas insuficientes",
			csv: "Pergunta,Opção 1,Opção 2,Opção 3,Opção 4,Resposta\n" +
				"Quanto é 2+2?,1,2\n",
			wantErr: true,
		},
		{
			name:    "CSV sem perguntas",
			csv:     "Pergunta,Opção 1,Opção 2,Opção 3,Opção 4,Resposta\n",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files := fstest.MapFS{
				"quiz.csv": &fstest.MapFile{
					Data: []byte(tt.csv),
				},
			}

			questions, err := ProcessCSV(files, "quiz.csv")

			if (err != nil) != tt.wantErr {
				t.Fatalf(
					"erro = %v, esperado erro = %v",
					err, tt.wantErr,
				)
			}

			if err == nil && len(questions) != tt.wantLen {
				t.Errorf(
					"quantidade = %d, esperado %d",
					len(questions), tt.wantLen,
				)
			}
		})
	}
}
