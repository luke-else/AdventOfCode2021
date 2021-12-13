package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	content := returnContent("../input")
	//content := returnContent("testInput")

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
