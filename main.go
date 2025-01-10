package main

import (
	"log"

	"github.com/bfreis/ebitentools/ebitenwrap"
	"github.com/bfreis/ebitentools/ebitenwrapfx"
	"github.com/bfreis/trijam-go/game"
	"github.com/hajimehoshi/ebiten/v2"
	"go.uber.org/fx"
)

func main() {
	run()
}

func run() {
	var g *ebitenwrap.Wrapper

	fxApp := fx.New(
		bootstrap(),
		fx.Populate(&g),
	)
	err := fxApp.Err()
	if err != nil {
		log.Fatalf("error creating fx app: %v", err)
	}

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Trijam")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	//ebiten.SetScreenClearedEveryFrame(false)
	//ebiten.SetVsyncEnabled(true)

	err = ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}

func bootstrap() fx.Option {
	return fx.Options(
		game.Module,
		ebitenwrapfx.Module,
		fx.Provide(func(g *game.Game) ebitenwrap.Game { return g }),
	)
}
