package engine

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/rginestou/check"
	"github.com/rginestou/go-opengl-engine/v1/input"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

// Engine ...
type Engine struct {
	WindowWidth  int
	WindowHeight int
	Title        string
}

// Displayer ...
type Displayer interface {
	Update(float64)
	Draw(float64)
}

var window *glfw.Window

// Start the engine
func (e *Engine) Start() {
	// Init window
	check.Panic(glfw.Init())

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Samples, 4)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// Create Window
	var err error
	window, err = glfw.CreateWindow(e.WindowWidth, e.WindowHeight, e.Title, nil, nil)
	check.Panic(err)

	window.MakeContextCurrent()

	// Set inputs callbacks
	window.SetKeyCallback(input.KeyCallBack)
	window.SetCursorPosCallback(input.MouseCallback)
	window.SetScrollCallback(input.ScrollCallback)
	// window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)

	input.Init(window.GetCursorPos())

	// Init GL context
	check.Panic(gl.Init())

	// Display version
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
}

// Loop executes the main loop
func (e *Engine) Loop(dis Displayer) {
	var cl clock
	cl.init()

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	glfw.SwapInterval(1)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Get time
		cl.tick()

		dis.Update(cl.getElapsed())
		dis.Draw(cl.getElapsed())

		window.SwapBuffers()

		// Detect inputs
		glfw.PollEvents()
	}
}

// Stop the engine's execution
func (e *Engine) Stop() {
	glfw.Terminate()
}
