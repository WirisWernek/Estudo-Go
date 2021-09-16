package main

import "fmt"

func main() {
	// Aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	resto := 10 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	// Atribuição
	var variavel1 string = "string"
	variavel2 := "string"
	fmt.Println(variavel1, variavel2)

	// Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 == 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// Logicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) // operador e
	fmt.Println(verdadeiro || falso) // operador ou
	fmt.Println(!verdadeiro)         // negação

	// Unários
	numero := 10
	numero++ // numero = numero + 1
	fmt.Println(numero)
	numero += 10 // numero = numero + 15
	fmt.Println(numero)
	numero-- // numero = numero - 1
	fmt.Println(numero)
	numero -= 3 // numero = numero -3
	fmt.Println(numero)
	numero *= 2 // numero = numero * 2
	fmt.Println(numero)
	numero /= 2 // numero = numero / 2
	fmt.Println(numero)
	numero %= 3 // numero = numero %3
	fmt.Println(numero)

	// Ternários não existem no go
	// Ex: texto := numero >5 ? "maior que 5" : "menor que 5"

}
