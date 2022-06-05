// Variáveis, valores e tipos
//
//

package main

// Sempre que quiser utilizar uma biblioteca externa, você pode importar diretamente
import "fmt"

func main() {
	a := 10      // Integer
	b := "Hello" // String
	c := 7.14    // Float64
	d := false   // Boolean
	e := `Welcome
	to
	go`
	//String (String in block)

	// Utilizando o %v é possível ver o valor apresentada na variável
	fmt.Println("Visualizando valores das variáveis")
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)

	// Utilizando o %T é possível visualizar o tipo da variável
	fmt.Println("Visualizando tipos das variáveis")
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)

	// Para criar uma variável sem precisar
	_valueEmpty := "Valor"
	fmt.Printf(_valueEmpty)
}
