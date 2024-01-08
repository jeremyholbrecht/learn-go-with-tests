package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbers []int, numbers2 []int) []int {
	return nil
}

func foo(nums ...int) {
	fmt.Println(nums, "")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
	fmt.Println("len:", len(nums))
}

func main() {
	bar := []int{1, 2, 3, 4}
	foo(bar...)
}
