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
func (o *Object) Create(filePath string) {
	g, err := LoadGeometryFormFile(filePath)
	check.Panic(err)

	// Create the Array Object based on the mesh
	o.vao = makeIndexedVAO(g.vertices, g.normals)
	o.triCount = int32(len(g.vertices) / 3)
}

// Draw ...
func (o Object) Draw() {
	gl.BindVertexArray(o.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, o.triCount)
}
