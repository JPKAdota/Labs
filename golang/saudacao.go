package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Obter o horário atual
	hora := time.Now().Hour()

	// Definir a saudação de acordo com o horário
	var saudacao string
	if hora >= 6 && hora < 12 {
		saudacao = "Bom dia"
	} else if hora >= 12 && hora < 18 {
		saudacao = "Boa tarde"
	} else {
		saudacao = "Boa noite"
	}

	// Solicitar o nome do usuário
	fmt.Print("Digite seu nome: ")
	var nome string
	fmt.Scanln(&nome)

	// Criar a saudação completa
	saudacaoCompleta := fmt.Sprintf("%s, %s", saudacao, nome)

	// Salvar a saudação em um arquivo txt
	arquivo, err := os.Create("saudacao.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo")
		return
	}
	defer arquivo.Close()

	_, err = arquivo.WriteString(saudacaoCompleta)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo")
		return
	}

	fmt.Println(saudacaoCompleta)
}
