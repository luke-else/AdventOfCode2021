package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/yourbasic/graph"
)

func main() {
	content := returnContent("../input")
	//content := returnContent("../testInput")

	height := len(*content)
	width := len((*content)[0])

	fullHeight := 5 * height
	fullWidth := 5 * width
	g := graph.New(fullHeight * fullWidth)

	for y := 0; y < fullHeight; y++ {
		for x := 0; x < fullWidth; x++ {
			c := Coordinate{
				X: x,
				Y: y,
			}
			myIdx := c.Y*fullWidth + c.X
			neighbours := c.Neighbours(fullWidth, fullHeight, false)
			for _, neighbour := range neighbours {
				neighbourCost := int64((*content)[neighbour.Y%height][neighbour.X%width])
				widthPenalty := neighbour.X / width
				heightPenalty := neighbour.Y / height
				neighbourCost += int64(widthPenalty + heightPenalty)
				for neighbourCost > 9 {
					neighbourCost -= 9
				}

				neighborIdx := neighbour.Y*fullWidth + neighbour.X
				g.AddCost(myIdx, neighborIdx, neighbourCost)
			}
		}
	}

	_, cost := graph.ShortestPath(g, 0, fullWidth*fullHeight-1)

	fmt.Println(cost)
}

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
		for _, num := range scanner.Text() {
			val, _ := strconv.Atoi(string(num))
			nums = append(nums, val)
		}
		content = append(content, nums)
	}

	return &content
}
