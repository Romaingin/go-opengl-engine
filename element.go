package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	mgl "github.com/go-gl/mathgl/mgl32"
	engine "github.com/rginestou/go-opengl-engine/v1"
)

// Element ...
type Element struct {
	object       engine.Object
	modelMatrix  mgl.Mat4
	modelUniform int32
}

func (e *Element) create(program uint32, vertices []float32, normals []float32, uvs []float32) {
	e.modelMatrix = mgl.Ident4()
	e.modelUniform = gl.GetUniformLocation(program, gl.Str("model\x00"))

	e.object.Create(vertices, normals, uvs)
}

func (e *Element) createFromFile(program uint32, filePath string) {
	e.modelMatrix = mgl.Ident4()
	e.modelUniform = gl.GetUniformLocation(program, gl.Str("model\x00"))

	e.object.CreateFromFile(filePath)
}

func (e *Element) draw() {
	gl.UniformMatrix4fv(e.modelUniform, 1, false, &e.modelMatrix[0])

	e.object.Draw()
}
