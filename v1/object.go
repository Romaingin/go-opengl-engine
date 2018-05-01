package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/rginestou/check"
)

// Object ...
type Object struct {
	vao      uint32
	triCount int32
}

// Create ...
func (o *Object) Create(vertices []float32, normals []float32, uvs []float32) {
	// Create the Array Object based on the mesh
	o.vao = makeIndexedVAO(vertices, normals, uvs)
	o.triCount = int32(len(vertices) / 3)
}

// CreateFromFile ...
func (o *Object) CreateFromFile(filePath string) {
	g, err := LoadGeometryFormFile(filePath)
	check.Panic(err)

	// Create the Array Object based on the mesh
	o.vao = makeIndexedVAO(g.vertices, g.normals, nil)
	o.triCount = int32(len(g.vertices) / 3)
}

// Draw ...
func (o Object) Draw() {
	gl.BindVertexArray(o.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, o.triCount)
}
