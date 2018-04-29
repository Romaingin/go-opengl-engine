package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/rginestou/go-opengl-engine/v1"
	"github.com/rginestou/go-opengl-engine/v1/input"
)

// The game type fits the displayer interface
// of the OpenGL Engine
type game struct {
	cam            engine.Camera
	camProjUniform int32
	camViewUniform int32

	basicProgram uint32
	elements     []element
}

func (g *game) init() {
	// Init camera
	g.cam.Init(windowWidth, windowHeight)

	// Load shaders
	g.basicProgram = engine.CreateShaderProgram("shaders/basic.vs", "shaders/basic.fs")

	// Load scene elements
	g.elements = make([]element, 1)
	g.elements[0].create(g.basicProgram, "lib/mesh/cube.obj")

	// Camera Uniforms
	gl.UseProgram(g.basicProgram)
	g.camProjUniform = gl.GetUniformLocation(g.basicProgram, gl.Str("projection\x00"))
	g.camViewUniform = gl.GetUniformLocation(g.basicProgram, gl.Str("view\x00"))
	gl.UniformMatrix4fv(g.camProjUniform, 1, false, &g.cam.Projection[0])
}

// Draw ...
func (g *game) Draw(dt float64) {
	gl.UseProgram(g.basicProgram)

	// Update camera view
	gl.UniformMatrix4fv(g.camViewUniform, 1, false, &g.cam.View[0])

	// Loop through objects
	for el := range g.elements {
		g.elements[el].draw()
	}
}

// Update ...
func (g *game) Update(dt float64) {
	g.cam.ComputeView()

	g.keyboardProcess(float32(dt))
}

func (g *game) keyboardProcess(dt float32) {
	// if input.KeyPressed(glfw.KeyEscape) {
	// 	s.window.SetShouldClose(true)
	// }

	if input.KeyPressed(glfw.KeyW) {
		g.cam.MoveForward(8.0 * dt)
	}

	if input.KeyPressed(glfw.KeyS) {
		g.cam.MoveForward(-8.0 * dt)
	}

	if input.KeyPressed(glfw.KeyD) {
		g.cam.MoveSide(8.0 * dt)
	}

	if input.KeyPressed(glfw.KeyA) {
		g.cam.MoveSide(-8.0 * dt)
	}
}

// var angle = 0.0

// func update(s *Scene) {
// 	elapsed := float32(s.clock.getElapsed())

// 	// Key actions
// 	keyboardProcess(s, elapsed)

// 	xpos, ypos, dx, dy := input.MousePosition()
// 	xpos = ypos
// 	ypos = xpos
// 	ypos = dy

// 	// angle += elapsed * float32(dx)
// 	s.camera.rotateY(-float32(0.002 * dx))
// 	//
// 	// model := mgl.HomogRotate3D(float32(angle), mgl.Vec3{0, 1, 0})
// 	// s.objects[0].SetModelMatrix(model)
// }
