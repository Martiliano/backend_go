package cleanup

//
// presenter => cleanup => fx.go
//

import "go.uber.org/fx"

var CleanupFx = fx.Options(
	fx.Provide(
		GetCleanupConfig,
	),
)
