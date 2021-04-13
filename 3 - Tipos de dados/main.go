package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32, int64
	// int de acordo com processador 32bits ou 64bits
	var numero1 int = 100
	fmt.Println(numero1)

	// uint8, uint16, uint32, uint64
	// uint
	// - somente números positivos
	// - de acordo com processador 32bits ou 64bits
	var numero2 uint = 100
	fmt.Println(numero2)

	// "rune" é igual a "int32"
	var numero3 rune = 12345
	fmt.Println(numero3)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123456789.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.66
	fmt.Println(numeroReal3)

	var str string = "texto"
	fmt.Println(str)

	texto2 := "Texto 2"
	fmt.Println(texto2)

	// Um caractere e aspas simples, ele irá retornar o número correspondente
	// na tabela ascii
	charAscii := 'B'
	fmt.Println(charAscii)

	// Valor "Zero"
	// Existem valores inicias padrão quando são declarados e não definido o valor

	// string: é uma string vazia
	var texto string
	fmt.Println(texto)

	// int: é o número zero
	var numero int16
	fmt.Println(numero)

	// Boleanos
	// Valor zero é false
	var boolean1 bool = true // or false
	fmt.Println(boolean1)

	// Error
	// Valor zero é <nil>
	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)
}
