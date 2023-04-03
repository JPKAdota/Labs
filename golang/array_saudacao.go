package main

import (
	"fmt"
	"os"
	"time"
)

func getSalutation() string {
	now := time.Now()
	hour := now.Hour()
	if hour >= 6 && hour < 12 {
		return "Bom dia"
	} else if hour >= 12 && hour < 18 {
		return "Boa tarde"
	} else {
		return "Boa noite"
	}
}

func main() {
	names := []string{"João", "Igor", "Richar", "Symon", "Tatiana", "Danilo"}
	file, err := os.Create("array_saudacoes_go.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	salutation := getSalutation()

	for _, name := range names {
		message := fmt.Sprintf("%s, %s!\n", salutation, name)
		fmt.Print(message)
		_, err := file.WriteString(message)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Saudações salvas no arquivo array_saudacoes_go.txt")
}

