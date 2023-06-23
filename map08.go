package main

import "fmt"

func main() {
	mapa := map[string]float64{
		"Pedro": 100.5,
		"Ana":   250,
		"Luan":  400.9,
		"Maria": 120,
	}

	fmt.Println(igualar(mapa))
}

func igualar(mapa map[string]float64) map[string]float64 {
	despesas := make(map[string]float64)

	return despesas
}
