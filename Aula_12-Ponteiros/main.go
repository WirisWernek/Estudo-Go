package main

import "fmt"

func main() {
	var variavel1 int = 2
	var variavel2 int = variavel1
	fmt.Println("Ponteiros")
	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	var variavel3 int = 34
	// Referenciação = a variavel4 armazena o endereço da variavel3
	var variavel4 *int = &variavel3
	fmt.Println(variavel3, variavel4)
	// Desreferenciação = valor encontrado no endereço armazenado na variavel
	fmt.Println(variavel3, *variavel4)
}
