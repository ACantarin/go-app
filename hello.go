package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramento = 5
const delay = 5

func main() {
	exibirIntroducao()

	for {
		exibirMenu()

		option := lerComando()

		switch option {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida")
			os.Exit(-1)
		}
	}
}

func exibirIntroducao() {
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

	fmt.Println(imprimirSeparador())
}

func exibirMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")
}

func lerComando() int {
	var option int
	fmt.Scan(&option)

	return option
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando monitoramento...")

	hosts := []string{
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/404",
	}

	for i := 0; i < monitoramento; i++ {
		for _, host := range hosts {
			testarHost(host)
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Println(imprimirSeparador())
}

func testarHost(host string) {
	response, _ := http.Get(host)

	if response.StatusCode == 200 {
		fmt.Println("O site", host, "está online")
	} else {
		fmt.Println("O site", host, "está fora do ar com o status", response.StatusCode)
	}
}

func imprimirSeparador() string {
	return "********************************************"
}
