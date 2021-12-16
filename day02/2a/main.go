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

	horizontal, depth := followGuidance(content)

	fmt.Println(horizontal * depth)

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

func followGuidance(content *[]string) (horizontal int, depth int) {

	for i := 0; i < len(*content); i++ {
		currentLine := strings.Split((*content)[i], " ")

		value, err := strconv.Atoi(currentLine[1])

		if err != nil {
			fmt.Println("Uh oh, couldn't find the key! There was an issue with the sonar")
		}

		switch currentLine[0] {
		//instructions say that up and down are reversed as we are in the submarine
		case "up":
			depth -= value
		case "down":
			depth += value
		default:
			horizontal += value
		}
	}
	return

}
