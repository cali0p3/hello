package main

import (
	"fmt"
	"os"
)

func main() {
	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	/* if comando == 1 {
		fmt.Println("Monitorando")
	} else if comando == 2{
		fmt.Println("Exibindo Logs")
	} else if comando == 3 {
		fmt.Println("Saindo do Programa")
	} else {
		fmt.Println("Não reconheço esse comando")
	} */

	switch comando {
	case 1:
		fmt.Println("Monitorando")
	case 2:
		fmt.Println("Exibindo Logs")
	case 3:
		fmt.Println("Saindo do Programa")
		os.Exit(0)
	default:
		fmt.Println("Não reconheço esse comando")
		os.Exit(-1)
	}

}
func exibeIntroducao() {
	nome := "Caliope"
	versao := 1.2
	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa eta na versão:", versao)

}

func leComando() int {
	var comandoLido int
	fmt.Scanf("%d", &comandoLido) //o comando fmt.Scan(&comando) pode ser usado visto que o tipo de variavel já foi declarado anteriormente.

	fmt.Println("O comando selecionado foi:", comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair do Programa")
}
