package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

func MouseCallback(w *glfw.Window, xpos float64, ypos float64) {
	s.mouseX = xpos
	s.mouseY = ypos
}
