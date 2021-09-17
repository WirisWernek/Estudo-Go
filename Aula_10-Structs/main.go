package main

import "fmt"

type usuario struct {
	nome     string
	unidade  uint8
	endereco endereco
}
type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario
	u.nome = "Wiris"
	u.unidade = 3

	enderecoexec := endereco{"Rua dos bobos", 12}
	usuario2 := usuario{"João", 21, enderecoexec}

	usuario3 := usuario{nome: "Marcia"}
	usuario4 := usuario{unidade: 34}

	fmt.Println("Arquivo Struct")
	fmt.Println(u)
	fmt.Println(usuario2)
	fmt.Println(usuario3)
	fmt.Println(usuario4)
}
