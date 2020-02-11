package main

import (
	"fmt"
)

func binarysearch(data []int, val int) int {
	var mid int
	var low int
	high := len(data) - 1
	for low <= high {
		mid = low + (high-low)/2
		if data[mid] == val {
			return mid
		}
		if data[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return -1

}

//array includes positive and negetive numbers
//sum must be positive cause if at least positive number exists, sum is positive
func findMaxSumSubarray(data []int) int {
	pureSum := 0
	x := 0
	for i := 0; i < len(data); i++ {
		x = x + data[i]
		if x < 0 { //from here subarray finish
			x = 0
		}
		if pureSum < x {
			pureSum = x
		}
	}
	return pureSum
}

//euclid's algorithm is gcd(m,n) = gcd(n,m%n)
func greatestCommonDivisor(x int, y int) int {
	if x < y {
		greatestCommonDivisor(y, x)
	}
	if x%y == 0 {
		return y
	}
	return greatestCommonDivisor(y, x%y)
}

func printArray(data []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()
}

func allPermutationOfAnIntegerArray(data []int, x int, size int) {
	if x == size-1 {
		printArray(data, size)
		return
	}

	for counter := x; counter < size; counter++ {
		data[x], data[counter] = data[counter], data[x]
		allPermutationOfAnIntegerArray(data, x+1, size)
		data[x], data[counter] = data[counter], data[x]
	}
}

func main() {
	var data [3]int
	for i := 0; i < len(data); i++ {
		data[i] = i
	}
	allPermutationOfAnIntegerArray(data[:], 0, len(data))
	fmt.Println(greatestCommonDivisor(12, 9))

}
