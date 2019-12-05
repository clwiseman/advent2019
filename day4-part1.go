package main

import (
	"fmt"
	"strconv"
	"strings"
)

var minValue = 137683
var maxValue = 596253

func hasDouble(num int) bool {
	// convert to string
	str := strconv.Itoa(num)
	// split digits into array
	array := make([]string, len(str))
	array = strings.Split(str, "")

	// check if digit is same as next
	for i := range array {
		if i == len(array)-1 {
			break
		}
		if array[i] == array[i+1] {
			return true
		}
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
	hasDouble := hasDouble(num)
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
