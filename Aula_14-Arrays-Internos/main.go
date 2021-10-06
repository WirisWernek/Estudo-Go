package main

import "fmt"

func main() {
	// a função make aloca na memoria o espaço necessário para alguma estrutura
	// ela recebe como parâmetros: 1º o tipo, 2º o tamanho e 3º a capcidade(máximo)
	// o array interno é o que é criado com a função make, que basicamente cria um array com o tamanho que você passou como máximo e retorna um slice com os valores que você passou como tamanho
	slice := make([]float32, 10, 11)
	fmt.Println(slice)
	// a função len recebe como parametro um array ou slice e retorna o seu tamanho
	fmt.Println(len(slice))
	// a função cap recebe como parametro um array ou slice e retorna a sua capacidade(máximo)
	fmt.Println(cap(slice))
	// o slice não possui limite pois toda vez que o limite é estourado o Go retorna uma nova slice com o dobro do tamanho maximo da anterior, ou seja se o máximo era 11 e você tentou colocar 12 o Go transforma a sua slice em uma slice com máximo de 24
	slice = append(slice, 0)
	slice = append(slice, 0)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	// se um valor máximo não for passado o Go assume que o máximo é o tamanho do slice
	slice2 := make([]float32, 6)
	fmt.Println(slice2)
	slice2 = append(slice2, 0)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
}
