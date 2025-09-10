package utils

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Função genérica para obter entrada do usuário
func GetUserInput[T any](message string) (input T) {
	fmt.Print(message)
	fmt.Scanln(&input)
	return
}

// Cria um menu a partir de uma lista de opções exibindo um título opcional
func CreateMenu(options []string, menuTitle ...string) {
	if len(menuTitle) > 0 {
		title := menuTitle[0]
		fmt.Println(title)
	}

	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
}

// Lê o conteúdo de um arquivo e retorna como string
func ReadFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	content := ""
	for {
		line, err := reader.ReadString('\n')
		content += line
		if err != nil {
			if err.Error() == "EOF" { // Fim do arquivo
				break
			}
			return "", err
		}
	}
	return content, nil
}

func RegistreLog(url string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo de log: %v\n", err)
		return
	}
	defer file.Close()

	statusText := "offline"
	if status {
		statusText = "online"
	}

	logEntry := fmt.Sprintf("Timestamp: %s - Site: %s - Status: %s \n", time.Now().Format(time.RFC3339), url, statusText)
	if _, err := file.WriteString(logEntry); err != nil {
		fmt.Printf("Erro ao escrever no arquivo de log: %v\n", err)
	}
}

func PrintLog() {
	content, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo de log: %v\n", err)
		return
	}

	fmt.Println("Conteúdo do arquivo de log:")
	fmt.Println(string(content))
}
