package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/rginestou/go-opengl-engine/element"
)

// Groups the scene's crucial elements
type Scene struct {
	window   *glfw.Window
	clock    *Clock
	camera   *Camera
	program  uint32
	elements []element.Element
}
