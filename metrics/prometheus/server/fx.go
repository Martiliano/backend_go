package prometheusServer

//
// metrics => prometheus => server => fx.go
//

import "go.uber.org/fx"

var InitPromthesiusServerFx = fx.Options(
	fx.Provide(
		InitPromthesiusServer,
	),
)
