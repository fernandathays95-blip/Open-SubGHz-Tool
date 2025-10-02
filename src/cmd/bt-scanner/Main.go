package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-ble/ble"
	"github.com/go-ble/ble/linux" // Se estiver usando Linux (recomendado)
)

// A duração do escaneamento (em segundos)
const scanDuration = 10 * time.Second 

func main() {
	fmt.Println("###################################################")
	fmt.Println("# Sistema 'Zero-Budget-Hack' - Módulo BT Scanner  #")
	fmt.Println("###################################################")
	
	// A biblioteca 'ble' precisa de um dispositivo de interface Bluetooth.
	// O 'linux.NewDevice()' tenta configurar o adaptador Bluetooth do seu sistema.
	d, err := linux.NewDevice()
	if err != nil {
		fmt.Printf("\nERRO: Falha ao abrir o adaptador Bluetooth: %v\n", err)
		fmt.Println("Certifique-se de que o Bluetooth está ligado e você tem permissões (sudo/root).")
		return
	}
	ble.SetDefaultDevice(d)

	fmt.Printf("\nIniciando escaneamento de Bluetooth por %v...\n", scanDuration)
	fmt.Println("---------------------------------------------------")
	
	// 'context.WithTimeout' define o tempo máximo que o escaneamento vai rodar.
	ctx, cancel := context.WithTimeout(context.Background(), scanDuration)
	defer cancel() // Garante que o contexto seja cancelado após o escaneamento

	// Escaneia por 10 segundos, chamando a função 'advHandler' para cada dispositivo encontrado.
	err = ble.Scan(ctx, true, advHandler, nil) 
	
	if err != nil && err != context.DeadlineExceeded {
		fmt.Printf("\nERRO durante o escaneamento: %v\n", err)
	}

	fmt.Println("---------------------------------------------------")
	fmt.Println("Escaneamento concluído.")
}

// advHandler é chamada para cada dispositivo que é encontrado durante o escaneamento.
func advHandler(a ble.Advertisement) {
	// A maioria dos celulares vai aparecer com um nome
	name := a.LocalName()
	if name == "" {
		name = "(Nome Desconhecido)"
	}
	
	// Imprime as informações mais relevantes
	fmt.Printf(" [Dispositivo Encontrado] MAC: %s | Nome: %s | RSSI: %d dBm\n", 
		a.Addr(),          // Endereço Bluetooth (similar ao MAC)
		name,              // O nome do dispositivo (ex: iPhone de...)
		a.RSSI())          // Força do sinal (RSSI)
}
