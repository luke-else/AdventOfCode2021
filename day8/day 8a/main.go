package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	content := returnContent("../input")
	//content := returnContent("testInput")

	matches := 0

	for _, line := range *content {
		split := strings.Split(line, " | ")
		rhs := split[1]
		digits := strings.Fields(rhs)

		for _, digit := range digits {
			segments := len(digit)
			if segments == 2 || segments == 4 || segments == 3 || segments == 7 {
				matches++
			}
		}
	}

	fmt.Println(matches)

}

func returnContent(path string) *[]string {
	//read file and return it as an array of integers

	file, err := os.Open(path)
	var content []string

	if err != nil {
		fmt.Println("Unlucky, the file didn't open")
		return &content
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	return &content
}
