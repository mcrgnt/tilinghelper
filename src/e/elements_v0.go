package e

import (
	"fmt"
	"strings"
	"tilinghelper/src/g"
	"tilinghelper/src/p"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	triangle = []float32{
		-0.5, 0.5, 0.5,
		-0.5, -0.5, 0.5,
		0.5, -0.5, 0.5,
	}
)

const (
	vertexShaderSource = `
		#version 460
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 460
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(1, 1, 1, 1.0);
		}
	` + "\x00"
)

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

	return
}

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

	return
}

func init() {
	println("v0")
}
