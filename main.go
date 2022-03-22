package main

import (
	"github.com/evrintobing17/MyGram/app/registry"
)

func main() {
	appRegistry := registry.NewAppRegistry()
	appRegistry.StartServer()
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
