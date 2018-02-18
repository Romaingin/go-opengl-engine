package element

import (
	"opengl_engine/object"
)

type ElementInterface interface {
	GetObject() object.Object
}

type Element struct {
	object object.Object
}

func (e *Element) Create(program uint32, mesh []float32) {
	e.object.Create(program, mesh)
}

func (e Element) GetObject() object.Object {
	return e.object
}
