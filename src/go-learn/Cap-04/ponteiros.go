// Ponteiros
//
//

package main

import "fmt"

func main() {

	a := 10
	fmt.Println(a)  // Mostrando o valor atribuido
	fmt.Println(&a) // Mostrando o endereçamento em memória daquele valor

	// Guardando o endereço de memoria da variável
	var ponteiro *int = &a
	fmt.Println(ponteiro) // Mostrando o endereço na memória

	*ponteiro = 50 // Alterando o valor mantendo o mesmo endereço em memória
	fmt.Println(*ponteiro)
	fmt.Println(ponteiro) // Mostrando o endereço na memória

	alterarEnderecamento(ponteiro)
}

// Exemplo funcional
func alterarEnderecamento(value *int) {
	*value = 200
	fmt.Println(*value)
	fmt.Println(value) // Mostrando o endereço na memória

}
