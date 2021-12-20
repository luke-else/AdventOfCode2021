package main

import (
	"AdventOfCode2021/shared"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//content := returnContent("../input")
	content := returnContent("../testInput")

	fmt.Println(*content)

	key := (*content)[0]

	input := [][]int{}

	for i := 2; i < len(*content); i++ {
		input = append(input, (*content)[i])
	}

	for x := 1; x < len(input)-1; x++ {
		for y := 1; y < len(input[0]); y++ {
			binary := GetSurroundingValue(&input, x, y)
			newValue := shared.BinaryToInteger(&binary)

			fmt.Println(newValue, key[newValue])

		}
	}

}

func GetSurroundingValue(input *[][]int, xVal int, yVal int) (binary string) {
	for x := xVal - 1; x <= xVal+1; x++ {
		for y := yVal - 1; y <= yVal+1; y++ {
			binary += strconv.Itoa((*input)[x][y])
		}
	}
	return
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
		line := scanner.Text()

		nums := []int{}

		for i := 0; i < len(line); i++ {
			num := 0

			if line[i] == '#' {
				num = 1
			}

			nums = append(nums, num)
		}
		content = append(content, nums)
	}

	return &content
}
