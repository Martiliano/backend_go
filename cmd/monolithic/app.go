package main

//
// cmd => monolithic => app.go
//

import (
	"go.uber.org/fx"
)

// GetApp : retornará o aplicativo fx.
// Fx App contém invocadores, provedores, ciclos de vida etc.
// Quando iniciamos o aplicativo usando o fx app, os provedores usados ​​serão inicializados primeiro.
// Depois disso, o invocador será invocado automaticamente.
// Nota: Os invocadores serão executados na mesma ordem.
func GetApp() *fx.App {
	opts := make([]fx.Option, 0)
	opts = GetProviderOptions()
	opts = append(opts, GetInvokersOptions())
	return fx.New(
		opts...,
	)
}
