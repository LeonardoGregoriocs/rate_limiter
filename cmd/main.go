package main

import (
	"time"

	"github.com/leonardogregoriocs/rate_limiter/config"
	dependencyinjector "github.com/leonardogregoriocs/rate_limiter/internal/pkg/dependency_injector"
)

func main() {
	cfg, err := config.GetConfigs(".")
	if err != nil {
		panic(err)
	}

	dependencyInjector := dependencyinjector.NewDependencyInjector(cfg)

	deps, err := dependencyInjector.Inject()
	if err != nil {
		panic(err)
	}

	deps.WebServer.Start()
	time.Sleep(time.Duration(time.Second.Minutes()) * 30)
}
