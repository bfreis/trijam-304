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

	switch g.currentScreen {
	case ScreenTitle:
		err = g.titleScreen.Update(tick)
		if err == nil && isButtonJustReleased(tick.InputState) {
			switch g.titleScreen.selectedOption {
			case 0: // Start
				g.mazeScreen, err = NewMazeScreen(g.titleScreen.playerSpeed, g.titleScreen.mazeSize)
				if err == nil {
					g.currentScreen = ScreenMaze
				}
			case 1: // Player Speed
				// Speed is handled in title screen
			case 2: // Maze Size
				// Size is handled in title screen
			case 3: // About
				g.currentScreen = ScreenAbout
			}
		}
	case ScreenMaze:
		err = g.mazeScreen.Update(tick)
		if err == nil && tick.InputState.Keyboard().IsKeyJustPressed(ebiten.KeyEscape) {
			g.currentScreen = ScreenTitle
		}
	case ScreenAbout:
		err = g.aboutScreen.Update(tick)
		if err == nil && (tick.InputState.Keyboard().IsKeyJustPressed(ebiten.KeyEscape) ||
			isButtonJustReleased(tick.InputState)) {
			g.currentScreen = ScreenTitle
		}
	}

	return err
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
