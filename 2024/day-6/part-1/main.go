package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func fileToString(pathToFile string) string {
	file, err := os.ReadFile(pathToFile)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func strToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num

}

const (
	UP    = iota
	RIGHT = iota
	DOWN  = iota
	LEFT  = iota
)

type Guard struct {
	Position     Position
	Direction    int
	TilesVisited map[[2]int]int
}

type Position struct {
	x, y int
}

func (g *Guard) Initialize(world World) {
	g.TilesVisited = make(map[[2]int]int)
	for i := range world.Map {
		for j := range world.Map[i] {
			if world.Map[i][j] == "^" {
				g.Position.x = j
				g.Position.y = i
				g.Direction = UP
			}
		}
	}
}

func (g *Guard) MoveForward(world World) {
	switch g.Direction {
	case UP:
		if g.Position.y == 0 || g.IsNotObstacle(world) {
			g.Position.y -= 1
		} else {
			g.Direction = RIGHT
			g.Position.x += 1
		}
	case RIGHT:
		if g.Position.x == len(world.Map[0]) || g.IsNotObstacle(world) {
			g.Position.x += 1
		} else {
			g.Direction = DOWN
			g.Position.y += 1
		}
	case DOWN:
		if g.Position.y == len(world.Map) || g.IsNotObstacle(world) {
			g.Position.y += 1
		} else {
			g.Direction = LEFT
			g.Position.x -= 1
		}
	case LEFT:
		if g.Position.x == 0 || g.IsNotObstacle(world) {
			g.Position.x -= 1
		} else {
			g.Direction = UP
			g.Position.y -= 1
		}
	}
}

func (g *Guard) isStillInWorld(world World) bool {
	position := g.Position
	if position.y < 0 || position.x < 0 || position.y > len(world.Map)-1 || position.x > len(world.Map[0])-1 {
		return false
	}
	return true
}

func (g *Guard) IsNotObstacle(world World) bool {
	position := g.Position
	if g.isStillInWorld(world) {
		switch g.Direction {
		case UP:
			if !slices.Contains(world.ObstaclesPosition, Position{x: position.x, y: position.y - 1}) {
				return true
			}
		case RIGHT:
			if !slices.Contains(world.ObstaclesPosition, Position{x: position.x + 1, y: position.y}) {
				return true
			}
		case DOWN:
			if !slices.Contains(world.ObstaclesPosition, Position{x: position.x, y: position.y + 1}) {
				return true
			}
		case LEFT:
			if !slices.Contains(world.ObstaclesPosition, Position{x: position.x - 1, y: position.y}) {
				return true
			}
		}
	}
	return false
}

type World struct {
	Map               [][]string
	ObstaclesPosition []Position
}

func (w *World) Initialize(data string) {
	lines := strings.Split(data, "\n")
	for _, rows := range lines {
		columns := strings.Split(rows, "")
		w.Map = append(w.Map, columns)
	}
	w.FindAll()
}

func (w *World) FindAll() {
	for i := range w.Map {
		for j := range w.Map[i] {
			if w.Map[i][j] == "#" {
				w.ObstaclesPosition = append(w.ObstaclesPosition, Position{x: j, y: i})
			}
		}
	}

}

func main() {
	data := fileToString("../input.txt")
	// 	data := `....#.....
	// .........#
	// ..........
	// ..#.......
	// .......#..
	// ..........
	// .#..^.....
	// ........#.
	// #.........
	// ......#...`

	var world World
	world.Initialize(data)

	var guard Guard
	guard.Initialize(world)
	for guard.isStillInWorld(world) {
		guard.MoveForward(world)
		if !guard.isStillInWorld(world) {
			break
		}
		guard.TilesVisited[[2]int{guard.Position.y, guard.Position.x}] += 1
	}

	fmt.Println(len(guard.TilesVisited))
}
