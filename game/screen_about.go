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
	text.Draw(screen, "A maze game created for Trijam #304", face7x13, opts)

	opts = &text.DrawOptions{}
	opts.GeoM.Translate(250, 300)
	text.Draw(screen, "Theme: \"one button adventure\"", face7x13, opts)

	opts = &text.DrawOptions{}
	opts.GeoM.Translate(250, 325)
	text.Draw(screen, "Time taken: 2h, with help from Cursor", face7x13, opts)

	opts = &text.DrawOptions{}
	opts.GeoM.Translate(250, 350)
	text.Draw(screen, "https://bfreis.itch.io/single-button-maze", face7x13, opts)

	opts = &text.DrawOptions{}
	opts.GeoM.Translate(300, 400)
	text.Draw(screen, "Press ESC to return", face7x13, opts)
}
