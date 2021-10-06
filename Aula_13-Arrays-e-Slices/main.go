package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")
	// array declarado com tamanho fixo
	var array1 [5]int
	fmt.Println(array1)
	// array declarado com inferência de tipo e valor em uma posição
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)
	// array declarado com tamanho baseado na quantidade de item que ele contem
	// mesmo dessa forma ele fica com o tamnho fixo
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// os slices não possuem um tamamho fixo, porém possuem um tipo diferente dos arrays
	slice := []int{10, 23, 53}
	fmt.Println(slice)
	// a inclusão de um elemento no slice é atraves da função append que adiciona o elemento no final do slice
	slice = append(slice, 45)
	// a função append recebe como paramêtro a variável que contém os elementos anteriores e o novo elemento a ser adicionado, e retorna um novo slice com  todos os elementos, este pode ser armazenando no próprio slice passado como parametro ou em um novo
	slice2 := []int{}
	slice2 = append(slice, 39)
	fmt.Println(slice)
	fmt.Println(slice2)

	// as slices podem ser resumidas como "fatias de arrays", assim sendo é possivel armazenar uma parte de um array em uma slice
	// o primeiro valor é incluido e o último é excluido
	// Ex: [1:3], as posições que serão armazenadas são: 1 e 2 apenas
	slice3 := array2[1:4]
	fmt.Println(slice3)

	// os slices possuem um comportamento similar aos ponteiros, alteração feitas nas variáveis que elas pegaram partes são refletidas nas slices
	array2[1] = "Posição alterada"
	fmt.Println(slice3)
}
