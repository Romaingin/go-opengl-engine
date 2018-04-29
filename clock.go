package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Clock struct {
	previousTime float64
	elapsed      float64
}

func (c *Clock) init() {
	c.previousTime = glfw.GetTime()
}

func (c *Clock) tick() {
	time := glfw.GetTime()
	c.elapsed = time - c.previousTime
	c.previousTime = time
}

func (c Clock) getElapsed() float64 {
	return glfw.GetTime() - c.previousTime
}
