package main

import (
	"go-opengl-engine/input"
	"github.com/go-gl/glfw/v3.2/glfw"
	// mgl "github.com/go-gl/mathgl/mgl32"
)

var angle = 0.0

func update(s *Scene) {
	elapsed := s.clock.getElapsed()

	if input.KeyPressed(glfw.KeyEscape) {
		s.window.SetShouldClose(true)
	}

	xpos, ypos, dx, dy := input.MousePosition()
	xpos = ypos
	ypos = xpos
	ypos = dy

	angle += elapsed * dx
	s.camera.rotateY(float32(0.002*dx))
	//
	// model := mgl.HomogRotate3D(float32(angle), mgl.Vec3{0, 1, 0})
	// s.objects[0].SetModelMatrix(model)
}
