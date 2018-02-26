package vao

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// returns a vertex array from the vertices provided
func Make(vertices []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

// returns a vertex array from the vertices provided
func MakeIndexed(vertices []float32, normals []float32) uint32 {
	// Create VAO buffer
	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	// var verticesEBO uint32
	// gl.GenBuffers(1, &verticesEBO)
	// gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, verticesEBO)
	// gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(verticesIndices), gl.Ptr(verticesIndices), gl.STATIC_DRAW)

	// Create vertices VBO buffer
	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	// var normalsEBO uint32
	// gl.GenBuffers(1, &normalsEBO)
	// gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, normalsEBO)
	// gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(normalsIndices), gl.Ptr(normalsIndices), gl.STATIC_DRAW)

	// Create Normals VBO buffer
	var NBO uint32
	gl.GenBuffers(1, &NBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, NBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(normals), gl.Ptr(normals), gl.STATIC_DRAW)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(1)

	// vertices
	// gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, verticesEBO)
	// gl.BindBuffer(gl.ARRAY_BUFFER, verticesVBO)

	// Normals
	// gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, normalsEBO)
	// gl.BindBuffer(gl.ARRAY_BUFFER, normalsVBO)

	return VAO
}
