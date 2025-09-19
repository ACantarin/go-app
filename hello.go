package main

import (
	"fmt"
	"time"
)

func main() {
	name := "André"
	version := 1.0

	meses := []string{
		"janeiro", "fevereiro", "março", "abril", "maio", "junho",
		"julho", "agosto", "setembro", "outubro", "novembro", "dezembro",
	}

	now := time.Now()
	dia := now.Day()
	mes := meses[now.Month()-1]
	ano := now.Year()

	fmt.Println("Bom dia,", name)
	fmt.Printf("Hoje é %d de %s de %d\n", dia, mes, ano)
	fmt.Printf("Você está usando a versão %.1f\n", version)

	fmt.Println("********************************************")

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")

	var option int
	fmt.Scan(&option)

	switch option {
	case 1:
		fmt.Println("Iniciando monitoramento...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo...")
	default:
		fmt.Println("Opção inválida")
	}
}
