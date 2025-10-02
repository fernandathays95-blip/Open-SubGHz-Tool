package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// O caminho para a pasta 'cmd' onde estão todas as ferramentas
const toolsPath = "cmd" 

func main() {
	fmt.Println("###################################################")
	fmt.Println("# Sistema 'Zero-Budget-Hack' - INICIALIZANDO      #")
	fmt.Println("###################################################")

	// 1. Encontrar todas as ferramentas disponíveis
	tools, err := findTools()
	if err != nil {
		log.Fatalf("Erro ao ler a pasta de comandos: %v", err)
	}

	if len(tools) == 0 {
		fmt.Println("Nenhuma ferramenta encontrada na pasta cmd. O sistema está vazio.")
		return
	}

	// 2. Apresentar o Menu ao Usuário
	reader := bufio.NewReader(os.Stdin)
	for {
		displayMenu(tools)

		fmt.Print("\nSelecione uma ferramenta (digite o número) ou 'q' para sair: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		if strings.ToLower(input) == "q" {
			fmt.Println("\nSistema desligado. Até mais!")
			return
		}

		// 3. Processar a escolha do Usuário
		choice, err := processChoice(input, tools)
		if err != nil {
			fmt.Println("\nERRO: Seleção inválida. Tente novamente.")
			continue
		}

		// 4. Rodar a Ferramenta Selecionada
		runTool(choice)
	}
}

// findTools lista todas as subpastas dentro da pasta 'cmd' (que são as ferramentas)
func findTools() ([]string, error) {
	entries, err := os.ReadDir(toolsPath)
	if err != nil {
		return nil, err
	}

	var tools []string
	// Itera sobre as entradas e só adiciona pastas (que são nossas ferramentas)
	for _, entry := range entries {
		// Ignora o próprio menu principal
		if entry.IsDir() && entry.Name() != "main_menu" && entry.Name() != "nfc-tools" {
			tools = append(tools, entry.Name())
		}
	}
	return tools, nil
}

// displayMenu exibe o menu interativo
func displayMenu(tools []string) {
	fmt.Println("\n===================================================")
	fmt.Println(" Ferramentas Disponíveis:")
	fmt.Println("===================================================")

	for i, tool := range tools {
		// Formata o nome para ficar mais legível no menu
		displayName := strings.ReplaceAll(tool, "-", " ")
		displayName = strings.ToTitle(displayName)
		fmt.Printf(" [%d] %s\n", i+1, displayName)
	}
}

// processChoice valida a entrada do usuário e retorna o nome da pasta (a ferramenta)
func processChoice(input string, tools []string) (string, error) {
	var index int
	_, err := fmt.Sscanf(input, "%d", &index)
	if err != nil || index < 1 || index > len(tools) {
		return "", fmt.Errorf("escolha inválida")
	}
	// O usuário escolhe 1, 2, 3... mas a lista é 0, 1, 2...
	return tools[index-1], nil
}

// runTool executa o script 'main.go' dentro da pasta da ferramenta selecionada
func runTool(toolName string) {
	fmt.Printf("\n>>> INICIANDO: %s...\n", toolName)
	
	// O comando 'go run' é usado para compilar e executar o script principal da ferramenta
	toolPath := fmt.Sprintf("./%s/main.go", toolName)
	
	// Prepara e executa o comando. NOTA: Este comando depende do 'go' estar instalado.
	cmd := exec.Command("go", "run", toolPath)
	
	// Redireciona a saída do script da ferramenta para o console do menu
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("\n[ERRO DE EXECUÇÃO] Falha ao rodar %s: %v\n", toolName, err)
	} else {
		fmt.Printf("\n<<< %s ENCERRADO. >>>\n", toolName)
	}
	fmt.Println("Pressione ENTER para voltar ao Menu...")
	bufio.NewReader(os.Stdin).ReadBytes('\n') // Espera o ENTER
}
