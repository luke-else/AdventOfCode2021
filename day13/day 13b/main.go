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
	//content := returnContent("testInput")

	sheet := make(map[Coordinate]bool)
	//var answer string

	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			sheet[Coordinate{X: x, Y: y}] = false
		}
	}

	for _, line := range *content {

		if strings.HasPrefix(line, "fold") {
			//Fold instructions
			instruction := strings.Split(line, "=")

			if strings.Contains(instruction[0], "x") {
				foldPoint, _ := strconv.Atoi(instruction[1])
				sheet = FoldX(sheet, foldPoint)
			}

			if strings.Contains(instruction[0], "y") {
				foldPoint, _ := strconv.Atoi(instruction[1])
				sheet = FoldY(sheet, foldPoint)
			}

		} else if line != "" {
			//mapping instructions
			coordinates := strings.Split(line, ",")
			x, _ := strconv.Atoi(coordinates[0])
			y, _ := strconv.Atoi(coordinates[1])

			sheet[Coordinate{X: x, Y: y}] = true

		}
	}

	for y := 0; y < 8; y++ {
		for x := 0; x < 200; x++ {
			if sheet[Coordinate{X: x, Y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}

}

func FoldX(sheet map[Coordinate]bool, foldPoint int) (folded map[Coordinate]bool) {
	folded = make(map[Coordinate]bool)
	for mark := range sheet {
		x := mark.X
		if x > foldPoint {
			//If the value is in the region that gets folded
			x = 2*foldPoint - x
		}

		folded[Coordinate{X: x, Y: mark.Y}] = true
	}
	return
}

func FoldY(sheet map[Coordinate]bool, foldPoint int) (folded map[Coordinate]bool) {
	folded = make(map[Coordinate]bool)
	for mark := range sheet {
		y := mark.X
		if y > foldPoint {
			//If the value is in the region that gets folded
			y = 2*foldPoint - y
		}

		folded[Coordinate{X: mark.X, Y: y}] = true
	}
	return
}

type Coordinate struct {
	X int
	Y int
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
