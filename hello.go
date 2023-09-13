package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome = "Calíope"
	nome_2 := "Roberta"      //usando o := podemos declarar uma variavel sem usar a palavra dedicada "var"
	var versao float32 = 1.1 //não é necessario declarar o tipo da variavel o Go decide qual tipo mais indicado de acordo com o valor atribuido a variavel.
	var idade int = 36
	fmt.Println("Hello, Eldorado!")
	fmt.Println("Olá senhor(a),", nome_2, nome, "!")
	fmt.Println("Sua idade é:", idade, ".")
	fmt.Println("A versão dessa aplicação é: ", versao, ".")

	fmt.Println("O tipo da variável versão é:", reflect.TypeOf(versao))
}
