package game

import (
	"image/color"

	"github.com/bfreis/ebitentools/ebitenwrap"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

const (
	mazeCellDisplayHeight = 32
	mazeCellDisplayWidth  = 32
	wallThickness         = 2.0
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
	text.Draw(screen, "Maze Game", basicfont.Face7x13, 320, 200, color.White)

	// Draw menu options
	for i, option := range s.options {
		y := 300 + i*40
		if i == s.selectedOption {
			text.Draw(screen, "> "+option, basicfont.Face7x13, 300, y, color.RGBA{255, 255, 0, 255})
		} else {
			text.Draw(screen, "  "+option, basicfont.Face7x13, 300, y, color.White)
		}
	}
}

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
	text.Draw(screen, "About Maze Game", basicfont.Face7x13, 320, 200, color.White)
	text.Draw(screen, "A simple maze game created for learning Go", basicfont.Face7x13, 250, 250, color.White)
	text.Draw(screen, "Press ESC to return", basicfont.Face7x13, 300, 350, color.White)
}

type MazeScreen struct {
	maze *Maze
}

func NewMazeScreen() (*MazeScreen, error) {
	const mazeString = `+--+--+--+
|  |  |  |
+--+--+--+
|     |  |
+--+--+--+
|  |  |  |
+--+--+--+`
	maze, err := ParseMaze(mazeString)
	if err != nil {
		return nil, err
	}
	return &MazeScreen{maze: maze}, nil
}

func (s *MazeScreen) Update(tick ebitenwrap.Tick) error {
	if tick.InputState.Keyboard().IsKeyJustPressed(ebiten.KeyEscape) {
		return nil
	}
	return nil
}

func (s *MazeScreen) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{40, 40, 40, 255})

	cellWidth := float64(mazeCellDisplayWidth)
	cellHeight := float64(mazeCellDisplayHeight)

	mazeWidth := float64(s.maze.Width) * cellWidth
	mazeHeight := float64(s.maze.Height) * cellHeight

	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	offsetX := (float64(sw) - mazeWidth) / 2
	offsetY := (float64(sh) - mazeHeight) / 2

	for y := 0; y < s.maze.Height; y++ {
		for x := 0; x < s.maze.Width; x++ {
			px := offsetX + float64(x)*cellWidth
			py := offsetY + float64(y)*cellHeight

			if s.maze.HasWall(x, y, North) {
				vector.StrokeLine(screen, float32(px), float32(py), float32(px+cellWidth), float32(py), float32(wallThickness), color.White, false)
			}
			if s.maze.HasWall(x, y, East) {
				vector.StrokeLine(screen, float32(px+cellWidth), float32(py), float32(px+cellWidth), float32(py+cellHeight), float32(wallThickness), color.White, false)
			}
			if s.maze.HasWall(x, y, South) {
				vector.StrokeLine(screen, float32(px), float32(py+cellHeight), float32(px+cellWidth), float32(py+cellHeight), float32(wallThickness), color.White, false)
			}
			if s.maze.HasWall(x, y, West) {
				vector.StrokeLine(screen, float32(px), float32(py), float32(px), float32(py+cellHeight), float32(wallThickness), color.White, false)
			}
		}
	}
}

type Game struct {
	currentScreen ScreenType
	titleScreen   *TitleScreen
	mazeScreen    *MazeScreen
	aboutScreen   *AboutScreen
}

func NewGame() (*Game, error) {
	mazeScreen, err := NewMazeScreen()
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
		if err == nil && tick.InputState.Keyboard().IsKeyJustPressed(ebiten.KeyEnter) {
			switch g.titleScreen.selectedOption {
			case 0: // Start
				g.currentScreen = ScreenMaze
			case 1: // About
				g.currentScreen = ScreenAbout
			}
		}
	case ScreenMaze:
		err = g.mazeScreen.Update(tick)
		if err == nil && tick.InputState.Keyboard().IsKeyJustPressed(ebiten.KeyEscape) {
			g.currentScreen = ScreenTitle
			g.titleScreen = NewTitleScreen() // Reset title screen state
		}
	case ScreenAbout:
		err = g.aboutScreen.Update(tick)
		if err == nil && tick.InputState.Keyboard().IsKeyJustPressed(ebiten.KeyEscape) {
			g.currentScreen = ScreenTitle
			g.titleScreen = NewTitleScreen() // Reset title screen state
		}
	}

	return err
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
