package main

import (
	"AdventOfCode2021/shared"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	content := returnContent("../input")
	//content := returnContent("../testInput")

	coordLeft := shared.Coordinate{
		X: shared.Min((*content)[0], (*content)[1]),
		Y: shared.Max((*content)[2], (*content)[3]),
	}
	coordRight := shared.Coordinate{
		X: shared.Max((*content)[0], (*content)[1]),
		Y: shared.Min((*content)[2], (*content)[3]),
	}

	answer := make(map[shared.Coordinate]bool)

	y := 0

	for {
		for x := 1; x <= coordRight.X; x++ {
			//Test with positive Y velocity
			probe := Probe{
				Position: &shared.Coordinate{X: 0, Y: 0},
				Velocity: &shared.Coordinate{X: x, Y: y},
			}
			success, _ := probe.Model(&coordLeft, &coordRight)
			if success {
				answer[shared.Coordinate{X: x, Y: y}] = true
			}

			//Test with negative Y velocity
			probe = Probe{
				Position: &shared.Coordinate{X: 0, Y: 0},
				Velocity: &shared.Coordinate{X: x, Y: -y},
			}
			success, _ = probe.Model(&coordLeft, &coordRight)
			if success {
				answer[shared.Coordinate{X: x, Y: -y}] = true
			}
		}
		fmt.Println(len(answer))
		y++
	}

}

type Probe struct {
	Position *shared.Coordinate
	Velocity *shared.Coordinate
}

func (p *Probe) Step() {
	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y

	if p.Velocity.X > 0 {
		p.Velocity.X--
	} else if p.Velocity.X < 0 {
		p.Velocity.X++
	}

	p.Velocity.Y--
}

func (p *Probe) Model(left *shared.Coordinate, right *shared.Coordinate) (success bool, maxY int) {
	maxY = 0
	for p.Position.X <= right.X && p.Position.Y >= right.Y {
		p.Step()
		if p.Position.Y > maxY {
			maxY = p.Position.Y
		}
		if CheckLocation(p, left, right) {
			return true, maxY
		}
	}
	return false, maxY
}

func CheckLocation(probe *Probe, left *shared.Coordinate, right *shared.Coordinate) bool {
	if probe.Position.X >= left.X && probe.Position.X <= right.X {
		if probe.Position.Y <= left.Y && probe.Position.Y >= right.Y {
			return true
		}
	}
	return false
}

func returnContent(path string) *[]int {
	//read file and return it as an array of integers

	file, err := os.Open(path)
	var content []int

	if err != nil {
		fmt.Println("Unlucky, the file didn't open")
		return &content
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	regex, _ := regexp.Compile(`[-+]?[\d]+`)

	strings := []string{}

	for scanner.Scan() {
		strings = regex.FindAllString(scanner.Text(), 4)
	}

	for _, val := range strings {
		num, _ := strconv.Atoi(val)
		content = append(content, num)
	}

	return &content
}
