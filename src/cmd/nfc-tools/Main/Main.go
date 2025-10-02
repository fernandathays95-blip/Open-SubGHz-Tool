package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ecc1/go-pn532/pn532"
)

// Defina aqui a porta serial ou método de comunicação com seu chip NFC.
// Por exemplo: "/dev/ttyUSB0" (para USB/Serial) ou um endereço SPI/I2C.
const devicePath = "/dev/ttyAMA0" // Exemplo para Raspberry Pi com Serial/UART

func main() {
	fmt.Println("###################################################")
	fmt.Println("# Sistema 'Zero-Budget-Hack' - Módulo NFC Tools   #")
	fmt.Println("###################################################")

	// 1. Conectar ao chip PN532
	fmt.Printf("\nConectando ao chip PN532 em %s...\n", devicePath)
	
	// Tentativa de abrir o dispositivo NFC
	nfc, err := pn532.Open(devicePath)
	if err != nil {
		log.Fatalf("ERRO: Não foi possível abrir o dispositivo PN532. Verifique o caminho (%s) e a conexão.\n%v", devicePath, err)
	}
	defer nfc.Close() // Garante que o chip seja fechado ao final

	// 2. Verificar a versão do firmware
	v, err := nfc.FirmwareVersion()
	if err != nil {
		log.Fatalf("ERRO: Falha ao ler a versão do firmware: %v", err)
	}
	fmt.Printf("Sucesso! PN532 encontrado. Versão do Firmware: %X\n", v)

	// 3. Loop de Leitura de Tags
	fmt.Println("\nPRONTO PARA ESCANEAR: Aproxime uma tag NFC ou cartão...")
	
	// Configura para ler tags tipo A (MiFare, mais comuns)
	target := pn532.TargetTypeA

	// Loop infinito para escutar tags
	for {
		// A função 'InListPassiveTarget' espera que uma tag seja colocada
		tag, err := nfc.InListPassiveTarget(target)
		
		if err == nil {
			fmt.Println("---------------------------------------------------")
			fmt.Printf("-> TAG NFC ENCONTRADA!\n")
			fmt.Printf("   UID (ID Único): %X\n", tag.UID)
			fmt.Printf("   Tipo de Tecnologia: %s\n", target.String())
			fmt.Println("---------------------------------------------------")
			
			// Espera um pouco para evitar leituras duplicadas imediatas
			time.Sleep(2 * time.Second) 
		} else if err != pn532.ErrTimeout {
			// Ignora o erro de timeout (que é normal quando não há tag)
			log.Printf("Erro durante a leitura: %v\n", err)
		}

		time.Sleep(200 * time.Millisecond) // Pequeno intervalo para não sobrecarregar
	}
}
