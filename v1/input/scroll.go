package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

// ScrollCallback ...
func ScrollCallback(w *glfw.Window, xoff float64, yoff float64) {
	s.scroll = yoff
}
