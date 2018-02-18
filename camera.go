package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	mgl "github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	// Matrices
	projection mgl.Mat4
	view mgl.Mat4

	// Vision
	eye mgl.Vec3
	aim mgl.Vec3

	// Uniforms
	projectionUniform int32
	viewUniform int32
}

// Init all the fields
func (cam *Camera) init(program uint32) {
	// Vectors
	cam.eye = mgl.Vec3{5,0.5,5}
	cam.aim = mgl.Vec3{-1,0,-1}.Normalize()

	// Uniforms
	cam.projectionUniform = gl.GetUniformLocation(program, gl.Str("projection\x00"))
	cam.viewUniform = gl.GetUniformLocation(program, gl.Str("camera\x00"))

	// Compute Matrices
	cam.computeProjection()
	cam.computeView()
}

func (cam *Camera) computeProjection() {
	cam.projection = mgl.Perspective(
		mgl.DegToRad(45.0),
		float32(500)/float32(500),
		0.1,
		100.0)

	gl.UniformMatrix4fv(cam.projectionUniform, 1, false, &cam.projection[0])
}

func (cam *Camera) computeView() {
	cam.view = mgl.LookAtV(
		cam.eye,
		cam.eye.Add(cam.aim),
		mgl.Vec3{0, 1, 0})

	gl.UniformMatrix4fv(cam.viewUniform, 1, false, &cam.view[0])
}

func (cam *Camera) rotateY(angle float32) {
	cam.aim = mgl.Rotate3DY(angle).Mul3x1(cam.aim)
}
