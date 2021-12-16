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
	var reverseBracketMap map[string]string = make(map[string]string)
	var bracketCost map[string]int = make(map[string]int)

	bracketMap[")"] = "("
	bracketMap["]"] = "["
	bracketMap["}"] = "{"
	bracketMap[">"] = "<"

	reverseBracketMap["("] = ")"
	reverseBracketMap["["] = "]"
	reverseBracketMap["{"] = "}"
	reverseBracketMap["<"] = ">"

	bracketCost[")"] = 1
	bracketCost["]"] = 2
	bracketCost["}"] = 3
	bracketCost[">"] = 4

	//reduce list to required incomplete set

	stackList := []shared.Stack{}
	autoCompletes := []int{}

	for _, row := range *content {

		stack := shared.Stack{}
		keep := true

		for _, char := range row {
			brack, found := bracketMap[string(char)]

			if !found {
				//If it is an opening bracket
				stack.Push(string(char))
			} else {
				if brack == stack.Peek().Value {
					stack.Pop()
				} else {
					keep = false
					// *content = *removeElement(content, i)
					break
				}
			}
		}

		if keep {
			stackList = append(stackList, stack)
		}

	}

	emptyStackNode := shared.Node{}

	for _, stack := range stackList {
		autocomplete := 0
		for *(stack.Peek()) != emptyStackNode {
			autocomplete *= 5
			autocomplete += bracketCost[reverseBracketMap[stack.Pop()]]
		}
		autoCompletes = append(autoCompletes, autocomplete)
	}

	autoCompletes = shared.MergeSort(autoCompletes, 0, len(autoCompletes)-1)

	mid := autoCompletes[len(autoCompletes)/2]

	fmt.Println(mid)

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
