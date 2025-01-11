package game

import (
	"github.com/bfreis/ebitentools/ebitenwrap"
	"github.com/hajimehoshi/ebiten/v2"
)

func isButtonJustReleased(input ebitenwrap.InputState) bool {
	// Check keyboard
	if input.Keyboard().IsKeyJustReleased(ebiten.KeyEnter) {
		return true
	}

	// Check mouse - left button
	if input.Mouse().IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		return true
	}

	// Check touch - any touch release counts
	touchIDs := input.Touch().AppendJustReleasedTouchIDs(nil)
	return len(touchIDs) > 0
}

type Game struct {
	currentScreen ScreenType
	titleScreen   *TitleScreen
	mazeScreen    *MazeScreen
	aboutScreen   *AboutScreen
}

func NewGame() (*Game, error) {
	mazeScreen, err := NewMazeScreen(SpeedMedium, SizeMedium)
	if err != nil {
		return nil, err
	}

	return &Game{
		currentScreen: ScreenTitle,
		titleScreen:   NewTitleScreen(),
		mazeScreen:    mazeScreen,
		aboutScreen:   NewAboutScreen(),
	}, nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.currentScreen {
	case ScreenTitle:
		g.titleScreen.Draw(screen)
	case ScreenMaze:
		g.mazeScreen.Draw(screen)
	case ScreenAbout:
		g.aboutScreen.Draw(screen)
	}
}

func (g *Game) Update(tick ebitenwrap.Tick) error {
	var err error
	var transition *ScreenTransition

	switch g.currentScreen {
	case ScreenTitle:
		transition, err = g.titleScreen.Update(tick)
		if err == nil && transition != nil {
			switch transition.NextScreen {
			case ScreenMaze:
				g.mazeScreen, err = NewMazeScreen(transition.PlayerSpeed, transition.MazeSize)
				if err == nil {
					g.currentScreen = ScreenMaze
				}
			case ScreenAbout:
				g.currentScreen = ScreenAbout
			}
		}
	case ScreenMaze:
		transition, err = g.mazeScreen.Update(tick)
		if err == nil && transition != nil {
			g.currentScreen = transition.NextScreen
		}
	case ScreenAbout:
		transition, err = g.aboutScreen.Update(tick)
		if err == nil && transition != nil {
			g.currentScreen = transition.NextScreen
		}
	}

	return err
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
