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
	//content := returnContent("../input")
	content := returnContent("../testInput")

	//Construct Game board... in reverse :) (Just makes it easier for stringing them together)
	positions := make(map[int]*shared.GamePosition)
	gameBoard := shared.GameBoard{
		Start: &shared.GamePosition{
			Position: 1,
			Next:     nil,
		},
	}
	positions[1] = gameBoard.Start
	next := gameBoard.Start
	var currentPosition *shared.GamePosition

	for i := 10; i > 1; i-- {
		currentPosition = new(shared.GamePosition)
		currentPosition.Position = i
		currentPosition.Next = next
		next = currentPosition

		positions[i] = currentPosition
	}
	gameBoard.Start.Next = currentPosition

	//Add Players to the game
	players := make(map[int]*shared.Player)
	numPlayers := 0

	for _, value := range *content {
		newPlayer := shared.Player{
			PlayerNum: value[0],
			Score:     0,
			Position:  positions[value[1]],
		}
		players[newPlayer.PlayerNum] = &newPlayer
		numPlayers++
	}

	//Play the GAMEEEEEE
	queue := shared.UniverseQueue{}
	wins := make((map[int]int))
	//Make the initial universe and add to queue

	queue.Add(&shared.Universe{Players: players})

	//Continue to process the universe

	for {
		current := queue.Pop()
		if current == nil {
			break
		}

		for _, p := range current.Players {
			RecurseUniverse(current.Players, &queue, p.PlayerNum, wins)
		}

	}

	fmt.Println(wins[1], wins[2])

}

func RecurseUniverse(players map[int]*shared.Player, queue *shared.UniverseQueue, player int, wins map[int]int) {
	for d := 1; d <= 3; d++ {
		//Create a new instance based on the current dice roll
		alteredPlayer := Roll(*players[player], d)

		if alteredPlayer.Score < 21 {
			newUniverse := new(shared.Universe)

			newPlayers := make(map[int]*shared.Player)
			for _, p := range players {
				newPlayer := *p
				newPlayers[newPlayer.PlayerNum] = &newPlayer
			}
			newPlayers[player] = alteredPlayer
			newUniverse.Players = newPlayers
			queue.Add(newUniverse)
		} else {
			wins[player]++
		}

	}
}

func Roll(p shared.Player, diceVal int) *shared.Player {

	//calculate where the position will be after n positions being moved
	position := p.Position
	for i := 0; i < diceVal; i++ {
		position = position.Next
	}

	p.Position = position

	p.Score += p.Position.Position

	return &p

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
