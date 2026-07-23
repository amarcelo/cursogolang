package divisao

import "errors"

func Dividir(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("divisão por zero")
	}

	return a / b, nil
}
