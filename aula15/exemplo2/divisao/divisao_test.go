package divisao

import "testing"

func TestDividir(t *testing.T) {

	resultado, err := Dividir(20, 5)

	if err != nil {
		t.Fatal("Não deveria retornar erro.")
	}

	if resultado != 5 {
		t.Errorf("Esperado 5, recebeu %d", resultado)
	}
}
