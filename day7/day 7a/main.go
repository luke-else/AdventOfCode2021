package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	content := returnContent("../input")
	//content := returnContent("testInput")

	list := mergeSort((*content), 0, len(*content)-1)

	position := list[len(list)/2]
	var cost float64

	for i := 0; i < len(list); i++ {
		cost += math.Abs(float64(position) - float64(list[i]))
	}

	fmt.Println(cost)
}

func returnContent(path string) *[]int {
	//read file and return it as an array of integers

	file, err := os.Open(path)
	var content []int

	if err != nil {
		fmt.Println("Unlucky, the file didn't open")
		return &content
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ",")
		for _, v := range text {
			num, _ := strconv.Atoi(v)
			content = append(content, num)
		}
	}

	return &content
}

//sorting algorithm

func mergeSort(nums []int, start int, end int) []int {
	if start == end {
		return []int{nums[start]}
	}

	var mid int = ((end - start) / 2) + start

	//Assign values back into Left and right
	left := mergeSort(nums, start, mid)
	right := mergeSort(nums, mid+1, end)

	var combined []int

	//Pointers for new array
	leftPointer, rightPointer := 0, 0

	for leftPointer <= len(left)-1 || rightPointer <= len(right)-1 {

		if leftPointer == len(left) {
			addValue(&combined, right[rightPointer], &rightPointer)
		} else if rightPointer == len(right) {
			addValue(&combined, left[leftPointer], &leftPointer)
		} else {
			if left[leftPointer] <= right[rightPointer] {
				addValue(&combined, left[leftPointer], &leftPointer)
			} else {
				addValue(&combined, right[rightPointer], &rightPointer)
			}
		}
	}
	return combined
}

func addValue(nums *[]int, value int, pointer *int) {
	*nums = append(*nums, value)
	*pointer++
}
