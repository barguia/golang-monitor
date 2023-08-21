package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 3
const delay = 5

func main() {

	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLog()
		case 0:
			fmt.Println("Saindo do programa.")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando.")
			os.Exit(-1)
		}
	}
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func exibeIntroducao() {
	nome := "Eduardo"
	versao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão ", versao)
	fmt.Println("")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("A opcao selecionada foi:", comandoLido)
	fmt.Println("")
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := lerSitesDoArquivo()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testanto site ", i, "site:", site)
			testaSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	response, error := http.Get(site)

	if error != nil {
		fmt.Println("Ocorreu um erro", error)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site", site, "está com problema. Status code: ", response.StatusCode)
		registraLog(site, false)
	}
}

func lerSitesDoArquivo() []string {
	file, error := os.Open("sites.txt")
	var sites []string
	if error != nil {
		fmt.Println("Ocorreu um erro", error)
	}

	leitor := bufio.NewReader(file)

	for {
		linha, erro := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		fmt.Println(linha)

		sites = append(sites, linha)

		if erro == io.EOF {
			break
		}
	}
	fmt.Println("")
	file.Close()

	return sites
}

func registraLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	log := time.Now().Format("02/01/2006 15:04:05") + " - " + site + " Online: " + strconv.FormatBool(status) + "\n"
	file.WriteString(log)

	file.Close()
}

func imprimeLog() {
	file, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
