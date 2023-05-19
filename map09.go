package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	fmt.Println(fibonacci(n))
}

func fibonacci(n int) map[int]int {
	mapa := make(map[int]int)
	mapa[0] = 1
	mapa[1] = 1

	for i := 2; i < n; i++ {
		mapa[i] = mapa[i-1] + mapa[i-2]
	}

	return mapa
}
