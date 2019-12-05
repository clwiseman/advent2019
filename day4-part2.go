package main

import (
	"fmt"
	"strconv"
	"strings"
)

var minValue = 137683
var maxValue = 596253

func hasDoubleButNotMore(num int) bool {
	// convert to string
	str := strconv.Itoa(num)
	// split digits into array
	array := make([]string, len(str))
	array = strings.Split(str, "")

	// check if three digits in a row, fail
	for i := range array {
		if i == len(array)-2 {
			break
		}
		if array[i] == array[i+1] && array[i] == array[i+2] {
			return false
		}
	}

	// if not three in a row are there two in row, pass
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
