package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	content := returnContent("../input")

	//content := returnContent("../testInput")

	oxygen, carbon := findOxygenAndCarbon(content, 0)

	fmt.Println(oxygen, carbon)

	fmt.Println(binaryToInteger(oxygen) * binaryToInteger(carbon))
}

func findOxygenAndCarbon(content *[]string, i int) (oxygen string, carbon string) {
	//recursion

	oxygen = (*findValues(content, 0, true))[0]

	carbon = (*findValues(content, 0, false))[0]

	return
}

//recursively find the values that fit the criteria
func findValues(content *[]string, i int, inclusive bool) (values *[]string) {

	if len(*content) <= 1 {
		return content
	}

	count := 0
	var bit byte = '0'
	for j := 0; j < len(*content); j++ {
		if ((*content)[j])[i] == '1' {
			count++
		}
	}

	if inclusive {
		if float64(count) >= float64(len(*content))/2 {
			bit = '1'
		}
	} else {
		if float64(count) < float64(len(*content))/2 {
			bit = '1'
		}
	}

	var newContent []string
	for j := 0; j < len(*content); j++ {
		if ((*content)[j])[i] == bit {
			newContent = append(newContent, (*content)[j])
		}
	}
	i++
	return findValues(&newContent, i, inclusive)
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
