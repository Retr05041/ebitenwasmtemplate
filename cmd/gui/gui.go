package gui

import (
	"ebitenwasmtemplate/cmd/lib"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Buttons []lib.Button
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for _, b := range g.Buttons {
			if b.Contains(x, y) {
				log.Println("Clicked:", b.Label)
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{30, 30, 30, 255})

	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	buttonWidth := 200
	buttonHeight := 40
	verticalSpacing := 10

	totalHeight := len(g.Buttons)*(buttonHeight+verticalSpacing) - verticalSpacing
	startY := (sh - totalHeight) / 2

	for i := range g.Buttons {
		g.Buttons[i].W = buttonWidth
		g.Buttons[i].H = buttonHeight
		g.Buttons[i].X = (sw - buttonWidth) / 2
		g.Buttons[i].Y = startY + i*(buttonHeight+verticalSpacing)
		g.Buttons[i].Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func Run() {
	game := &Game{
		Buttons: []lib.Button{
			{X: 220, Y: 100, W: 200, H: 40, Label: "Start Game"},
			{X: 220, Y: 150, W: 200, H: 40, Label: "Options"},
			{X: 220, Y: 200, W: 200, H: 40, Label: "Quit"},
		},
	}
	ebiten.SetWindowTitle("Main Menu")
	ebiten.SetWindowSize(960, 540)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
