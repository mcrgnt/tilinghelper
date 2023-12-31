package e

import (
	"fmt"
	"strings"
	"tilinghelper/pkg/g"
	"tilinghelper/pkg/p"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	z        float32 = -0.9
	triangle         = []float32{
		-0.5, 0.5, z,
		-0.5, -0.5, z,
		0.5, -0.5, z,
	}
)

const (
	// vertexShaderSource = `
	// 	#version 460
	// 	in vec3 vp;
	// 	void main() {
	// 		gl_Position = vec4(vp, 1.0);
	// 	}
	// ` + "\x00"

	vertexShaderSource = `
	#version 460
	in vec3 vp;
	void main() {
	gl_Position = vec4(vp, 1.0);
	}
	` + "\x00"

	uniform
	fragmentShaderSource = `
		#version 460
		uniform vec4 inDiffuseColor;
		out vec4 outDiffuseColor;
		void main() {
			outDiffuseColor = inDiffuseColor;
		}
	` + "\x00"
)

func GetEbo(window *glfw.Window) (vao uint32, err error) {
	defer glfw.DetachCurrentContext()
	window.MakeContextCurrent()

	if err = g.GlErrorHelper(); err != nil {
		return
	}

	var z float32 = 0.9

	vertices := []float32{
		-0.6, -0.6, z,
		-0.6, 0.4, z,
		0.4, 0.4, z,
		0.4, -0.6, z,
	}

	indices := []uint32{0, 1, 2, 2, 3, 0}

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, int32(0), uintptr(0))
	gl.EnableVertexAttribArray(0)

	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(indices), gl.Ptr(indices), gl.STATIC_DRAW)

	gl.BindVertexArray(0)

	return
}

func GetVao(window *glfw.Window) (vao uint32, err error) {
	defer glfw.DetachCurrentContext()
	window.MakeContextCurrent()

	if err = g.GlErrorHelper(); err != nil {
		return
	}

	var vbo uint32
	// get vertex buffer objects bunch id
	gl.GenBuffers(1, &vbo)

	// set vbo type
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	// set vbo data
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(triangle), gl.Ptr(triangle), gl.STATIC_DRAW)

	// get vertex array objects bunch id
	gl.GenVertexArrays(1, &vao)

	// bind vao to (current context?)
	gl.BindVertexArray(vao)

	// some needs for vao to be rendered
	gl.EnableVertexAttribArray(0)

	// again...
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	// set attributes?
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	gl.BindVertexArray(0)

	return
}

func compileShader(source string, shaderType uint32) (shader uint32, err error) {
	shader = gl.CreateShader(shaderType)

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

		err = fmt.Errorf("failed to compile %v: %v", source, log)
	} else {
		fmt.Println("good compile", shader)
	}

	return
}

func InitElements(window *glfw.Window, program uint32) (err error) {
	defer glfw.DetachCurrentContext()
	window.MakeContextCurrent()

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return
	}

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	p.Validate(program)
	p.LinkStatus(program)

	return
}
