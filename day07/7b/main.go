package main

import (
	"AdventOfCode2021/shared"
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

	crabs := shared.MergeSort(*content, 0, len(*content)-1)

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
