package helper

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

func ToInt(s string, countOptions int) (int, error) {
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0, errors.New("não é permitido caractere diferente de número, por favor insira um número")
	}

	if i == 0 || i > countOptions {
		return 0, errors.New("não existe a opção digitada")
	}

	return i, nil
}

func ReadAnswer(reader *bufio.Reader, answers chan<- string) {
	for {
		input, err := reader.ReadString('\n')

		if err != nil {
			close(answers)
			return
		}

		answers <- strings.TrimSpace(input)
	}
}
