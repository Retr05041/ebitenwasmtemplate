package lib

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	X, Y, W, H int
	Label      string
}

func (b *Button) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(b.X), float32(b.Y), float32(b.W), float32(b.H), color.RGBA{100, 100, 200, 255}, false)
	textX := b.X + 10
	textY := b.Y + b.H/2 - 4
	ebitenutil.DebugPrintAt(screen, b.Label, textX, textY)
}

func (b *Button) Contains(x, y int) bool {
	return x >= b.X && x <= b.X+b.W && y >= b.Y && y <= b.Y+b.H
}
