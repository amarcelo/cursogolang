package calculadora

import "testing"

func TestMultiplicar(t *testing.T) {

	testes := []struct {
		nome     string
		a        int
		b        int
		esperado int
	}{
		{"2 x 3", 2, 3, 6},
		{"5 x 0", 5, 0, 0},
		{"4 x 4", 4, 4, 16},
		{"3 x 3", 3, 3, 8},
	}

	for _, teste := range testes {

		resultado := Multiplicar(teste.a, teste.b)

		if resultado != teste.esperado {
			t.Errorf("%s: esperado %d recebeu %d",
				teste.nome,
				teste.esperado,
				resultado)
		}
	}
}
