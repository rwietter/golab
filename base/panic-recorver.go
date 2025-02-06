package base

import (
	"errors"
	"fmt"
)

func soma(n1 int, n2 int) {
	if recov := recover(); recov != nil {
		med := (n1 + n2) / 2

		if med < 7 {
			fmt.Println("Ops... A média é inferior a 7, ou seja,", med)
		}
	}
}

func media(n1, n2 int) string {
	defer soma(n1, n2)
	med := (n1 + n2) / 2

	if med < 7 {
		panic("A média é inferior a 7!")
	}

	return "A média é maior que 7!"
}

/**
* Panic interrompe a execução do programa, mas antes de interromper, ele executa funções defer.
* Recover é usado para recuperar de um panic.
* Panic é usado para erros irreversíveis, como falha de alocação de memória.
* Recover só é útil para evitar que um programa pare de funcionar.
 */
func PanicRecover() {
	result := media(8, 10)

	if errors.New(result) != nil && result != "" {
		fmt.Println(result)
	}
}
