package engine

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/rginestou/check"
)

// Load a file given its name
func loadShaderFile(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
	check.LogAndPanic(err)

	return string(dat)
}

// Compile the shader
func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

// CreateShaderProgram load the shader files and create a OpenGL program
func CreateShaderProgram(vertexPath, fragmentPath string) uint32 {
	vertexShaderSource := loadShaderFile(vertexPath) + "\x00"
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	check.Panic(err)

	fragmentShaderSource := loadShaderFile(fragmentPath) + "\x00"
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	check.Panic(err)

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)

	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		panic(fmt.Errorf("failed to link program: %v", log))
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program
}
