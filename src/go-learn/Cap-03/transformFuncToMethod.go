// Structs no GoLang
//
//

package main

import "fmt"

// Constitue do conceito de orientação a objeto, sendo possível criar classes e métodos
func main() {
	carro := Carro{
		Nome:   "Celta",
		Marca:  "Chevrolet",
		Modelo: "1.0",
		Valor:  2.000,
	}

	carro.montar()
}

type Carro struct {
	Nome   string
	Marca  string
	Modelo string
	Valor  float64
}

func (c Carro) montar() {
	fmt.Printf("Montando %v da marca %v modelo %v e R$ %v", c.Nome, c.Marca, c.Modelo, c.Valor)
}
