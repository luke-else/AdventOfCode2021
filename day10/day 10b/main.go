package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	content := returnContent("../input")
	//content := returnContent("testInput")

	var bracketMap map[byte]byte = make(map[byte]byte)
	var reverseBracketMap map[byte]byte = make(map[byte]byte)
	var bracketCost map[byte]int = make(map[byte]int)

	bracketMap[')'] = '('
	bracketMap[']'] = '['
	bracketMap['}'] = '{'
	bracketMap['>'] = '<'

	reverseBracketMap['('] = ')'
	reverseBracketMap['['] = ']'
	reverseBracketMap['{'] = '}'
	reverseBracketMap['<'] = '>'

	bracketCost[')'] = 1
	bracketCost[']'] = 2
	bracketCost['}'] = 3
	bracketCost['>'] = 4

	//reduce list to required incomplete set

	stackList := []Stack{}
	autoCompletes := []int{}

	for _, row := range *content {

		stack := Stack{}
		keep := true

		for _, char := range row {
			brack, found := bracketMap[char]

			if !found {
				//If it is an opening bracket
				stack.Push(char)
			} else {
				if brack == stack.Peek().char {
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

	emptyStackNode := StackNode{}

	for _, stack := range stackList {
		autocomplete := 0
		for *(stack.Peek()) != emptyStackNode {
			autocomplete *= 5
			autocomplete += bracketCost[reverseBracketMap[stack.Pop()]]
		}
		autoCompletes = append(autoCompletes, autocomplete)
	}

	autoCompletes = mergeSort(autoCompletes, 0, len(autoCompletes)-1)

	mid := autoCompletes[len(autoCompletes)/2]

	fmt.Println(mid)

}

// func removeElement(slice *[][]byte, i int) (new *[][]byte) {
// 	new = &[][]byte{}
// 	for j, row := range *slice {
// 		if i != j {
// 			*new = append(*new, row)
// 		}
// 	}
// 	return
// }

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
