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

	binary := shared.HexToBinary(content)

	pointer := 0
	version := 0

	for pointer < len(binary)-6 {
		//Get version from first 3 Bits
		current := ""
		current = binary[pointer : pointer+3]
		version += shared.BinaryToInteger(&current)
		pointer += 3

		//determine packet type ID from next 2 bits
		current = ""
		current = binary[pointer : pointer+3]
		typeID := shared.BinaryToInteger(&current)
		pointer += 3

		if typeID == 4 {
			//literal value
			for binary[pointer] == '1' {
				pointer += 5
			}
			pointer += 5

		} else {
			//operator value
			if binary[pointer] == '1' {
				pointer += 12
			} else {
				pointer += 16
			}
		}
	}

	fmt.Println(version)

}

func returnContent(path string) *string {
	//read file and return it as an array of integers

	file, err := os.Open(path)
	var content string

	if err != nil {
		fmt.Println("Unlucky, the file didn't open")
		return &content
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content = scanner.Text()
	}

	return &content
}
