package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	mgl "github.com/go-gl/mathgl/mgl32"
	"github.com/rginestou/go-opengl-engine/v1"
	"github.com/rginestou/go-opengl-engine/v1/input"
)

// The game type fits the displayer interface
// of the OpenGL Engine
type game struct {
	cam            engine.Camera
	camProjUniform int32
	camOrthUniform int32
	camViewUniform int32

	gui     GUI
	texture uint32

	basicProgram uint32
	elements     []Element
}

func (g *game) init() {
	// Load shaders
	g.basicProgram = engine.CreateShaderProgram("shaders/basic.vs", "shaders/basic.fs")

	// Load scene elements
	g.elements = make([]Element, 1)
	g.elements[0].createFromFile(g.basicProgram, "lib/mesh/cube.obj")

	// Load GUI
	g.gui.buttons = make([]Button, 1)
	vertices := []float32{
		0.0, 0.0, 0.0,
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		1.0, 1.0, 0.0,
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
	}
	uvs := []float32{
		0.0, 0.0,
		1.0, 0.0,
		0.0, 1.0,
		1.0, 1.0,
		1.0, 0.0,
		0.0, 1.0,
	}
	g.gui.buttons[0].create(g.basicProgram, vertices, nil, uvs)

	// Camera Set up
	g.cam.SetEyeAim(mgl.Vec3{2, 2.5, 2}, mgl.Vec3{-1, -1, -1})
	g.cam.ComputeProjection(windowWidth, windowHeight)
	g.cam.ComputeOrtho(windowWidth, windowHeight)

	gl.UseProgram(g.basicProgram)

	g.texture = engine.TextureFromFile("lib/img/rock.png")
	textureUniform := gl.GetUniformLocation(g.basicProgram, gl.Str("tex\x00"))
	gl.Uniform1i(textureUniform, 0)

	g.camProjUniform = gl.GetUniformLocation(g.basicProgram, gl.Str("projection\x00"))
	g.camOrthUniform = gl.GetUniformLocation(g.basicProgram, gl.Str("ortho\x00"))
	g.camViewUniform = gl.GetUniformLocation(g.basicProgram, gl.Str("view\x00"))
	gl.UniformMatrix4fv(g.camProjUniform, 1, false, &g.cam.Projection[0])
}

// Draw ...
func (g *game) Draw(dt float64) {
	gl.UseProgram(g.basicProgram)

	// Update camera view
	gl.UniformMatrix4fv(g.camViewUniform, 1, false, &g.cam.View[0])

	// Loop through Scene elements
	// for el := range g.elements {
	// g.elements[el].draw()
	// }

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, g.texture)

	// Loop through GUI elements
	for b := range g.gui.buttons {
		g.gui.buttons[b].draw()
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

	s := input.Scroll()
	if s != 0.0 {
		eye := g.cam.Eye
		aim := g.cam.Aim
		g.cam.SetEyeAim(eye.Add(aim.Mul(float32(s)*dt*30)), aim)
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
