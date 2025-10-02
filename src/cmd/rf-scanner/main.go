package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sdr-go/sdr/rtl" // Biblioteca para RTL-SDR
)

// Frequência comum para controles remotos (Sub-GHz) em Hz
const frequency = 433920000 

// Taxa de amostragem (quantos dados por segundo)
const sampleRate = 2048000 

func main() {
	fmt.Println("###################################################")
	fmt.Println("# Sistema 'Zero-Budget-Hack' - Módulo RF/SDR      #")
	fmt.Println("###################################################")

	// 1. Encontrar o Dispositivo RTL-SDR
	count := rtl.GetDeviceCount()
	if count == 0 {
		log.Fatal("ERRO: Nenhum dispositivo RTL-SDR (dongle de rádio) encontrado.\nConecte um e garanta que o driver 'rtl-sdr' esteja instalado.")
	}

	fmt.Printf("\nDispositivos RTL encontrados: %d. Usando o Dispositivo 0.\n", count)

	// 2. Abrir o Dispositivo
	dev, err := rtl.Open(0)
	if err != nil {
		log.Fatalf("ERRO: Falha ao abrir o dispositivo: %v", err)
	}
	defer dev.Close()
	
	// 3. Configurar o Rádio
	fmt.Printf("Configurando o rádio para %d Hz...\n", frequency)

	// Define a taxa de amostragem
	if err = dev.SetSampleRate(sampleRate); err != nil {
		log.Fatalf("Falha ao definir Sample Rate: %v", err)
	}

	// Define a frequência
	if err = dev.SetCenterFreq(frequency); err != nil {
		log.Fatalf("Falha ao definir a Frequência: %v", err)
	}
	
	// 4. Iniciar a Recepção (Escuta)
	fmt.Println("PRONTO: Escutando na frequência...")

	// Buffer para armazenar os dados I/Q (dados de rádio)
	buffer := make([]byte, 16*1024) 
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Escaneia por 10 segundos
	defer cancel()
	
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\nEscaneamento de 10 segundos concluído.")
			return
		default:
			// Lê os dados brutos de rádio
			n, err := dev.ReadSync(buffer)
			if err != nil {
				log.Printf("Erro de leitura: %v", err)
				return
			}
			
			// Se dados forem lidos, é um sinal de que algo foi detectado
			if n > 0 {
				fmt.Printf(" [DADOS DETECTADOS] %d bytes lidos. Sinal de RF encontrado! \n", n)
				// Aqui, a lógica de decodificação de controles remotos entraria,
				// analisando os padrões dos bytes lidos (buffer[:n]).
			}
		}
		time.Sleep(10 * time.Millisecond) // Pequeno intervalo
	}
}
