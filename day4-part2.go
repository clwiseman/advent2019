package main

import (
	"fmt"
	"strconv"
	"strings"
)

var minValue = 137683
var maxValue = 596253

func indexOf(element string, data []string) int {
	for i, v := range data {
		if element == v {
			return i
		}
	}
	return -1 //not found.
}

func lastIndexOf(element string, data []string) int {
	indices := make([]int, len(data))
	for i, v := range data {
		if element == v {
			indices = append(indices, i)
		}
	}
	if len(indices) > 0 {
		return indices[len(indices)-1]
	}
	return -1 //not found.
}

func hasDoubleButNotMore(num int) bool {
	// convert to string
	str := strconv.Itoa(num)
	// split digits into array
	array := make([]string, len(str))
	array = strings.Split(str, "")

	// empty array to hold duplicates
	duplicates := make([]string, len(str))

	// check if same digit in array, push to empty array
	for i, v := range array {
		if i == len(array)-1 {
			break
		}
		//append doubles to duplicates array
		if array[i] == array[i+1] {
			duplicates = append(duplicates, v)
		}
	}
	//count if single digit in array
	for _, v := range duplicates {
		if indexOf(v, duplicates) != lastIndexOf(v, duplicates) {
			continue
		}
		return true
	}

	return false
}

func noDecrease(num int) bool {
	// convert to string
	str := strconv.Itoa(num)
	// split digits into array
	array := make([]string, len(str))
	array = strings.Split(str, "")

	// check if digit is greater than next
	for i := range array {
		if i == len(array)-1 {
			break
		}
		if array[i] > array[i+1] {
			return false
		}
	}
	return true
}

func isPassword(num int) bool {
	hasDouble := hasDoubleButNotMore(num)
	noDecrease := noDecrease(num)

	if hasDouble && noDecrease {
		return true
	}
	return false
}

func howManyPasswords(min int, max int) int {
	count := 0
	for i := min; i <= max; i++ {
		if isPassword(i) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(howManyPasswords(minValue, maxValue))
}
