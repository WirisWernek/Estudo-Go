package main

import "fmt"

func main() {

	// os maps são similares as structs porém eles precisam ter todos os dados dentro de si do mesmo tipo
	// eles são declarados com o tipo entre parênteses sedo o das chaves e o tipo após o parênteses sendo o dos valores para as chaves
	// algo que não pode acontecer é as chaves ou os valores serem de tipos diferentes, todas as chaves tem que ser do mesmo tipo e todos os valores tem que ser do mesmo tipo porem os valorem podem ser de tipos diferentes das chaves
	// Ex-1:
	usuario1 := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	// Ex-2:
	usuario2 := map[string]int{
		"senha": 1234,
		"login": 21154,
	}

	fmt.Println(usuario1)
	// para acessar um campo em específico do map se utiliza [nomeDoCampo]
	fmt.Println(usuario1["nome"])
	fmt.Println(usuario2)

	// Pode se utilizar os map como valores para as chaves
	usuario3 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}
	fmt.Println(usuario3)
	// para se deletar uma chave utiliza a função delete, onde se passa o nome do map e a chave a ser deletada
	delete(usuario3, "curso")
	fmt.Println(usuario3)

	usuario1["signo"] = "Capricornio"
	fmt.Println(usuario1)
}
