package main

import (
	"AdventOfCode2021/shared"
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	//content := returnContent("../input")
	content := returnContent("../testInput")

	binary := shared.HexToBinary(content)

	pointer := 0
	version := 0
	typeID := 0
	lengthTypeID := "0"
	length := 11
	values := []int{}

	answer := 0

	for pointer < len(binary)-6 {
		//Get version from first 3 Bits
		current := ""
		current = binary[pointer : pointer+3]
		version += shared.BinaryToInteger(&current)
		pointer += 3

		//determine packet type ID from next 2 bits
		current = ""
		current = binary[pointer : pointer+3]
		typeID = shared.BinaryToInteger(&current)
		pointer += 3

		if typeID == 4 {
			//literal value
			value := ""
			for {
				//continue adding the 4 bit values while the packets continue
				pointer++
				value += binary[pointer : pointer+4]
				pointer += 4
				if binary[pointer-5] == '0' {
					break
				}
			}
			answer = shared.BinaryToInteger(&value)

		} else {
			//operator value
			//find the lenth of the subpackets definition
			lengthTypeID = string(binary[pointer])

			if lengthTypeID == "0" {
				length = 15
			} //default = 11
			pointer++

			//subpacket length value
			temp := binary[pointer : pointer+length]
			length = shared.BinaryToInteger(&temp)
			fmt.Println(length, lengthTypeID)

			pointer += length

			//assert into values based on typeID (0 = total length, 1 = )
			if lengthTypeID == "1" {
				//add the series of 11 bit numbers
				for i := 0; i < length; i++ {
					packetString := binary[pointer : pointer+11]
					values = append(values, shared.BinaryToInteger(&packetString))
					pointer += 11
				}
				fmt.Println(values)
			} else {
				fmt.Println("here")
				//add the series of ... bit numbers
			}

		}
	}

	fmt.Println(answer)

}

func ComputePackets(version int, typeID int, values []int) (answer int) {
	answer = 0
	switch typeID {
	case 0:
		//sum
		for _, value := range values {
			answer += value
		}
	case 1:
		//product
		for _, value := range values {
			answer = answer * value
		}
	case 2:
		//min
		min := math.MaxInt
		for _, value := range values {
			if value < min {
				min = value
			}
		}
	case 3:
		//max
		for _, value := range values {
			if value > answer {
				answer = value
			}
		}
	case 5:
		//greater than
		if values[0] > values[1] {
			answer = 1
		}
	case 6:
		//less than
		if values[0] < values[1] {
			answer = 1
		}
	case 7:
		//equal to
		if values[0] == values[1] {
			answer = 1
		}
	}

	return
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
