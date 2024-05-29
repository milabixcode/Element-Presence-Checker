package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	readSlice := func(prompt string, length int) []string {
		var slice []string
		fmt.Println(prompt)
		fmt.Print(" Insira os elementos (ou pressione Enter para finalizar): ")
		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "" {
				break
			}

			if len(input) > length {
				fmt.Printf("Error: String '%s' é maior que %d characters. Insira uma string válida \n", input, length)
				continue
			}
	
			//  Preenche a string com zeros se for menor que o comprimento requerido
			for len(input) < length {
				input = "0" + input
			}

			slice = append(slice, input)
		}
		return slice
	}

	// Primeiro slice
	slice1 := readSlice("Insira os elementos do primeiro slice: ", 6)

	// Segundo slice
	slice2 := readSlice("Insira os elementos do segundo slice: ", 6)

	presenteNaLista, ausenteNaLista := checkElements(slice1, slice2)
	fmt.Println("Elementos presentes: ", presenteNaLista)
	fmt.Println("Elementos ausentes: ", ausenteNaLista)

}

func checkElements(slice1, slice2 []string) (presenteNaLista, ausenteNaLista []string) {
	// Armazena a presença de cada string do 'slice2'
	elementMap := make(map[string]bool)

	// Iteração no 'slice2' e marcação de cada string como presente no map
	for _, item := range slice2 {
		elementMap[item] = true
	}

	// Iteração no slice1 e verificação se cada string está presente no 'elementMap'
	for _, item := range slice1 {
		if elementMap[item] {
			presenteNaLista = append(presenteNaLista, item)
		} else {
			ausenteNaLista = append(ausenteNaLista, item)
		}
	}
	return presenteNaLista, ausenteNaLista
}