package game

import (
	"image/color"

	"github.com/bfreis/ebitentools/ebitenwrap"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type AboutScreen struct{}

func NewAboutScreen() *AboutScreen {
	return &AboutScreen{}
}

func (s *AboutScreen) Update(tick ebitenwrap.Tick) error {
	if tick.InputState.Keyboard().IsKeyJustPressed(ebiten.KeyEscape) {
		return nil
	}
	return nil
}

func (s *AboutScreen) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{40, 40, 40, 255})

	opts := &text.DrawOptions{}
	opts.GeoM.Translate(320, 200)
	text.Draw(screen, "About Maze Game", face7x13, opts)

	opts = &text.DrawOptions{}
	opts.GeoM.Translate(250, 250)
	text.Draw(screen, "A simple maze game created for learning Go", face7x13, opts)

	opts = &text.DrawOptions{}
	opts.GeoM.Translate(300, 350)
	text.Draw(screen, "Press ESC to return", face7x13, opts)
}
