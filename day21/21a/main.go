package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	content := returnContent("../input")
	// content := returnContent("../testInput")

	//Construct Game board... in reverse :) (Just makes it easier for stringing them together)
	positions := make(map[int]*GamePosition)
	gameBoard := GameBoard{
		Start: &GamePosition{
			Position: 1,
			Next:     nil,
		},
	}
	positions[1] = gameBoard.Start
	next := gameBoard.Start
	var currentPosition *GamePosition

	for i := 10; i > 1; i-- {
		currentPosition = new(GamePosition)
		currentPosition.Position = i
		currentPosition.Next = next
		next = currentPosition

		positions[i] = currentPosition
	}
	gameBoard.Start.Next = currentPosition

	//Add Players to the game
	players := make(map[int]*Player)
	numPlayers := 0

	for _, value := range *content {
		newPlayer := Player{
			PlayerNum: value[0],
			Score:     0,
			Position:  positions[value[1]],
		}
		players[newPlayer.PlayerNum] = &newPlayer
		numPlayers++
	}

	//Play the GAMEEEEEE
	diceVal := 1
	numRolls := 0
	winner := 0

	for {
		for i := 1; i <= numPlayers; i++ {
			p := players[i]
			if p.Roll(&diceVal, &numRolls) {
				winner = p.PlayerNum
				break
			}
		}

		if winner > 0 {
			break
		}
	}

	loser := 1
	if winner == 1 {
		loser = 2
	}

	fmt.Println(numRolls * players[loser].Score)

}

type GameBoard struct {
	Start *GamePosition
}

type GamePosition struct {
	Position int
	Next     *GamePosition
}

type Player struct {
	PlayerNum int
	Score     int
	Position  *GamePosition
}

func (p *Player) Roll(diceVal *int, numRolls *int) bool {
	//simulate 3 consecutive dice rolls
	dice := (*diceVal + 1) * 3
	*diceVal += 3
	*numRolls += 3

	//calculate where the position will be after n positions being moved
	position := p.Position
	for i := 0; i < dice; i++ {
		position = position.Next
	}

	p.Position = position

	p.Score += p.Position.Position

	//Return bool for if player has won
	return p.Score >= 1000

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

	regex, _ := regexp.Compile(`[0-9]+`)

	for scanner.Scan() {

		line := regex.FindAllString(scanner.Text(), 2)
		nums := []int{}

		for i := 0; i < len(line); i++ {
			num, _ := strconv.Atoi(line[i])
			nums = append(nums, num)
		}
		content = append(content, nums)
	}

	return &content
}
