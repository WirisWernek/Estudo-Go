package main

import "fmt"

// o tipo ao lado dos parênteses indica o tipo que será retornado pela funçao
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// dessa forma você pode receber mais de um retorno
func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(30, 20)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}
	f("Wiris")
	resultadosoma, resultadosubtracao := calculos(10, 15)
	fmt.Println(resultadosoma, resultadosubtracao)
	// o _ indica que o go deve ignorar um dos valores retornados
	resultadosoma1, _ := calculos(41, 23)
	fmt.Println(resultadosoma1)

}
