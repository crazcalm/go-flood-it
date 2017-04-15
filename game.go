package main

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
)

//Game
type Game struct {
	Board *Board
	CurrentColor tl.Attr
	Turn int
}

//ChangeColor
func (g *Game) ChangeColor () tl.Attr {
	colorList := [] tl.Attr{tl.ColorBlack, tl.ColorBlue, tl.ColorCyan, tl.ColorRed,
							tl.ColorGreen, tl.ColorYellow, tl.ColorWhite, tl.ColorMagenta}
	numOfColors := len(colorList)
	g.CurrentColor = colorList[rand.Intn(numOfColors)]
	return g.CurrentColor
}


//Board to the game
type Board struct {
	Hieght int
	Width  int
	Grid []*Clickable
}


//ColorRect
type ColorRect struct {
	*tl.Rectangle
	Game *Game
}

func (c ColorRect) Tick (ev tl.Event) {
	x, y := c.Position()
	if ev.Type == tl.EventMouse && ev.MouseX == x && ev.MouseY == y {
		c.SetColor(c.Game.ChangeColor())
	}
}

//Clickable ... add docs
type Clickable struct {
	*tl.Rectangle
	Game *Game
}

//Tick ... add docs
func (c *Clickable) Tick (ev tl.Event) {
	x, y := c.Position()
	if ev.Type == tl.EventMouse && ev.MouseX == x && ev.MouseY == y {
		c.SetColor(c.Game.CurrentColor)
		c.Game.Turn += 1
	}
}

//NewClickable ... add docs
func NewClickable(x, y, w, h int, col tl.Attr, g *Game) *Clickable {
	return &Clickable{tl.NewRectangle(x, y, w, h, col,), g}
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
	game := Game{&board, white, 0}


	gameColor := &ColorRect{tl.NewRectangle(20, 20, 1, 1, white), &game}

	g := tl.NewGame()

	// Used to created the board
	var tempt *Clickable
	for w:=0; w < board.Width; w++{
		for h:=0; h < board.Hieght; h++{
			tempt = NewClickable(h, w, 1, 1, white, &game)
			g.Screen().AddEntity(tempt)
			board.Grid = append(board.Grid, tempt)
		}
	}
	g.Screen().AddEntity(gameColor)
	g.Start()
}
