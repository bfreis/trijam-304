package game

import (
	"image/color"
	"time"

	"github.com/bfreis/ebitentools/ebitenwrap"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	mazeCellDisplayHeight = 32
	mazeCellDisplayWidth  = 32
	wallThickness         = 2.0
	rotationInterval      = time.Second
	winMessageScale       = 3.0
)

type MazeScreen struct {
	maze                   *Maze
	playerX                int
	playerY                int
	playerDirection        MazeDirection
	ticksSinceLastRotation int
	hasWon                 bool
	exitDirection          MazeDirection // Direction where player exited the maze
}

func NewMazeScreen() (*MazeScreen, error) {
	const mazeString = `+--+--+--+
|  |  |   
+  +--+  +
|        |
+  +  +--+
|  |     |
+--+--+--+`
	maze, err := ParseMaze(mazeString)
	if err != nil {
		return nil, err
	}
	return &MazeScreen{
		maze:                   maze,
		playerX:                0,
		playerY:                0,
		playerDirection:        MazeDirection(North),
		ticksSinceLastRotation: 0,
		hasWon:                 false,
	}, nil
}

func (s *MazeScreen) Update(tick ebitenwrap.Tick) error {
	if tick.InputState.Keyboard().IsKeyJustPressed(ebiten.KeyEscape) {
		return nil
	}

	if s.hasWon {
		return nil
	}

	// Rotate player direction every rotationInterval
	s.ticksSinceLastRotation++
	if s.ticksSinceLastRotation >= int(rotationInterval.Seconds()*float64(tick.TPS)) {
		s.playerDirection = MazeDirection((int(s.playerDirection) + 1) % 4)
		s.ticksSinceLastRotation = 0
	}

	// Move player when enter is pressed
	if tick.InputState.Keyboard().IsKeyJustPressed(ebiten.KeyEnter) {
		nextX, nextY := s.playerX, s.playerY
		switch s.playerDirection {
		case MazeDirection(North):
			nextY--
		case MazeDirection(East):
			nextX++
		case MazeDirection(South):
			nextY++
		case MazeDirection(West):
			nextX--
		}

		// Check if movement would lead to winning
		if !s.maze.HasWall(s.playerX, s.playerY, s.playerDirection) &&
			(nextX < 0 || nextX >= s.maze.Width || nextY < 0 || nextY >= s.maze.Height) {
			s.hasWon = true
			s.exitDirection = s.playerDirection
			return nil
		}

		// Check if movement is valid (within bounds and no wall)
		if nextX >= 0 && nextX < s.maze.Width &&
			nextY >= 0 && nextY < s.maze.Height &&
			!s.maze.HasWall(s.playerX, s.playerY, s.playerDirection) {
			s.playerX = nextX
			s.playerY = nextY
		}
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

	// Draw maze walls
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

	// Calculate player position, adjusting for win state
	playerX, playerY := s.playerX, s.playerY
	if s.hasWon {
		switch s.exitDirection {
		case MazeDirection(North):
			playerY = -1
		case MazeDirection(East):
			playerX = s.maze.Width
		case MazeDirection(South):
			playerY = s.maze.Height
		case MazeDirection(West):
			playerX = -1
		}
	}

	// Draw player
	playerRadius := float32(mazeCellDisplayWidth) / 4
	playerPosX := float32(offsetX + float64(playerX)*cellWidth + cellWidth/2)
	playerPosY := float32(offsetY + float64(playerY)*cellHeight + cellHeight/2)

	// Draw player body
	vector.DrawFilledCircle(screen, playerPosX, playerPosY, playerRadius, color.RGBA{255, 200, 0, 255}, false)

	// Draw direction indicator
	indicatorLength := playerRadius * 1.2
	var dx, dy float32
	switch s.playerDirection {
	case MazeDirection(North):
		dx, dy = 0, -indicatorLength
	case MazeDirection(East):
		dx, dy = indicatorLength, 0
	case MazeDirection(South):
		dx, dy = 0, indicatorLength
	case MazeDirection(West):
		dx, dy = -indicatorLength, 0
	}
	vector.StrokeLine(screen,
		playerPosX, playerPosY,
		playerPosX+dx, playerPosY+dy,
		wallThickness*2, color.RGBA{255, 100, 0, 255}, false)

	// Draw win message if player has won
	if s.hasWon {
		// Draw semi-transparent dark overlay
		vector.DrawFilledRect(screen, 0, 0, float32(sw), float32(sh), color.RGBA{0, 0, 0, 180}, false)

		const winMessage = "YOU WON!"

		opts := &text.DrawOptions{}
		opts.GeoM.Scale(winMessageScale, winMessageScale)
		opts.GeoM.Translate(
			float64(sw)/2-float64(len(winMessage)*7)*winMessageScale/2, // Approximate width based on monospace font
			float64(sh)/2-float64(13)*winMessageScale/2,                // Approximate height based on font size
		)
		opts.ColorScale.Scale(1, 1, 0, 1) // Yellow color
		text.Draw(screen, winMessage, face7x13, opts)
	}
}
