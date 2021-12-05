package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	content := returnContent("../input")
	//content := returnContent("testInput")

	gameMap := [1000][1000]int{}

	fmt.Println(run(content, &gameMap))
}

func run(content *[][]int, gameMap *[1000][1000]int) int {
	for i := 0; i < len((*content)); i++ {
		fillMap(gameMap, (*content)[i][0], (*content)[i][1], (*content)[i][2], (*content)[i][3])
	}

	count := 0

	for i := 0; i < len(gameMap); i++ {
		for j := 0; j < len(gameMap[i]); j++ {
			if gameMap[i][j] > 1 {
				count++
			}
		}
	}
	return count
}

func fillMap(gameMap *[1000][1000]int, x1 int, y1 int, x2 int, y2 int) {

	if x1 == x2 {
		for i := int(math.Min(float64(y1), float64(y2))); i <= int(math.Max(float64(y1), float64(y2))); i++ {
			(*gameMap)[i][x1]++
		}
	} else if y1 == y2 {
		//fill horizontal line
		for i := int(math.Min(float64(x1), float64(x2))); i <= int(math.Max(float64(x1), float64(x2))); i++ {
			(*gameMap)[y1][i]++
		}
	}
}

func returnContent(path string) *[][]int {
	//read file and return it as an array of integers

	file, err := os.Open(path)
	var content [][]int

	if err != nil {
		fmt.Println("Unlucky, the file didn't open")
		return &content
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	regex := regexp.MustCompile("([0-9]+)")
	for scanner.Scan() {
		stringValues := regex.FindAllString(scanner.Text(), 4)
		intValues := []int{}
		for _, v := range stringValues {
			num, _ := strconv.Atoi(v)
			intValues = append(intValues, num)
		}
		content = append(content, intValues)
	}

	return &content
}
