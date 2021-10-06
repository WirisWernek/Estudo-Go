package main

import (
	"fmt"
	"time"
)

func main() {
	// só existe uma estrutura de repetição em Go, o for, ele consegue funcionar tanto como while tanto como for de outras linguagens
	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		// a biblioteca time permite utilizar funções que trabalham com unidades de tempo
		// a função sleep faz o programa esperar por um determinado tempo, no exemplo o tempo é de 1 segundo
		time.Sleep(time.Second)
	}
	fmt.Println(i)
	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}
	nomes := [3]string{"João", "Davi", "Lucas"}
	// a função range serve para percorrer uma estrutura e devolve algo(indice e valor) que pode ser armazenado em uma variável(iteração)
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
	// para ignorar uma informação que o range retorna pode se colocar _ no lugar da variável que armazenaria a informação que sera descartada
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	// a iteração por letras retorna o código ASCII e não o caracter pois n existe char em Go
	// para que ele pegue a letra utiliza-se a função string(variavel) para converter o código ASCII em uma string(que vai atuar como um char)
	for indice, letra := range "Palavra" {
		fmt.Println(indice, string(letra))
	}

	// pode se usar a iteração com maps
	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// não é possivel utilizar a iteração com structs
	// type usuarioStruct struct {
	// 	nome      string
	// 	sobrenome string
	// }
	// usuario2 := usuarioStruct{"Zé", "Júnior"}
	// for chave, valor := range usuario2 {
	// 	fmt.Println(chave, valor)
	// }

	// loop infinito
	var cont int
	cont = 0
	for {
		fmt.Println(cont)
		cont++
	}
}
