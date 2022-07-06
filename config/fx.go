package config

//
// microservices => config => fx.go
//

import "go.uber.org/fx"

var ProviderFx = fx.Options(
	fx.Provide(
		GetConfig,
	),
)
