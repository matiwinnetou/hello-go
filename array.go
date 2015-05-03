package main

import "fmt"

func main() {
	var array []int = []int{1, 2, 3, 4, 5}

	var sum1 int = 0
	for i := 0; i < len(array); i++ {
		sum1 += array[i]
	}

	var sum2 int = 0
	for _, value := range array {
		sum2 += value
	}

	fmt.Println("array:", array)
	fmt.Println("sum1:", sum1)
	fmt.Println("sum2:", sum2)

	var m map[string]int = make(map[string]int)
	m["key"] = 1

	fmt.Println("m:", m)

	fmt.Println("average:", average(array...))

	var i uint64 = 1
	for {
		var f = factorial(i)
		fmt.Println("fact:", f, i)

		if f == 0 {
			break
		}
		i = i + 1
	}

	fmt.Println("factorial:", factorial(70))
}

func factorial(x uint64) uint64 {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func average(xs ...int) int {
	var total int = 0
	for _, v := range xs {
		total += v
	}

	return total / len(xs)
}
