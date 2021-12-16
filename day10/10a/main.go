package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	content := returnContent("../input")
	//content := returnContent("../testInput")

	var bracketMap map[byte]byte = make(map[byte]byte)
	var bracketCost map[byte]int = make(map[byte]int)

	answer := 0

	bracketMap[')'] = '('
	bracketMap[']'] = '['
	bracketMap['}'] = '{'
	bracketMap['>'] = '<'

	bracketCost[')'] = 3
	bracketCost[']'] = 57
	bracketCost['}'] = 1197
	bracketCost['>'] = 25137

	for _, row := range *content {
		var stack Stack = Stack{}
		for _, char := range row {
			brack, found := bracketMap[char]

			if !found {
				//If it is an opening bracket
				stack.Push(char)
			} else {
				if brack == stack.Peek().char {
					stack.Pop()
				} else {
					answer += bracketCost[char]
					fmt.Println("Illegal,", char, "found", answer)
					break
				}
			}
		}
	}

	fmt.Println(answer)

}

func returnContent(path string) *[][]byte {
	//read file and return it as an array of integers

	file, err := os.Open(path)
	var content [][]byte

	if err != nil {
		fmt.Println("Unlucky, the file didn't open")
		return &content
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content = append(content, []byte(scanner.Text()))
	}

	return &content
}
