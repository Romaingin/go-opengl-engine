package main

import (
	"github.com/Romaingin/go-opengl-engine/object"
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Uses OpenGL to draw the scene objects
func draw(s *Scene) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(s.program)

	// Update camera view
	s.camera.computeView()

	// Loop through objects and draw them
	for _, el := range(s.elements) {
		drawObject(el.GetObject())
	}

	// Disable VAO
	gl.BindVertexArray(0);

	// Maintenance
	s.window.SwapBuffers()
}

func drawObject(o object.ObjectInterface) {
	// Update model matrix
	model := o.GetModelMatrix()
	gl.UniformMatrix4fv(o.GetModelMatrixUniform(), 1, false, &model[0])

	// Draw buffer's content
	gl.BindVertexArray(o.GetVAO())
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}
