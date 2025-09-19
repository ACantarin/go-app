package main

import (
	"fmt"
	"time"
)

func main() {
	var name string = "André"
	var version float32 = 1.0

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
}
