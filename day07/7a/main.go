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

	list := shared.MergeSort((*content), 0, len(*content)-1)

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
