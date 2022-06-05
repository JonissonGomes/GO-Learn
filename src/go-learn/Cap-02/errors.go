// Tratando erros com GO
//
//

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Sempre quando for trabalhar com requests onde há a possibilidade de retorno sobre erros, trabalhar com essa base
	res, err := http.Get("http://google.com.br")

	// Tratando o erro
	if err != nil {
		log.Fatal("Erro detecado", err.Error())
	}

	// Após tratamento de erro
	fmt.Println(res)
}
