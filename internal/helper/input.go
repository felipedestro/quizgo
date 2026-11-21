package helper

import (
	"errors"
	"strconv"
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
