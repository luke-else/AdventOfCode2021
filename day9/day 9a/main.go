package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content := returnContent("../input")
	//content := returnContent("testInput")
	answer := 0

	for y, row := range *content {
		for x, value := range row {

			low := true

			for _, n := range *neighbours(x, y, len((*content)[0]), len(*content)) {
				//Check for each value if the current one we are looking
				//at is greater than any of its neighbours
				if value >= (*content)[n[0]][n[1]] {
					low = false
					break
				}
			}
			if low {
				answer += (1 + value)
			}
		}
	}

	fmt.Println(answer)
}

func neighbours(x int, y int, lenX int, lenY int) *[][]int {
	var neighbours [][]int

	if x > 0 {
		//Cell left
		neighbours = append(neighbours, []int{y, x - 1})
	}
	if x < lenX-1 {
		//Cell right
		neighbours = append(neighbours, []int{y, x + 1})
	}
	if y > 0 {
		//Cell Above
		neighbours = append(neighbours, []int{y - 1, x})
	}
	if y < lenY-1 {
		//Cell Below
		neighbours = append(neighbours, []int{y + 1, x})
	}
	return &neighbours
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
		var values []int
		for _, val := range strings.Split(scanner.Text(), "") {
			num, _ := strconv.Atoi(val)
			values = append(values, num)
		}
		content = append(content, values)
	}

	return &content
}
