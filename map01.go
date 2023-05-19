package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	frase := scanner.Text()

	fmt.Println(contarPalavras(frase))
}

func contarPalavras(frase string) map[string]int {
	ocorrencias := make(map[string]int)
	palavras := strings.Fields(frase)

	for _, p := range palavras {
		ocorrencias[p]++
	}

	return ocorrencias
}
