package math

import (
	"fmt"
	"os"
	"strconv"
)

func Add(numbers []string) int {
	var total int = 0
	for _, number := range numbers {
		num := ConvertToInt(number)
		total += num
	}
	return total
}

func Sub(numbers []string) int {
	var total int = ConvertToInt(numbers[0])
	for _, number := range numbers[1:] {
		num := ConvertToInt(number)
		total -= num
	}
	return total
}

func Multiply(numbers []string) int {
	var total int = 1
	for _, number := range numbers {
		num := ConvertToInt(number)
		total *= num
	}
	return total
}

func Division(numbers []string) int {
	var total int = ConvertToInt(numbers[0])
	//go through the elements in numbers array, starting from index 1
	for _, number := range numbers[1:] {
		num := ConvertToInt(number)
		total = total / num
	}
	return total
}

func ConvertToInt(val string) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println("Couldn't add a number with a characters.")
		os.Exit(1)
	}
	return num
}

func BinarySearch(target int, numbers []string) (int, int) {
	low := 0
	high := len(numbers)
	for high > low {
		mid := low + (high-low)/2
		if ConvertToInt(numbers[mid]) == target {
			return ConvertToInt(numbers[mid]), mid
		} else if ConvertToInt(numbers[mid]) > target {
			high = mid + 1
		} else {
			low = mid - 1
		}
	}
	return -1, -1
}
