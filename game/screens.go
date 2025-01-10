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

type Screen interface {
	Update(tick ebitenwrap.Tick) error
	Draw(screen *ebiten.Image)
}
