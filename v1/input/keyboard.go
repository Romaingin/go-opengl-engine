package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Alter the keys' state
func KeyCallBack(w *glfw.Window, k glfw.Key, st int, a glfw.Action, mk glfw.ModifierKey) {
	if a == glfw.Press {
		s.keys[k] = true
	} else if a == glfw.Release {
		s.keys[k] = false
	}
}
