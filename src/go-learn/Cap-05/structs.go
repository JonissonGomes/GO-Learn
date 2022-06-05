// Struct
//
//

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	// Formas de declaração de structs

	// 1ª Forma
	cliente := Cliente{
		name:  "Filho de Jonis",
		email: "jonis@gmail.com",
		cpf:   12345678900,
	}
	fmt.Println(cliente)

	// 2ª Forma
	cliente2 := Cliente{"Jonis", "jonis21@gmail.com", 12345678900}
	fmt.Println(cliente2)
	// Buscando dados específico a partir da 2ª forma
	fmt.Printf("Nome: %s Email: %s CPF: %d\n", cliente2.name, cliente2.email, cliente2.cpf)

	// Chamando struct com "Herança"
	cliente3 := ClienteEstrangeiro{
		Cliente: Cliente{
			name:  "Jonisson",
			email: "jonisson@gmail.com",
			cpf:   12345678900,
		},
		nacionalidade: "EUA",
	}
	fmt.Printf("Nome: %s Email: %s CPF: %d Nacionalidade: %s\n", cliente3.name, cliente3.email, cliente3.cpf, cliente3.nacionalidade)

	cliente.ExibirCliente()
	cliente2.ExibirCliente()
	cliente3.ExibirCliente()

	// Criando JSON a partir de uma Struct
	//
	//
	// Marshal transforma a struct para json
	clienteJson, err := json.Marshal(cliente2)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(clienteJson)) // Converte o byte para string e mostra o JSON

	// Como ler um arquivo JSON
	createJsonCliente := `{"name": "John", "email":"john@gmail.com", "cpf":12345678900, "nacionalidade":"EUA"}`
	cliente4 := ClienteEstrangeiro{}
	json.Unmarshal([]byte(createJsonCliente), &cliente4) // Alocando o JSON no mesmo endereço de memória do cliente4
	fmt.Println(cliente4)
}

type Cliente struct {
	name  string
	email string
	cpf   int
}

type ClienteEstrangeiro struct {
	Cliente
	nacionalidade string
	// nacionalidade string `json: "pais"` // Tratando a leitura do json, e adicionando valores padrão
}

func (c Cliente) ExibirCliente() {
	fmt.Println("Cliente: ", c.name)
}
