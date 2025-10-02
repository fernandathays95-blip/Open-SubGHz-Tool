package main

import (
	"fmt"
	"log"
	
	// Importamos a biblioteca gopacket pcap
	"github.com/google/gopacket/pcap"
)

func main() {
	fmt.Println("###################################################")
	fmt.Println("# Sistema 'Zero-Budget-Hack' - Módulo Wi-Fi Scanner #")
	fmt.Println("###################################################")
	
	fmt.Println("\nPASSOS PARA VER OS DISPOSITIVOS:")
	fmt.Println("1. Listar Interfaces de Rede...")

	// Lista todas as interfaces de rede disponíveis no sistema
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err) // Se houver um erro, o programa para e mostra a mensagem.
	}

	// Itera sobre todas as interfaces encontradas e as imprime
	fmt.Println("\nInterfaces Encontradas:")
	for _, device := range devices {
		// Esta linha nos ajuda a identificar o adaptador Wi-Fi correto
		fmt.Printf("-> Nome: %s\n", device.Name) 
		fmt.Printf("   Descrição: %s\n", device.Description)
		
		// Imprime os endereços IP associados à interface (se houver)
		for _, addr := range device.Addresses {
			fmt.Printf("   Endereço IP: %s\n", addr.IP)
		}
		fmt.Println("---")
	}

	fmt.Println("\n===================================================")
	fmt.Println("PRÓXIMO PASSO: Escolha a interface Wi-Fi acima (Ex: wlan0 ou eth0).")
}
