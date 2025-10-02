# 🐬 Zero-Budget-Hack (Sistema Multi-Ferramentas Open-Source)

> Um projeto DIY (Faça Você Mesmo) de engenharia reversa para criar um sistema de ferramentas portátil, inspirado em dispositivos como o Flipper Zero, usando **Go (Golang)** para velocidade e eficiência.

---

## 🧠 Sobre o Projeto

Este projeto nasceu da ideia de que não precisamos de hardware caro para aprender sobre comunicação sem fio e sistemas de baixo nível. O **Zero-Budget-Hack** é um sistema modular que visa replicar funcionalidades de escaneamento, captura e gerenciamento de rádio, Bluetooth, Wi-Fi e sistemas.

O sistema é construído inteiramente em **Go**, aproveitando sua excelente capacidade de lidar com operações de rede e sistema.

## 🛠️ Módulos e Funcionalidades

O sistema é composto por módulos que podem ser executados via um menu principal interativo:

| Módulo             | Pasta                  | Funcionalidade Principal                                      | Status |
| :----------------- | :--------------------- | :------------------------------------------------------------ | :----- |
| **Menu Principal** | `cmd/main_menu`        | Central de comando que lista e executa todos os programas.    | ✅ |
| **Wi-Fi Scanner** | `cmd/wifi-scanner`     | Escaneamento de redes e identificação de dispositivos (celulares).| ✅ |
| **BT Scanner** | `cmd/bt-scanner`       | Detecção de dispositivos Bluetooth Low Energy (BLE).          | ✅ |
| **RF Scanner/Replay**| `cmd/rf-scanner`       | Lógica de captura, decodificação e repetição de códigos de Rádio Frequência (433MHz). | ✅ |
| **NFC Tools** | `cmd/nfc-tools`        | Esqueleto para comunicação de baixo nível com chips PN532.    | ✅ |
| **Camera Tools** | `cmd/camera-tools`     | Varredura de rede para identificar câmeras IP (RTSP/HTTP).    | ✅ |
| **Sys Manager** | `cmd/sys-manager`      | Monitoramento de processos, uso de CPU e temperatura do hardware. | ✅ |

## 🚀 Como Executar

Para rodar o sistema, você precisa ter o **Go** (versão 1.18+) instalado.

1.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/seu-usuario/Zero-Budget-Hack.git](https://github.com/seu-usuario/Zero-Budget-Hack.git)
    cd Zero-Budget-Hack
    ```

2.  **Baixe as dependências:**
    ```bash
    go mod tidy
    ```

3.  **Inicie o Sistema (Menu Principal):**
    ```bash
    go run ./cmd/main_menu/main.go
    ```

## ⚠️ Nota sobre Hardware

As ferramentas de baixo nível (**RF**, **BT**, **NFC**) são códigos de demonstração da lógica. Para que elas funcionem de fato, é necessário **hardware especializado** conectado ao seu sistema (ex: Adaptador Wi-Fi em Modo Monitor, SDR/RTL-SDR, e um chip PN532).

---
