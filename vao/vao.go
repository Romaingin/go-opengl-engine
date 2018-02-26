package vao

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// returns a vertex array from the points provided
func Make(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

// returns a vertex array from the points provided
func MakeIndexed(program uint32, points []float32, normals []float32) uint32 {
	// Get attribute locations
	vert, free := gl.Strs("vert")
	defer free()
	vertexAttribVert := uint32(gl.GetAttribLocation(program, *vert))
	norm, free := gl.Strs("norm")
	defer free()
	vertexAttribNorm := uint32(gl.GetAttribLocation(program, *norm))

	// Create VAO buffer
	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	// var pointsEBO uint32
	// gl.GenBuffers(1, &pointsEBO)
	// gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, pointsEBO)
	// gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(pointsIndices), gl.Ptr(pointsIndices), gl.STATIC_DRAW)

	// Create Points VBO buffer
	var pointsVBO uint32
	gl.GenBuffers(1, &pointsVBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, pointsVBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)
	gl.VertexAttribPointer(vertexAttribVert, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(vertexAttribVert)

	// var normalsEBO uint32
	// gl.GenBuffers(1, &normalsEBO)
	// gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, normalsEBO)
	// gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(normalsIndices), gl.Ptr(normalsIndices), gl.STATIC_DRAW)

	// Create Normals VBO buffer
	var normalsVBO uint32
	gl.GenBuffers(1, &normalsVBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, normalsVBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(normals), gl.Ptr(normals), gl.STATIC_DRAW)
	gl.VertexAttribPointer(vertexAttribNorm, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(vertexAttribNorm)


	// Points
	// gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, pointsEBO)
	// gl.BindBuffer(gl.ARRAY_BUFFER, pointsVBO)

	// Normals
	// gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, normalsEBO)
	// gl.BindBuffer(gl.ARRAY_BUFFER, normalsVBO)

	return VAO
}
