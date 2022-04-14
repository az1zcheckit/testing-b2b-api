package main

import (
	"b2b-api/internal"
	"b2b-api/internal/router"
	"go.uber.org/fx"
	"runtime"
)

func main() {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)

	app := fx.New(
		internal.Modules,
		router.EntryPoint,
	)
	app.Run()
}
