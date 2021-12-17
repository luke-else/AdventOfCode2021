package main

import (
	"AdventOfCode2021/shared"
	"bufio"
	"fmt"
	"os"
)

func main() {
	content := returnContent("../input")
	//content := returnContent("../testInput")

	gamma, epsilon := findGammaAndEpsilon(content)

	fmt.Println(shared.BinaryToInteger(&gamma) * shared.BinaryToInteger(&epsilon))

}

func findGammaAndEpsilon(content *[]string) (gamma string, epsilon string) {
	for i := 0; i < len((*content)[1]); i++ {
		count := 0

		//Loop through list checking index[i] for each string
		for j := 0; j < len(*content); j++ {
			if ((*content)[j])[i] == '1' {
				count++
			}
		}

		if count >= len(*content)/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}

	}
	return
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
