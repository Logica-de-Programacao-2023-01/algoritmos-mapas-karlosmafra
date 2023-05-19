package main

import "fmt"

func main() {
	map1 := map[string]int{
		"maça":    10,
		"banana":  6,
		"laranja": 15,
	}

	map2 := map[string]int{
		"melancia": 5,
		"maça":     15,
		"uva":      8,
	}

	fmt.Println(mapas(map1, map2))
}

func mapas(map1, map2 map[string]int) map[string]int {
	mapa := make(map[string]int)

	for c, v := range map1 {
		mapa[c] = v
	}

	for c, v := range map2 {
		mapa[c] = v
	}

	return mapa
}