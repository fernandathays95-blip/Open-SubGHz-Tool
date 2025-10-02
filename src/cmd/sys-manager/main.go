package main

import (
	"fmt"
	"log"
	"time"

	// Importa a biblioteca de utilitários do sistema
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	fmt.Println("###################################################")
	fmt.Println("# Sistema 'Zero-Budget-Hack' - Módulo Gerenciador #")
	fmt.Println("###################################################")

	// 1. GERENCIAMENTO DE TEMPERATURA (Se disponível)
	fmt.Println("\n--- Status de Saúde do Hardware ---")
	
	// Nota: A leitura de temperatura (sensors.Temperatures()) pode exigir
	// permissões de root e drivers específicos no Linux.
	
	temp, err := host.SensorsTemperatures()
	if err == nil && len(temp) > 0 {
		for _, t := range temp {
			fmt.Printf(" [Temperatura] %s: %.2f°C\n", t.SensorKey, t.Temperature)
		}
	} else {
		fmt.Println(" [Temperatura] Leitura indisponível. Requer permissões/drivers de sistema.")
	}

	// 2. GERENCIAMENTO DE PROGRAMAS (Processos)
	fmt.Println("\n--- Programas/Processos em Execução (Top 5) ---")
	
	processes, err := process.Processes()
	if err != nil {
		log.Fatalf("Erro ao listar processos: %v", err)
	}

	// Listar os processos (mostrando apenas os 5 primeiros para simplicidade)
	for i, p := range processes {
		if i >= 5 {
			break
		}
		
		name, _ := p.Name()
		cpuPercent, _ := p.CPUPercent()
		
		fmt.Printf(" [%d] %s (PID: %d) | Uso de CPU: %.1f%%\n", i+1, name, p.Pid, cpuPercent)
	}

	// 3. Status Geral da CPU
	percent, _ := cpu.Percent(time.Second, false)
	fmt.Printf("\n[Status Geral] Uso da CPU: %.1f%%\n", percent[0])
}
