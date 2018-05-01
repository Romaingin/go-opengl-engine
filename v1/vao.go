package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

const sizeOfFloat32 = 4

// returns a vertex array from the vertices provided
func makeVAO(vertices []float32) uint32 {
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
func makeIndexedVAO(vertices []float32, normals []float32, uvs []float32) uint32 {
	// Create VAO buffer
	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	// Create Vertices buffer
	if vertices != nil {
		var VBO uint32
		gl.GenBuffers(1, &VBO)
		gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
		gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*sizeOfFloat32, gl.Ptr(vertices), gl.STATIC_DRAW)

		gl.EnableVertexAttribArray(0)
		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	}

	// Create Normals buffer
	if normals != nil {
		var NBO uint32
		gl.GenBuffers(1, &NBO)
		gl.BindBuffer(gl.ARRAY_BUFFER, NBO)
		gl.BufferData(gl.ARRAY_BUFFER, len(normals)*sizeOfFloat32, gl.Ptr(normals), gl.STATIC_DRAW)

		gl.EnableVertexAttribArray(1)
		gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 0, nil)
	}

	// Create UVs buffer
	if uvs != nil {
		var UBO uint32
		gl.GenBuffers(1, &UBO)
		gl.BindBuffer(gl.ARRAY_BUFFER, UBO)
		gl.BufferData(gl.ARRAY_BUFFER, len(uvs)*sizeOfFloat32, gl.Ptr(uvs), gl.STATIC_DRAW)

		gl.EnableVertexAttribArray(2)
		gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 0, nil)
	}

	// Disable VAO
	gl.BindVertexArray(0)

	return VAO
}
