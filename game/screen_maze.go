package game

import (
	"image/color"

	"github.com/bfreis/ebitentools/ebitenwrap"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	mazeCellDisplayHeight = 32
	mazeCellDisplayWidth  = 32
	wallThickness         = 2.0
)

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
