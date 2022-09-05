package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var elems, lookingFor int

	fmt.Print("Number of elements? ")
	fmt.Scan(&elems)

	var array = make([]int, elems)

	start := time.Now()
	for i := 0; i < elems; i++ {
		array[i] = i
	}
	duration := time.Since(start)
	fmt.Print("Time of generation: ")
	fmt.Println(duration)

	fmt.Println()

	fmt.Println("What are we looking for:")
	fmt.Scan(&lookingFor)

	var binSearch, binResult = binarySearch(&array, lookingFor)
	if binResult {
		fmt.Println("Found: " + strconv.Itoa(binSearch))
	} else {
		fmt.Println("Nothing found!")
	}
	var simpleSearch, simpleResult = simpleSearch(&array, lookingFor)
	if simpleResult {
		fmt.Println("Found: " + strconv.Itoa(simpleSearch))
	} else {
		fmt.Println("Nothing found!")
	}
}

func binarySearch(array *[]int, lookingFor int) (int, bool) {
	var mid, assumption int

	min := 0
	high := len(*array) - 1

	count := 0
	start := time.Now()
	for min <= high {
		count++
		mid = (min + high) / 2
		assumption = mid
		if assumption == lookingFor {
			duration := time.Since(start)
			fmt.Print("Binary search time: ")
			fmt.Println(duration)
			fmt.Println("Steps: " + strconv.Itoa(count))
			return assumption, true
		}
		if assumption > lookingFor {
			high = mid - 1
		} else {
			min = mid + 1
		}
	}
	duration := time.Since(start)
	fmt.Println("Steps: " + strconv.Itoa(count))
	fmt.Print("Binary search time: ")
	fmt.Println(duration)

	return 0, false
}

func simpleSearch(array *[]int, lookingFor int) (int, bool) {
	count := 0
	start := time.Now()
	for value := range *array {
		count++
		if value == lookingFor {
			duration := time.Since(start)
			fmt.Print("Simple search time: ")
			fmt.Println(duration)
			fmt.Println("Steps: " + strconv.Itoa(count))
			return value, true
		}
	}
	duration := time.Since(start)
	fmt.Println("Steps: " + strconv.Itoa(count))
	fmt.Print("Simple search time: ")
	fmt.Println(duration)

	return 0, false
}
