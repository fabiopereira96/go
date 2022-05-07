package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1:]
	numbers := make([]int, len(input))

	for index, n := range input {
		number, err := strconv.Atoi(n)

		if err != nil {
			fmt.Printf("%s não é um número válido!\n", n)
			os.Exit(1)
		}
		numbers[index] = number
	}

	fmt.Println(quicksort(numbers))
}

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	n := make([]int, len(numbers))
	copy(n, numbers)

	pivotIndex := len(n) / 2
	pivot := n[pivotIndex]

	n = append(n[:pivotIndex], n[pivotIndex+1:]...)

	less, greather := partition(n, pivot)

	return append(append(quicksort(less), pivot), quicksort(greather)...)
}

func partition(numbers []int, pivot int) (less []int, greather []int) {
	for _, n := range numbers {
		if n <= pivot {
			less = append(less, n)
		} else {
			greather = append(greather, n)
		}
	}
	return less, greather
}
