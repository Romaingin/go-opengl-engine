package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/rginestou/go-opengl-engine/input"
	// mgl "github.com/go-gl/mathgl/mgl32"
)

var angle = 0.0

func update(s *Scene) {
	elapsed := float32(s.clock.getElapsed())

	// Key actions
	keyboardProcess(s, elapsed)

	xpos, ypos, dx, dy := input.MousePosition()
	xpos = ypos
	ypos = xpos
	ypos = dy

	// angle += elapsed * float32(dx)
	s.camera.rotateY(-float32(0.002 * dx))
	//
	// model := mgl.HomogRotate3D(float32(angle), mgl.Vec3{0, 1, 0})
	// s.objects[0].SetModelMatrix(model)
}

func keyboardProcess(s *Scene, elapsed float32) {
	if input.KeyPressed(glfw.KeyEscape) {
		s.window.SetShouldClose(true)
	}

	if input.KeyPressed(glfw.KeyW) {
		s.camera.MoveForward(8.0 * elapsed)
	}

	if input.KeyPressed(glfw.KeyS) {
		s.camera.MoveForward(-8.0 * elapsed)
	}

	if input.KeyPressed(glfw.KeyD) {
		s.camera.MoveSide(8.0 * elapsed)
	}

	if input.KeyPressed(glfw.KeyA) {
		s.camera.MoveSide(-8.0 * elapsed)
	}
}
