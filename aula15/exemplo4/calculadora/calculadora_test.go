package calculadora

import "testing"

// Benchmark para medir o desempenho da função Somar.
func BenchmarkSomar(b *testing.B) {

	// O laço será executado automaticamente milhares ou milhões de vezes.
	for i := 0; i < b.N; i++ {

		// Código cuja performance será medida.
		Somar(10, 20)
	}
}
