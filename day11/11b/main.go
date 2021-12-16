package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	content := returnContent("../input")
	//content := returnContent("../testInput")

	octopuses := map[Coordinate]int{}

	answer := 0

	width := 0
	height := len(*content)

	//Insert octopuses into a map
	for y, row := range *content {
		width = len(row)
		for x, val := range row {
			octopuses[Coordinate{X: x, Y: y}] = val
		}
	}

	i := 0
	for {
		i++
		flashers := map[Coordinate]bool{}
		for coords, energy := range octopuses {
			octopuses[coords] = energy + 1
		}

		for {
			doneFlashing := true

			for coords, energy := range octopuses {
				if energy > 9 && !flashers[coords] {
					doneFlashing = false
					flashers[coords] = true
					for _, neighbour := range coords.Neighbours(width, height, true) {
						octopuses[neighbour]++
					}
				}
			}

			if doneFlashing {
				break
			}
		}

		for coords, flashed := range flashers {
			if flashed {
				octopuses[coords] = 0
			}
		}

		allFlashed := true
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if !flashers[Coordinate{
					X: x,
					Y: y,
				}] {
					allFlashed = false
				}
			}
		}
		if allFlashed {
			answer = i
			break
		}
	}

	//fmt.Println(*content)
	fmt.Println(answer)
}

//Coordinate Class
type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) Neighbours(gridWidth int, gridHeight int, diagonal bool) (out []Coordinate) {
	spaceLeft := c.X > 0
	spaceRight := c.X < gridWidth-1
	spaceUp := c.Y > 0
	spaceDown := c.Y < gridHeight-1

	if spaceLeft {
		out = append(out, Coordinate{c.X - 1, c.Y})
	}
	if spaceRight {
		out = append(out, Coordinate{c.X + 1, c.Y})
	}
	if spaceUp {
		out = append(out, Coordinate{c.X, c.Y - 1})
	}
	if spaceDown {
		out = append(out, Coordinate{c.X, c.Y + 1})
	}

	if diagonal {
		if spaceUp && spaceLeft {
			out = append(out, Coordinate{c.X - 1, c.Y - 1})
		}
		if spaceUp && spaceRight {
			out = append(out, Coordinate{c.X + 1, c.Y - 1})
		}
		if spaceDown && spaceLeft {
			out = append(out, Coordinate{c.X - 1, c.Y + 1})
		}
		if spaceDown && spaceRight {
			out = append(out, Coordinate{c.X + 1, c.Y + 1})
		}
	}

	return
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

	for scanner.Scan() {
		nums := []int{}
		for _, char := range scanner.Text() {
			num, _ := strconv.Atoi(string(char))
			nums = append(nums, num)
		}
		content = append(content, nums)
	}

	return &content
}
