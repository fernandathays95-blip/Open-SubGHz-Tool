package main

import (
	"fmt"
	"net"
	"time"
)

// Portas comuns usadas por câmeras IP, DVRs e NVRs
var cameraPorts = []string{"80", "8080", "554", "8554", "443", "37777"} 

// Altere para o prefixo da sua rede (ex: "192.168.1.")
const networkPrefix = "192.168.0." 
const timeout = 500 * time.Millisecond

func main() {
	fmt.Println("###################################################")
	fmt.Println("# Sistema 'Zero-Budget-Hack' - Módulo Câmera      #")
	fmt.Println("###################################################")

	fmt.Printf("\nIniciando varredura de câmeras na rede %sX...\n", networkPrefix)

	// Varre os últimos 255 IPs da rede (de 1 a 255)
	for i := 1; i <= 255; i++ {
		ip := fmt.Sprintf("%s%d", networkPrefix, i)
		
		// Verifica as portas comuns de câmera para cada IP
		for _, port := range cameraPorts {
			address := fmt.Sprintf("%s:%s", ip, port)
			
			// Tenta estabelecer uma conexão TCP. Se conectar, significa que a porta está aberta
			conn, err := net.DialTimeout("tcp", address, timeout)
			
			if err == nil {
				fmt.Printf(" -> Câmera/Dispositivo de Vídeo encontrado em: %s\n", address)
				// Aqui você adicionaria lógica para tentar identificar o fabricante (banner grabbing)
				conn.Close()
			}
		}
	}
	
	fmt.Println("\nVarredura de portas comuns de câmera concluída.")
}
