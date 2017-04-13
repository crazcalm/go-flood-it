package main

import (
	tl "github.com/JoelOtter/termloop"
)

//Board to the game
type Board struct {
	Hieght int
	Width  int
	Grid []*Clickable
}


//Clickable ... add docs
type Clickable struct {
	*tl.Rectangle
	Grid *Board
}

//Tick ... add docs
func (c *Clickable) Tick (ev tl.Event) {
	x, y := c.Position()
	if ev.Type == tl.EventMouse && ev.MouseX == x && ev.MouseY == y {
		c.SetColor(tl.ColorRed)
	}
}

//NewClickable ... add docs
func NewClickable(x, y, w, h int, col tl.Attr, b *Board) *Clickable {
	return &Clickable{tl.NewRectangle(x, y, w, h, col,), b}
}


/*
I need to add a bar of colors that allow me to change
the current game color so that I can test out clicking
on the grid with different colors.
*/
func main(){
	white := tl.ColorWhite
	hieght := 10
	width := 20
	grid := []*Clickable{}
	board := Board{hieght, width, grid}

	gameColor := tl.NewRectangle(20, 20, 1, 1, tl.ColorGreen)

	g := tl.NewGame()

	// Used to created the board
	var tempt *Clickable
	for w:=0; w < board.Width; w++{
		for h:=0; h < board.Hieght; h++{
			tempt = NewClickable(h, w, 1, 1, white, &board)
			g.Screen().AddEntity(tempt)
			board.Grid = append(board.Grid, tempt)
		}
	}
	g.Screen().AddEntity(gameColor)
	g.Start()
}
