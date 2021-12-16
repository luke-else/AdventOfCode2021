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
	//content := returnContent("../testInput")

	crabs := mergeSort(*content, 0, len(*content)-1)

	min, max := crabs[0], crabs[len(crabs)-1]

	dists := countDists(min, max)
	minDist := 100000000000000000
	for i := min; i <= max; i++ {
		s := 0
		failed := false
		for _, start := range crabs {
			s += dists[int(math.Abs(float64(start-i)))]
			if s > minDist {
				failed = true
				break
			}
		}
		if failed {
			continue
		}
		if s < minDist {
			minDist = s
		}
	}
	fmt.Println(minDist)
}

func countDists(min int, max int) map[int]int {
	dists := make(map[int]int)
	for i := min; i < max; i++ {
		temp := 1
		s := 0
		for j := i; j < max; j++ {
			s += temp
			temp++
		}
		dists[max-i] = s
	}
	return dists
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
