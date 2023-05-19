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

	for i := range palavrasLetras(frase) {
		fmt.Println(i, palavrasLetras(frase)[i])
	}

}

func palavrasLetras(frase string) map[string]map[string]int {
	palavras := strings.Fields(frase)
	mapPs := make(map[string]map[string]int)

	for p, plv := range palavras {
		mapLs := make(map[string]int)
		for l := range palavras[p] {
			c := string(palavras[p][l])
			mapLs[c]++
		}
		mapPs[plv] = mapLs
	}

	return mapPs
}
