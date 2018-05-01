package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Stores all the input information
type state struct {
	keys map[glfw.Key]bool

	mouseX    float64
	mouseY    float64
	oldMouseX float64
	oldMouseY float64

	scroll float64
}

var s state

// Init ...
func Init(x, y float64) {
	s.keys = make(map[glfw.Key]bool)

	s.mouseX = x
	s.mouseY = y
	s.oldMouseX = x
	s.oldMouseY = y
}

// KeyPressed ...
func KeyPressed(k glfw.Key) bool {
	return s.keys[k]
}

// MousePosition returns the mouse position as well as the delta mouvement
func MousePosition() (float64, float64, float64, float64) {
	mouseDeltaX := s.mouseX - s.oldMouseX
	mouseDeltaY := s.mouseY - s.oldMouseY
	s.oldMouseX = s.mouseX
	s.oldMouseY = s.mouseY
	return s.mouseX, s.mouseY, mouseDeltaX, mouseDeltaY
}

// Scroll ...
func Scroll() float64 {
	val := s.scroll
	s.scroll = 0.0
	return val
}
