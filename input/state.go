package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Stores all the input information
type state struct {
	keys map[glfw.Key]bool
	mouseX float64
	mouseY float64
	oldMouseX float64
	oldMouseY float64
}

var s state

func Init(x, y float64) {
	s.keys = make(map[glfw.Key]bool)

	s.mouseX = x
	s.mouseY = y
	s.oldMouseX = x
	s.oldMouseY = y
}

func KeyPressed(k glfw.Key) bool {
	return s.keys[k]
}

// Returns the mouse position as well as the delta mouvement
func MousePosition() (float64, float64, float64, float64) {
	mouseDeltaX := s.mouseX - s.oldMouseX
	mouseDeltaY := s.mouseY - s.oldMouseY
	s.oldMouseX = s.mouseX
	s.oldMouseY = s.mouseY
	return s.mouseX, s.mouseY, mouseDeltaX, mouseDeltaY
}
