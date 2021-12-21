package shared

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
