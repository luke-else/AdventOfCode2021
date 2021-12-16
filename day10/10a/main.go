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

	var bracketMap map[string]string = make(map[string]string)
	var bracketCost map[string]int = make(map[string]int)

	answer := 0

	bracketMap[")"] = "("
	bracketMap["]"] = "["
	bracketMap["}"] = "{"
	bracketMap[">"] = "<"

	bracketCost[")"] = 3
	bracketCost["]"] = 57
	bracketCost["}"] = 1197
	bracketCost[">"] = 25137

	for _, row := range *content {
		var stack shared.Stack = shared.Stack{}
		for _, char := range row {
			brack, found := bracketMap[string(char)]

			if !found {
				//If it is an opening bracket
				stack.Push(string(char))
			} else {
				if brack == stack.Peek().Value {
					stack.Pop()
				} else {
					answer += bracketCost[string(char)]
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
