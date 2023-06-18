package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 2
const delay = 5

func main() {
	Introduction()

	for {
		Menu()

		comando := leCommand()

		switch comando {
		case 1:
			StartMonitoring   ()
		case 2:
			fmt.Println("Showing Logs...")
			PrintLogs()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
			os.Exit(-1)
		}
	}

}

func Introduction() {
	name := "Lorena"
	version := 1.2
	fmt.Println("Hello, mr/ms", name)
	fmt.Println("This is version", version)
}

func Menu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")
}

func leCommand() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("The chosen command is", comandoLido)
	fmt.Println("")

	return comandoLido
}

func StartMonitoring() {
	fmt.Println("Monitoring...")
	sites := leSitesDoArquivo()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testing website", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testWebsite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "is running!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "got some problems. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leArchiveFiles() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("An error occurred:", err)
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

func recordLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func PrintLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}
