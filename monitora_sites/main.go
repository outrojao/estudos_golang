package main

import (
	"fmt"
	"monitora_sites/utils"
	"net/http"
	"os"
	"strings"
	"time"
)

func monitorWebsites(urls ...string) {
	const attempts int = 5
	const interval time.Duration = 5 * time.Second
	for i := 0; i < attempts; i++ {
		fmt.Printf("Tentativa %d de monitoramento:\n", i+1)
		for _, url := range urls {
			fmt.Printf("Monitorando o site: %s\n", url)
			testWebsite(url)
			fmt.Println("-----------------------------")
		}
		if i < attempts-1 {
			time.Sleep(interval)
		}
	}
}

func testWebsite(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro ao acessar %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("O site %s está online.\n", url)
		utils.RegistreLog(url, true)
	} else {
		fmt.Printf("O site %s está offline. Status: %d\n", url, resp.StatusCode)
		utils.RegistreLog(url, false)
	}
}

func getURLsToMonitor() (urls []string) {
	for {
		utils.CreateMenu([]string{"Adicionar URL", "Monitorar"})
		switch choice := utils.GetUserInput[int]("Escolha uma opção: "); choice {
		case 1:
			url := utils.GetUserInput[string]("Digite a URL do site: ")
			urls = append(urls, url)
		case 2:
			return urls
		default:
			fmt.Println("Opção inválida.")
		}
	}
}

func main() {
	fmt.Println("Bem-vindo ao monitorador de sites!")
	for {
		utils.CreateMenu(
			[]string{
				"Monitorar site",
				"Carregar URLs de um arquivo",
				"Exibir logs",
				"Sair",
			},
			"Menu:",
		)

		switch choice := utils.GetUserInput[int]("Escolha uma opção: "); choice {
		case 1: // Monitorar site
			urls := getURLsToMonitor()
			monitorWebsites(urls...)
		case 2: // Carregar URLs de um arquivo
			filePath := utils.GetUserInput[string]("Digite o caminho do arquivo: ")
			content, err := utils.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Erro ao ler o arquivo: %v\n", err)
				continue
			}
			lines := strings.Split(content, "\n")
			var urls []string
			for _, line := range lines {
				url := strings.TrimSpace(line)
				if url != "" {
					urls = append(urls, url)
				}
			}
			monitorWebsites(urls...)
		case 3: // Exibir logs
			utils.PrintLog()
		case 4: // Sair
			os.Exit(0)
		default: // Opção inválida
			fmt.Println("Opção inválida.")
			os.Exit(0)
		}
	}
}
