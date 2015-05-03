package main

import "fmt"

func main() {
	var mapa map[string]int = make(map[string]int)
	mapa["key"] = 1

	var val, ok = mapa["key"]
	if ok {
		fmt.Println("val found:", val)
	}
}
