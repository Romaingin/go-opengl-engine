package engine

import (
	mgl "github.com/go-gl/mathgl/mgl32"
)

// Camera ...
type Camera struct {
	// Matrices
	Projection mgl.Mat4
	View       mgl.Mat4

	// Vision
	eye mgl.Vec3
	aim mgl.Vec3
	up  mgl.Vec3
}

// Init all the fields
func (cam *Camera) Init(windowWidth, windowHeight int) {
	// Vectors
	cam.eye = mgl.Vec3{5, 0.5, 5}
	cam.aim = mgl.Vec3{-1, 0, -1}.Normalize()
	cam.up = mgl.Vec3{0, 1, 0}

	// Compute Matrices
	cam.ComputeProjection(windowWidth, windowHeight)
	cam.ComputeView()
}

// ComputeProjection ...
func (cam *Camera) ComputeProjection(windowWidth, windowHeight int) {
	cam.Projection = mgl.Perspective(
		mgl.DegToRad(45.0),
		float32(windowWidth)/float32(windowHeight),
		0.1,
		100.0)
}

// ComputeView ...
func (cam *Camera) ComputeView() {
	cam.View = mgl.LookAtV(
		cam.eye,
		cam.eye.Add(cam.aim),
		cam.up)
}

// RotateY ...
func (cam *Camera) RotateY(angle float32) {
	cam.aim = mgl.Rotate3DY(angle).Mul3x1(cam.aim)
}

// MoveForward ...
func (cam *Camera) MoveForward(step float32) {
	cam.eye = cam.eye.Add(cam.aim.Mul(step))
}

// MoveSide ...
func (cam *Camera) MoveSide(step float32) {
	cam.eye = cam.eye.Add(cam.aim.Cross(cam.up).Mul(step))
}
