package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(countGreater(returnContent("../input")))
}

func countGreater(input *[]int) (count int) {
	//Find how many elements have increased from previous

	for i := 3; i < len(*input); i++ {
		sum1 := (*input)[i-1] + (*input)[i-2] + (*input)[i-3]
		sum2 := (*input)[i-1] + (*input)[i-2] + (*input)[i]
		if sum2 > sum1 {
			count++
		}
	}
	return
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
		value, _ := strconv.Atoi(scanner.Text())
		content = append(content, value)
	}

	return &content
}
