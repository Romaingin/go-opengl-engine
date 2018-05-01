package engine

import (
	mgl "github.com/go-gl/mathgl/mgl32"
)

// Camera ...
type Camera struct {
	// Matrices
	Projection mgl.Mat4
	Ortho      mgl.Mat4
	View       mgl.Mat4

	// Vision
	Eye mgl.Vec3
	Aim mgl.Vec3
}

var up = mgl.Vec3{0, 1, 0}

// SetEyeAim ...
func (cam *Camera) SetEyeAim(eye, aim mgl.Vec3) {
	cam.Eye = eye
	cam.Aim = aim.Normalize()
}

// ComputeProjection ...
func (cam *Camera) ComputeProjection(windowWidth, windowHeight int) {
	cam.Projection = mgl.Perspective(
		mgl.DegToRad(45.0),
		float32(windowWidth)/float32(windowHeight),
		0.1,
		100.0)
}

// ComputeOrtho ...
func (cam *Camera) ComputeOrtho(windowWidth, windowHeight int) {
	cam.Ortho = mgl.Ortho(0.0, float32(windowWidth), float32(windowHeight), 0.0, 0.0, 1.0)
}

// ComputeView ...
func (cam *Camera) ComputeView() {
	cam.View = mgl.LookAtV(
		cam.Eye,
		cam.Eye.Add(cam.Aim),
		up)
}

// RotateY ...
func (cam *Camera) RotateY(angle float32) {
	cam.Aim = mgl.Rotate3DY(angle).Mul3x1(cam.Aim)
}

// MoveForward ...
func (cam *Camera) MoveForward(step float32) {
	cam.Eye = cam.Eye.Add(cam.Aim.Mul(step))
}

// MoveSide ...
func (cam *Camera) MoveSide(step float32) {
	cam.Eye = cam.Eye.Add(cam.Aim.Cross(up).Mul(step))
}
