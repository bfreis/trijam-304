package game

import (
	"image/color"

	"github.com/bfreis/ebitentools/ebitenwrap"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TitleScreen struct {
	selectedOption int
	options        []string
	tickCounter    int
}

func NewTitleScreen() *TitleScreen {
	return &TitleScreen{
		selectedOption: 0,
		options:        []string{"Start", "About"},
		tickCounter:    0,
	}
}

func (s *TitleScreen) Update(tick ebitenwrap.Tick) error {
	s.tickCounter++
	if s.tickCounter >= int(tick.TPS) { // Switch every second
		s.selectedOption = (s.selectedOption + 1) % len(s.options)
		s.tickCounter = 0
	}
	return nil
}

func (s *TitleScreen) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{40, 40, 40, 255})

	// Draw title
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(320, 200)
	text.Draw(screen, "Maze Game", face7x13, opts)

	// Draw menu options
	for i, option := range s.options {
		y := 300 + i*40
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(300, float64(y))
		if i == s.selectedOption {
			opts.ColorScale.Scale(1, 1, 0, 1) // Yellow
			text.Draw(screen, "> "+option, face7x13, opts)
		} else {
			opts.ColorScale.Scale(1, 1, 1, 1) // White
			text.Draw(screen, "  "+option, face7x13, opts)
		}
	}
}
