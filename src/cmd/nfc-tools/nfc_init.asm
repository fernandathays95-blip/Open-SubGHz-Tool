; ----------------------------------------------------------------------
; Arquitetura: x86-64 (Linux)
; Montador: NASM
; Objetivo: Exemplo básico de I/O em Assembly para ilustrar a base do NFC
; ----------------------------------------------------------------------

section .data
    msg db "Iniciando Módulo NFC em Assembly...", 0x0a ; String a ser impressa (0x0a é a quebra de linha)
    len equ $ - msg                                    ; Calcula o tamanho da string

section .text
    global _start 

_start:
    ; -------------------------------------------------
    ; 1. Lógica do 'Bootloader' (Simulação)
    ; Em um hardware real, aqui você inicializaria
    ; o chip PN532 (SPI/I2C).
    ; -------------------------------------------------

    ; MOVE RDI, 0x1000 ; Exemplo: Endereço de memória do registrador I2C

    ; Envia o comando para o Kernel para escrever na tela (sys_write)
    mov rax, 1          ; System call number 1 (sys_write)
    mov rdi, 1          ; File descriptor 1 (stdout - a tela)
    mov rsi, msg        ; Endereço da string (msg)
    mov rdx, len        ; Comprimento da string (len)
    syscall             ; Executa a chamada do sistema (imprime a mensagem)
    
    ; -------------------------------------------------
    ; 2. Lógica de Comunicação NFC
    ; Aqui você teria loops e instruções bit-a-bit:
    ;
    ; mov al, 0x01   ; byte para enviar (endereço de um registrador do PN532)
    ; out 0x20, al   ; Envia o byte para a porta I/O (simulação)
    ; -------------------------------------------------

    ; Envia o comando para o Kernel para sair (sys_exit)
    mov rax, 60         ; System call number 60 (sys_exit)
    mov rdi, 0          ; Código de saída 0 (sucesso)
    syscall             ; Executa a chamada do sistema

