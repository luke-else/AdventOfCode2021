package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content := returnContent("../input")
	//content := returnContent("testInput")
	boards, nums := loadBoards(content)
	fmt.Println(run(boards, nums))
}

func run(boards *[]Board, nums []int) int {
	var bingo int
	var n int
	for i := 0; i < len(nums); i++ {
		bingo = callNumber(nums[i], boards)
		if bingo != -1 {
			fmt.Println("found value", bingo, nums[i], (*boards)[bingo])
			n = nums[i]
			break
		}
	}

	//return the answer
	return returnAnswer(&(*boards)[bingo], n)
}

func loadBoards(content *[]string) (boards *[]Board, nums []int) {
	boards = new([]Board)
	newBoard := Board{
		Hash: make(map[int]location, 25),
	}
	boardNum := 0
	row := 0

	for i := 0; i < len(*content); i++ {

		if i == 0 {
			//add nums
			numlist := strings.Split((*content)[i], ",")

			for i := 0; i < len(numlist); i++ {
				num, _ := strconv.Atoi(numlist[i])
				nums = append(nums, num)
			}

			i++
		} else {

			if (*content)[i] != "" {

				regex := regexp.MustCompile("([0-9]+)")
				values := regex.FindAllString((*content)[i], 5)

				for j := 0; j < len(values); j++ {
					value, _ := strconv.Atoi(values[j])
					newBoard.Values[row][j] = boardValue{Value: value, Visited: false}
					newBoard.Hash[value] = location{x: row, y: j}
				}
				row++

				if row == 5 {
					boardNum++
					i++
					*boards = append(*boards, newBoard)
					newBoard = Board{
						Hash: make(map[int]location, 25),
					}
					row = 0
				}

			}

		}

	}
	return
}

func callNumber(n int, boards *[]Board) int {
	for i := 0; i < len(*boards); i++ {
		location, present := (*boards)[i].Hash[n]
		if present {
			//Change the value to visited
			(*boards)[i].Values[location.x][location.y].Visited = true
			if checkBingo(&(*boards)[i], location.x, location.y) {
				return i
			}
		}
	}
	return -1
}

func checkBingo(board *Board, x int, y int) bool {
	checkVertical := true
	//check if bingo for a given row and column

	//Check row
	for i := 0; i < 5; i++ {
		if !(*board).Values[x][i].Visited {
			checkVertical = false
			break
		}
	}

	checkHorizontal := true
	//Check column
	for i := 0; i < 5; i++ {
		if !(*board).Values[i][y].Visited {
			checkHorizontal = false
			break
		}
	}

	return checkVertical || checkHorizontal
}

func returnAnswer(board *Board, n int) (answer int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !(*board).Values[i][j].Visited {
				answer += (*board).Values[i][j].Value
			}
		}
	}
	answer = answer * n
	return
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
