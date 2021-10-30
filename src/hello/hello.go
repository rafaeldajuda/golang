package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitorametos = 2
const delay = 5

func main() {

	exibeIntoducao()

	for {

		exibeMenu()

		// if comando == 1 {
		// 	fmt.Println("Iniciando monitoramento...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo logs...")
		// } else if comando == 0 {
		// 	fmt.Println("Saindo do programa")
		// } else {
		// 	fmt.Println("Não conheço este comando")
		// }

		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

	}

}

func exibeIntoducao() {
	nome := "Rafael"
	versao := 1.1
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando monitoramento...")

	sites := leSitesArquivo()

	for i := 0; i < monitorametos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			// fmt.Printf("Testando site %d: %s", i, site)
			testeSite(site)
		}
		fmt.Println("")
		if i+1 != monitorametos {
			time.Sleep(delay * time.Second)
		}
	}

	fmt.Println("")

}

func testeSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problema. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}

func exibeNomes() {
	nomes := []string{"Rafael", "Daniel", "Douglas"}
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Aparecida")

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes))
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "João")

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes))
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Fernando")

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes))
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Fagner")

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes))
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	// var sites [4]string
	// sites[0] = "http://random-status-code.herokuapp.com"
	// sites[1] = "https://www.alura.com.br"
	// sites[2] = "https://www.caelum.com.br"
	// fmt.Println(sites)
	// fmt.Println(reflect.TypeOf(sites))
}
