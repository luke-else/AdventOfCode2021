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

type Universe struct {
	Players map[int]*Player
}

type QueueItem struct {
	Universe *Universe
	Next     *QueueItem
}

type UniverseQueue struct {
	Head *QueueItem
	Tail *QueueItem
}

func (q *UniverseQueue) Pop() (universe *Universe) {
	universe = nil
	if q.Head != nil {
		universe = q.Head.Universe
		q.Head = q.Head.Next
		if q.Head == nil {
			q.Tail = nil
		}
	}
	return
}

func (q *UniverseQueue) Add(universe *Universe) {
	newUniverse := new(QueueItem)
	newUniverse.Universe = universe

	if q.Tail != nil {
		q.Tail.Next = newUniverse
	} else {
		q.Head = newUniverse
	}

	q.Tail = newUniverse
}
