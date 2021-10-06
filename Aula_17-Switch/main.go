package main

import "fmt"

// não é necessário utilizar o break, o switch sai automaticamente
func diasDaSemana1(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}
}

//  é possivel fazer as verificações dentro do case
func diasDaSemana2(numero int) string {
	var diasDaSemana string
	switch {
	case numero == 1:
		diasDaSemana = "Domingo"
		// a cláusula fallthrough faz com que a proxima condição seja executada sem antes ser verificada, descomente para testar
		// fallthrough
	case numero == 2:
		diasDaSemana = "Segunda-Feira"
	case numero == 3:
		diasDaSemana = "Terça-Feira"
	case numero == 4:
		diasDaSemana = "Quarta-Feira"
	case numero == 5:
		diasDaSemana = "Quinta-Feira"
	case numero == 6:
		diasDaSemana = "Sexta-Feira"
	case numero == 7:
		diasDaSemana = "Sábado"
	default:
		diasDaSemana = "Número Inválido"
	}
	return diasDaSemana
}

// a cláusula fallthrough faz com que a proxima condição seja executada sem antes ser verificada
func main() {

	dia1 := diasDaSemana1(1)
	fmt.Println(dia1)
	dia2 := diasDaSemana2(5)
	fmt.Println(dia2)
}
