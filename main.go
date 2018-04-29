package main

import (
	engine "github.com/rginestou/go-opengl-engine/v1"
)

const windowWidth = 1200
const windowHeight = 800

// Global OpenGL Engine
var e engine.Engine

func main() {
	// Config engine
	e.Title = "OpenGL 3D Engine"
	e.WindowWidth = windowWidth
	e.WindowHeight = windowHeight

	// Instanciate displayer
	var g game

	// Start app
	e.Start()
	g.init()

	// Main Loop
	e.Loop(&g)

	// Exit
	e.Stop()
}
