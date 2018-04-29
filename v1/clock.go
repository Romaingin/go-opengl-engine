package engine

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Clock ...
type clock struct {
	previousTime float64
	elapsed      float64
}

func (c *clock) init() {
	c.previousTime = glfw.GetTime()
}

func (c *clock) tick() {
	time := glfw.GetTime()
	c.elapsed = time - c.previousTime
	c.previousTime = time
}

func (c clock) getElapsed() float64 {
	return c.elapsed
}
