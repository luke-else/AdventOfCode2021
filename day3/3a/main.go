package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	content := returnContent("../input")
	//content := returnContent("testInput")

	gamma, epsilon := findGammaAndEpsilon(content)

	fmt.Println(binaryToInteger(gamma) * binaryToInteger(epsilon))

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

func binaryToInteger(input string) (value int) {
	n := 0
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == '1' {
			value += (int(math.Pow(float64(2), float64(n))))
		}
		n++
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
