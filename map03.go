package main

import "fmt"

func main() {
	mapa := map[int]int{
		0: 5,
		1: 3,
		2: 8,
		3: 12,
		4: 7,
	}

	fmt.Println(somar(mapa))
}

func somar(mapa map[int]int) int {
	var sum int
	for _, v := range mapa {
		sum += v
	}

	return sum
}
