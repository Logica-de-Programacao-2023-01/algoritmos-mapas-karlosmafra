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
	str1 := scanner.Text()

	fmt.Print("Digite outra frase: ")
	scanner.Scan()
	str2 := scanner.Text()

	fmt.Print("Digite outra frase: ")
	scanner.Scan()
	str3 := scanner.Text()

	map1 := cPalavras(str1)
	map2 := cPalavras(str2)
	map3 := cPalavras(str3)

	var lista = []map[string]int{
		map1, map2, map3,
	}

	fmt.Println(juntarMapas(lista))
}

func cPalavras(frase string) map[string]int {
	ocorrencias := make(map[string]int)
	palavras := strings.Fields(frase)

	for _, p := range palavras {
		ocorrencias[p]++
	}

	return ocorrencias
}

func juntarMapas(lista []map[string]int) map[string]int {
	mapa := make(map[string]int)

	for m := range lista {
		for str, val := range lista[m] {
			mapa[str] += val
		}
	}

	return mapa
}
