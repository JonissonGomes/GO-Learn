<!-- Pacotes -->

<!-- Simbolos / Operadores -->
:= Declaração direta(Onde a linguagem irá identificar qual a tipagem do dado a partir do valor atribuido)
OBS: Só pode ser usada dentro de um bloco de código {}

= Atribuição

== Comparação

<!-- variáveis -->
var (Utilizada para escopo global)
OBS: Quando não atribuida a algum tipo, irá identificar a tipagem automáticamente

int age

age := 

<!-- Visibilidade de uma função -->
## Caso uma função comece com letra maiúscula, exemplo:

func `ValidateFiled()` {
    fmt.Println("Campo validadeo)
}

Essa função é considerada `pública`, ou seja, ela pode ser visivel fora do pacote na qual está inserida

## Caso uma função comece com letra minúscula, exemplo:

func `PrepareData`() {
    fmt.Println("Dados preparados")
}

Essa função é considerada `privada`, ou seja, ela só pode ser visualizada e utilizada, dentro daquele pacote.