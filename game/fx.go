package game

import "go.uber.org/fx"

var Module = fx.Module("game",
	fx.Provide(
		NewGame,
	),
)
