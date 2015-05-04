package main

import "fmt"

func main() {
	var array []int = []int{1, 2, 3, 4, 5}

	var a [2]string = [2]string{"hello", "world"}
	fmt.Println(a[0], a[1])
	fmt.Println(a)

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

	fmt.Println("average:", average(array...))
}

func average(xs ...int) int {
	var total int = 0
	for _, v := range xs {
		total += v
	}

	return total / len(xs)
}
