// Tipos de funções
//
//

package main

import "fmt"

func main() {

	// Usando função simples
	fmt.Println("Retorno da função simples")
	resultadoSoma := soma(2, 2)
	fmt.Println(resultadoSoma)
	fmt.Println("-------------------------")

	// Usando função de múltiplo sretorno
	fmt.Println("Retorno da função de multiplo retorno")
	resultadoMutiplicacao, str := multiplicacao(2, 2)
	fmt.Println(resultadoMutiplicacao, str)
	fmt.Println("-------------------------")

	// Usando função nomeada
	fmt.Println("Retorno da função nomeada")
	resultadoSubtrair := subtrair(2, 1)
	fmt.Println(resultadoSubtrair)
	fmt.Println("-------------------------")

	// Usando função variádica
	fmt.Println("Retorno da função variádica")
	resultadoMutiplicaTudo := multiplicaTudo(4, 8, 12, 20)
	fmt.Println(resultadoMutiplicaTudo)
	fmt.Println("-------------------------")

	// Funções anônimas
	fmt.Println("Retorno da função anônima")
	anonymFunc := func(arrayParams ...int) func() int {
		result := 0

		for _, value := range arrayParams {
			result += value
		}

		return func() int {
			return result * 2
		}
	}

	fmt.Println(anonymFunc(24)())
	fmt.Println("-------------------------")

}

// Função com returno simples
func soma(a int, b int) int {
	return a + b
}

// Função com multiplos retornos
func multiplicacao(a int, b int) (string, int) {
	return "O retorno da expressao sera: ", a * b
}

// Função com retorno nomeado
func subtrair(a int, b int) (result int) {
	result = a - b
	return
}

// Funções variádicas
func multiplicaTudo(x ...int) (result int) {

	for _, value := range x {
		result += value
	}

	return result
}
