package main

import "fmt"

func main() {
	var i uint64 = 0
	for {
		i = i + 1
		var f = factorial(i)

		if f == 0 {
			break
		}
	}

	i-- //decrement by one, which was last working
	fmt.Println("fact, max i:", i)
	fmt.Println("fact(i:", i, ") is:", factorial(i))
}

func factorial(x uint64) uint64 {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}
