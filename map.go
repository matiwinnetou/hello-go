package main

import "fmt"

func main() {
	var mapa map[string]int = make(map[string]int)
	mapa["key1"] = 1
	mapa["key2"] = 2

	for key, value := range mapa {
		fmt.Println("Key:", key, "Value:", value)
	}
	var val, ok = mapa["key1"]
	if ok {
		fmt.Println("val found:", val)
	}
}
