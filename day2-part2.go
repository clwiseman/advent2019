package main

import (
	"fmt"
)

var array = [181]int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 9, 1, 19, 1, 19, 5, 23, 1, 23, 5, 27, 2, 27, 10, 31, 1, 31, 9, 35, 1, 35, 5, 39, 1, 6, 39, 43, 2, 9, 43, 47, 1, 5, 47, 51, 2, 6, 51, 55, 1, 5, 55, 59, 2, 10, 59, 63, 1, 63, 6, 67, 2, 67, 6, 71, 2, 10, 71, 75, 1, 6, 75, 79, 2, 79, 9, 83, 1, 83, 5, 87, 1, 87, 9, 91, 1, 91, 9, 95, 1, 10, 95, 99, 1, 99, 13, 103, 2, 6, 103, 107, 1, 107, 5, 111, 1, 6, 111, 115, 1, 9, 115, 119, 1, 119, 9, 123, 2, 123, 10, 127, 1, 6, 127, 131, 2, 131, 13, 135, 1, 13, 135, 139, 1, 9, 139, 143, 1, 9, 143, 147, 1, 147, 13, 151, 1, 151, 9, 155, 1, 155, 13, 159, 1, 6, 159, 163, 1, 13, 163, 167, 1, 2, 167, 171, 1, 171, 13, 0, 99, 2, 0, 14, 0}

func opCode1(list [181]int, one int, two int, three int) [181]int {
	value := list[one] + list[two]
	list[three] = value
	return list
}

func opCode2(list [181]int, one int, two int, three int) [181]int {
	value := list[one] * list[two]
	list[three] = value
	return list
}

func intCode(array [181]int, noun int, verb int) int {
	//Initialize the memory
	list := array

	//Set the inputs
	list[1] = noun
	list[2] = verb

	//Run the IntCode program
	for i := 0; i <= len(list); i += 4 {
		if i+1 > len(list) || i+2 > len(list) || i+3 > len(list) {
			break
		}
		opCode := list[i]

		if opCode == 1 {
			list = opCode1(list, list[i+1], list[i+2], list[i+3])
		}
		if opCode == 2 {
			list = opCode2(list, list[i+1], list[i+2], list[i+3])
		}
		if opCode == 99 {
			break
		}
	}

	return list[0]
}

func solver(value int, list [181]int) int {
	noun, verb := 0, 0

	for ; noun < 99; noun++ {
		verb = 0

		for ; verb < 99; verb++ {
			if intCode(list, noun, verb) == value {
				return 100*noun + verb
			}
		}
	}
	return 100*noun + verb
}

func main() {
	fmt.Println(solver(19690720, array))
}
