package jaeger

//
// metrics => opentracing => trace => jaeger => fx.go
//

import "go.uber.org/fx"

var JaegerTracerFx = fx.Options(
	fx.Provide(
		InitJaeger,
	),
)
