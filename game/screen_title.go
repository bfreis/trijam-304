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

type MazeSize int

const (
	SizeSmall MazeSize = iota
	SizeMedium
	SizeBig
)

func (s MazeSize) String() string {
	switch s {
	case SizeSmall:
		return "Small"
	case SizeMedium:
		return "Medium"
	case SizeBig:
		return "Big"
	default:
		return "Unknown"
	}
}

func (s MazeSize) Dimensions() (int, int) {
	switch s {
	case SizeSmall:
		return 3, 3
	case SizeMedium:
		return 10, 10
	case SizeBig:
		return 20, 20
	default:
		return 10, 10
	}
}

type TitleScreen struct {
	selectedOption int
	options        []string
	playerSpeed    PlayerSpeed
	mazeSize       MazeSize
	tickCounter    int
}

func NewTitleScreen() *TitleScreen {
	return &TitleScreen{
		selectedOption: 0,
		options:        []string{"Start", "Player Speed", "Maze Size", "About"},
		playerSpeed:    SpeedMedium,
		mazeSize:       SizeMedium,
		tickCounter:    0,
	}
}

func (s *TitleScreen) Update(tick ebitenwrap.Tick) (*ScreenTransition, error) {
	s.tickCounter++
	if s.tickCounter >= tick.TPS { // Switch every second
		s.selectedOption = (s.selectedOption + 1) % len(s.options)
		s.tickCounter = 0
	}
	if isButtonJustReleased(tick.InputState) {
		switch s.options[s.selectedOption] {
		case "Player Speed":
			s.playerSpeed = PlayerSpeed((int(s.playerSpeed) + 1) % 3)
		case "Maze Size":
			s.mazeSize = MazeSize((int(s.mazeSize) + 1) % 3)
		case "Start":
			return &ScreenTransition{
				NextScreen:  ScreenMaze,
				PlayerSpeed: s.playerSpeed,
				MazeSize:    s.mazeSize,
			}, nil
		case "About":
			return &ScreenTransition{
				NextScreen: ScreenAbout,
			}, nil
		}
	}
	return nil, nil
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
		switch option {
		case "Player Speed":
			menuText = option + ": " + s.playerSpeed.String()
		case "Maze Size":
			menuText = option + ": " + s.mazeSize.String()
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
