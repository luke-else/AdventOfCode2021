package main

import (
	"AdventOfCode2021/shared"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	content := returnContent("../input")
	//content := returnContent("../testInput")

	octopuses := map[shared.Coordinate]int{}

	answer := 0

	width := 0
	height := len(*content)

	//Insert octopuses into a map
	for y, row := range *content {
		width = len(row)
		for x, val := range row {
			octopuses[shared.Coordinate{X: x, Y: y}] = val
		}
	}

	for i := 0; i < 100; i++ {
		flashers := map[shared.Coordinate]bool{}

		for coord := range octopuses {
			octopuses[coord]++
		}

		for {
			complete := true

			for coord, energy := range octopuses {
				if energy > 9 && !flashers[coord] {
					complete = false
					flashers[coord] = true
					answer++
					for _, neighbour := range coord.Neighbours(width, height, true) {
						octopuses[neighbour]++
					}
				}
			}

			if complete {
				break
			}
		}

		for coord, flashed := range flashers {
			if flashed {
				octopuses[coord] = 0
			}
		}
	}

	fmt.Println(*content)
	fmt.Println(answer)
}

func returnContent(path string) *[][]int {
	//read file and return it as an array of integers

	file, err := os.Open(path)
	var content [][]int

	if err != nil {
		fmt.Println("Unlucky, the file didn't open")
		return &content
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		nums := []int{}
		for _, char := range scanner.Text() {
			num, _ := strconv.Atoi(string(char))
			nums = append(nums, num)
		}
		content = append(content, nums)
	}

	return &content
}
