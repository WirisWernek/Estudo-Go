package main

import "fmt"

func main() {
	numero := 10
	// a condição do if e else if pode ou não ser utilizado os ()
	// O Go não aceita que o if e else sejam declarados sem as {}
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}
	// no Go é possível inicializar uma váriável junto com o if, isso se chama if init
	// Porém ela fica atrelada ao escopo do if e else, ou seja, ela não existe no resto da execução
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que 0")
	} else if outroNumero < 0 {
		fmt.Println("Número é menor que 0")
	} else {
		fmt.Println("Número é igual a 0")
	}
}
