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

	i := 0
	for {
		i++
		flashers := map[shared.Coordinate]bool{}
		for coords, energy := range octopuses {
			octopuses[coords] = energy + 1
		}

		for {
			doneFlashing := true

			for coords, energy := range octopuses {
				if energy > 9 && !flashers[coords] {
					doneFlashing = false
					flashers[coords] = true
					for _, neighbour := range coords.Neighbours(width, height, true) {
						octopuses[neighbour]++
					}
				}
			}

			if doneFlashing {
				break
			}
		}

		for coords, flashed := range flashers {
			if flashed {
				octopuses[coords] = 0
			}
		}

		allFlashed := true
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if !flashers[shared.Coordinate{
					X: x,
					Y: y,
				}] {
					allFlashed = false
				}
			}
		}
		if allFlashed {
			answer = i
			break
		}
	}

	//fmt.Println(*content)
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
