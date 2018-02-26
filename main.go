package main

import (
	"log"
	"runtime"
	"github.com/Romaingin/go-opengl-engine/element"
	"github.com/Romaingin/go-opengl-engine/input"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const windowWidth = 800
const windowHeight = 600

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

// Main routine
// Launches the window and OpenGL context
func main() {
	// Init window
	window := initGlfw()
	defer glfw.Terminate()

	// Init GL context
	err := gl.Init()
	check(err)

	// Display version
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	// Load shader program
	program, err := createShaderProgram()
	check(err)
	gl.UseProgram(program)

	// Create Camera
	var camera Camera
	camera.init(program)

	var clock Clock
	clock.init()

	// Create elements to be drawn
	elements := make([]element.Element, 2)
	elements[0].CreateFromFile(program, "lib/mesh/cube.obj")
	elements[1].CreateFromFile(program, "lib/mesh/floor.obj")

	// Create main scene
	var scene = Scene{window, &clock, &camera, program, elements}

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	glfw.SwapInterval(1)

	for !window.ShouldClose() {
		// Get time
		clock.tick()
		// if (clock.getElapsed() > 1.0 / 60) {
		// 	println(1.0 / clock.getElapsed())
		// }

		// Draw & Update routines
		draw(&scene)
		update(&scene)

		// Detect inputs
		glfw.PollEvents()
	}
}

// Initializes glfw and returns a Window to use.
func initGlfw() *glfw.Window {
	err := glfw.Init()
	check(err)

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Samples, 4)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// Create Window
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "OpenGL GO Engine", nil, nil)
	check(err)
	window.MakeContextCurrent()

	// Set inputs callbacks
	window.SetKeyCallback(input.KeyCallBack)
	window.SetCursorPosCallback(input.MouseCallback)
	window.SetScrollCallback(input.ScrollCallback)

	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	input.Init(window.GetCursorPos())

	return window
}
