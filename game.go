package main

import "fmt"

var (
	boardRow  int  = 10
	boardCol  int  = 10
	debugMode bool = false
)

type Game struct {
	board      [][]string
	p1MoveNext bool
}

type Point struct {
	x int
	y int
}

func New() *Game {
	board := make([][]string, boardRow)

	for row := range board {
		singleRow := make([]string, boardCol)
		for col := range singleRow {
			// singleRow[col] = fmt.Sprintf("[%v, %v]", row, col)
			singleRow[col] = "_"
			//fmt.Print(singleRow[col] + " ")
		}
		board[row] = singleRow
		//fmt.Println(board[row])
		//fmt.Printf("\n")
	}
	return &Game{
		board:      board,
		p1MoveNext: true,
	}
}

func (g *Game) PlayerMove(point Point) {
	if g.p1MoveNext {

	}
	if debugMode {
		g.printBoard()
	}
}

func (g *Game) printBoard() {
	for row := range g.board {
		fmt.Println(g.board[row])
	}
}
