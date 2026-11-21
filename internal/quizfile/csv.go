package quizfile

import (
	"encoding/csv"
	"fmt"
	"io/fs"

	"github.com/felipedestro/quiz/internal/helper"
	"github.com/felipedestro/quiz/internal/model"
)

func ProcessCSV(files fs.FS, fileSelect string) ([]model.Question, error) {
	file, err := files.Open(fileSelect)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	questions := make([]model.Question, 0, len(records))

	for index, record := range records {
		if index == 0 {
			continue
		}

		if len(record) < 6 {
			return nil, fmt.Errorf("linha %d: CSV inválido", index+1)
		}

		correctAnswer, err := helper.ToInt(record[5], 4)
		if err != nil {
			return nil, fmt.Errorf(
				"linha %d: %w", index+1, err,
			)
		}

		question := model.Question{
			Text:    record[0],
			Options: record[1:5],
			Answer:  correctAnswer,
		}

		questions = append(questions, question)
	}

	return questions, nil
}
