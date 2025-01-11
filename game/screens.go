package game

import (
	"github.com/bfreis/ebitentools/ebitenwrap"
	"github.com/hajimehoshi/ebiten/v2"
)

type ScreenType int

const (
	ScreenTitle ScreenType = iota
	ScreenMaze
	ScreenAbout
)

type ScreenTransition struct {
	NextScreen ScreenType
	// For maze screen, we need to pass these parameters
	PlayerSpeed PlayerSpeed
	MazeSize    MazeSize
}

type Screen interface {
	Update(tick ebitenwrap.Tick) (*ScreenTransition, error)
	Draw(screen *ebiten.Image)
}
