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
	fmt.Println(content)

	days := 80

	for i := 0; i < days; i++ {
		for j := 0; j < len(*content); j++ {
			(*content)[j]--
			if (*content)[j] < 0 {
				(*content)[j] = 6
				(*content) = append((*content), 9)
			}
		}
	}

	fmt.Println(len(*content))
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
