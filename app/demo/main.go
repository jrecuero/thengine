package main

import (
	"fmt"

	"github.com/jrecuero/thengine/pkg/engine"
)

func main() {
	fmt.Println("ThEngine test")
	appEngine := engine.NewEngine()
	appEngine.Init()
	appEngine.Run()
}
