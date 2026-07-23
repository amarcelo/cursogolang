package calculadora

import "testing"

func TestSomar(t *testing.T) {
	resultado := Somar(5, 5)
	esperado := 10

	if resultado != esperado {
		t.Errorf("Esperado %d, mas recebeu %d", esperado, resultado)
	}
}
