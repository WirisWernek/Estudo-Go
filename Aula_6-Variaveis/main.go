package main

import "fmt"

func main() {
	// declaração explicita
	var variavel1 string = "variável 1"
	// declaração implicita
	variavel2 := "variável 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)
	var (
		variavel3 string = "lalala"
		variavel4 string = "ieieie"
	)
	variavel5, variavel6 := "variavel 5", "variavel 6"
	fmt.Println(variavel3, variavel4)
	fmt.Println(variavel5, variavel6)

	// declaração de constantes- as formas acima tb se aplicam
	const constante1 string = "Constante 1"
	fmt.Println(constante1)
	// Inversão de valores
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}
