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

	// for i := 0; i < len(gameMap); i++ {
	// 	fmt.Println(gameMap[i])
	// }
}

func run(content *[][]int, gameMap *[1000][1000]int) int {
	for i := 0; i < len((*content)); i++ {
		fillMap(gameMap, (*content)[i])
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

func fillMap(gameMap *[1000][1000]int, coords []int) {

	coord1 := []int{coords[0], coords[1]}
	coord2 := []int{coords[2], coords[3]}

	if coord1[0] > coord2[0] {
		//swap values so coord1 is far left value
		coord1, coord2 = coord2, coord1
	}

	//4 types of line

	//diagonal up,
	//diagonal down,
	//horizontal
	//vertical  ()

	//Check initially for vertical and horizontal lines

	if coord1[0] == coord2[0] {
		//Continue down vertical line
		for i := int(math.Min(float64(coord1[1]), float64(coord2[1]))); i <= int(math.Max(float64(coord1[1]), float64(coord2[1]))); i++ {
			(*gameMap)[i][coord1[0]]++
		}
	} else if coord1[1] == coord2[1] {
		//fill horizontal line in
		for i := coord1[0]; i <= coord2[0]; i++ {
			(*gameMap)[coord1[1]][i]++
		}
	} else {
		//Now check for diagonal lines

		if coord1[1] <= coord2[1] {
			//Diagonal up and to the right
			height := coord1[1]

			for i := coord1[0]; i <= coord2[0]; i++ {
				(*gameMap)[height][i]++
				height++
			}
		} else {
			//Diagonal down and to the right
			height := coord1[1]

			for i := coord1[0]; i <= coord2[0]; i++ {
				(*gameMap)[height][i]++
				height--
			}
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
