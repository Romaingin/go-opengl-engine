package engine

import (
	"os"

	"github.com/rginestou/check"
	"github.com/sheenobu/go-obj/obj"
)

// Geometry ...
type Geometry struct {
	vertices []float32
	normals  []float32
}

// LoadGeometryFormFile ...
func LoadGeometryFormFile(filePath string) (Geometry, error) {
	// load our OBJ
	f, err := os.Open(filePath)
	check.Panic(err)

	o, err := obj.NewReader(f).Read()
	check.Log(err)

	// New geometry to handle vertices and normals
	var g Geometry

	// Convert object into vertices for OpenGL
	for _, f := range o.Faces {
		for _, p := range f.Points {
			vx := p.Vertex
			nx := p.Normal

			g.normals = append(g.normals,
				[]float32{float32(nx.Z), float32(nx.Y), float32(nx.X)}...)

			g.vertices = append(g.vertices,
				[]float32{float32(vx.Z), float32(vx.Y), float32(vx.X)}...)
		}
	}

	// var g Geometry
	// g.vertices = []float32{
	// 	1.0, 1.0, 0.0,
	// 	0.0, 0.0, 0.0,
	// 	1.0, 0.0, 0.0,
	// 	1.0, 1.0, 0.0,
	// 	1.0, 0.0, 0.0,
	// 	1.0, 0.0, -1.0,
	// }

	// g.normals = []float32{
	// 	0.0, 0.0, 1.0,
	// 	0.0, 0.0, 1.0,
	// 	0.0, 0.0, 1.0,
	// 	1.0, 0.0, 0.0,
	// 	1.0, 0.0, 0.0,
	// 	1.0, 0.0, 0.0,
	// }

	return g, nil
}
