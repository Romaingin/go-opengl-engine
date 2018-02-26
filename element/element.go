package element

import (
	"github.com/Romaingin/go-opengl-engine/object"
)

type ElementInterface interface {
	GetObject() object.Object
}

type Element struct {
	object object.Object
}

func (e *Element) Create(program uint32, g object.Geometry) {
	e.object.Create(program, g)
}

func (e *Element) CreateFromFile(program uint32, filePath string) {
	geometry, err := object.LoadGeometryFormFile(filePath)
	check(err)

	e.object.Create(program, geometry)
}

func (e Element) GetObject() object.Object {
	return e.object
}
