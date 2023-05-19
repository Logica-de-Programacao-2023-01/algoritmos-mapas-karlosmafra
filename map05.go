package main

import "fmt"

func main() {
	var str string
	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&str)

	fmt.Println(contarLetras(str))
}

func contarLetras(str string) map[string]int {
	ocorrencias := make(map[string]int)

	for i := range str {
		c := string(str[i])
		ocorrencias[c]++
	}

	return ocorrencias
}
