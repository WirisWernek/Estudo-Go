package main

import (
	"errors"
	"fmt"
)

func main() {
	// tipos de numeros
	// int
	// int8 suporta 8 bits até 127 = 2⁷
	// int16 suporta 16 bits até 32767 = 2¹⁵
	// int32 suporta 32 bits até 2147483647 = 2³¹
	// int64 suporta 64 bits até 9,223372037×10¹⁸ = 2⁶³
	var numero1 int8 = 127
	// dessa forma o int assume o valor da arquitetura do computador. Ex: arm64 = int64, x32 = int32
	var numero2 int = -128
	fmt.Println(numero1)
	fmt.Println(numero2)

	// uint - unsygned int - não suporta valores negativos
	// uint8 suporta 8 bits = 2⁷
	// uint16 suporta 16 bits = 2¹⁵
	// uint32 suporta 32 bits = 2³¹
	// uint64 suporta 64 bits = 2⁶³
	var numero3 uint32 = 55
	fmt.Println(numero3)

	// alias
	// int32 = rune
	var numero4 rune = 4566
	fmt.Println(numero4)

	// uint8 = byte
	var numero5 byte = 124
	fmt.Println(numero5)

	// float
	// float32suporta 32 bits = 2³¹
	// float64 suporta 64 bits = 2⁶³
	// nao pode ser declarado domente "float"pode ser declarado por inferência :=, fessa forma o float assume o valor da arquitetura do computador: Ex numero := 121.21111 , assume que é um tipo float

	var numero6 float32 = 124.44545456465
	fmt.Println(numero6)

	var numero7 float64 = 87874488.1464454545876984
	fmt.Println(numero7)

	numero8 := 48546946489.4565465456465446
	fmt.Println(numero8)

	// strings
	// aceita inferência e deve ser sempre declarado com aspas duplas

	var str1 string = "variavel tipo string"
	fmt.Println(str1)

	str2 := "variavel string ehhh"
	fmt.Println(str2)

	// char
	// por padrao o char nao existe no go, a declaração em aspas simples retorna o numero do caracter na tabela ASCC, só aceita um caracter e é convertido a um int
	char := 'b'
	fmt.Println(char)

	// Conceito valor zero
	// variáveis declaradas mas n inicializadas com um valor assumem o valor zero do tipo da variável. Ex: string = " ", int = 0, boll = false, float = 0, error = nil

	var str3 string
	fmt.Println(str3)

	var numero9 int
	fmt.Println(numero9)

	// bool
	var booleano1 bool = false
	fmt.Println(booleano1)

	var booleano2 bool = true
	fmt.Println(booleano2)

	// error
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
