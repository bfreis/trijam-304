package game

import (
	"image/color"

	"github.com/bfreis/ebitentools/ebitenwrap"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type PlayerSpeed int

const (
	SpeedLow PlayerSpeed = iota
	SpeedMedium
	SpeedHigh
)

func (s PlayerSpeed) String() string {
	switch s {
	case SpeedLow:
		return "Low"
	case SpeedMedium:
		return "Medium"
	case SpeedHigh:
		return "High"
	default:
		return "Unknown"
	}
}

func (s PlayerSpeed) RotationsPerSecond() float64 {
	switch s {
	case SpeedLow:
		return 1.0
	case SpeedMedium:
		return 2.0
	case SpeedHigh:
		return 4.0
	default:
		return 1.0
	}
}

type TitleScreen struct {
	selectedOption int
	options        []string
	playerSpeed    PlayerSpeed
	tickCounter    int
}

func NewTitleScreen() *TitleScreen {
	return &TitleScreen{
		selectedOption: 0,
		options:        []string{"Start", "Player Speed", "About"},
		playerSpeed:    SpeedLow,
		tickCounter:    0,
	}
}

func (s *TitleScreen) Update(tick ebitenwrap.Tick) error {
	s.tickCounter++
	if s.tickCounter >= int(tick.TPS) { // Switch every second
		s.selectedOption = (s.selectedOption + 1) % len(s.options)
		s.tickCounter = 0
	}
	if tick.InputState.Keyboard().IsKeyJustPressed(ebiten.KeyEnter) {
		if s.options[s.selectedOption] == "Player Speed" {
			s.playerSpeed = PlayerSpeed((int(s.playerSpeed) + 1) % 3)
		}
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

		menuText := option
		if option == "Player Speed" {
			menuText = option + ": " + s.playerSpeed.String()
		}

		if i == s.selectedOption {
			opts.ColorScale.Scale(1, 1, 0, 1) // Yellow
			text.Draw(screen, "> "+menuText, face7x13, opts)
		} else {
			opts.ColorScale.Scale(1, 1, 1, 1) // White
			text.Draw(screen, "  "+menuText, face7x13, opts)
		}
	}
}
