//go:build v1

package e

import (
	"fmt"
	"reflect"
	"strings"
	"tilinghelper/src/g"
	"unsafe"

	"github.com/go-gl/gl/v4.6-core/gl"
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

	depth float32 = 0.0
)

var (
	square = []float32{
		-0.5, -0.5, depth,
		-0.5, 0.5, depth,
		0.5, 0.5, depth,
		0.5, -0.5, depth,
	}
	squareIdx = []uint{
		0, 1, 2,
		0, 2, 3,
	}
)

func GetSquare() (vbo uint32, pIndexes unsafe.Pointer, err error) {
	var vao, ebo uint32

	pIndexes = gl.Ptr(squareIdx) //unsafe.Pointer(&squareIdx[0])

	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)

	gl.BindVertexArray(vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, int(reflect.TypeOf(square[0]).Size())*len(square), gl.Ptr(square), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, int(reflect.TypeOf(squareIdx[0]).Size())*len(squareIdx), gl.Ptr(squareIdx), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindVertexArray(0)

	return
}

func GetSquareInvalid() (vbo uint32, pIndexes unsafe.Pointer, err error) {
	/// VBO SECTION
	gl.GenBuffers(1, &vbo)
	g.GlErrorHelper()
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	g.GlPanicHelper("1")
	gl.BufferData(gl.ARRAY_BUFFER, int(reflect.TypeOf(square[0]).Size())*len(square), gl.Ptr(square), gl.STATIC_DRAW)
	g.GlPanicHelper("2")

	// gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	// g.GlPanicHelper("3")
	// // setup shader input args??
	// gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	// g.GlPanicHelper("4")
	// gl.EnableVertexAttribArray(0)
	// g.GlPanicHelper("5")

	/// VAO SECTION
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	//g.GlErrorHelper()
	g.GlPanicHelper("11")
	gl.BindVertexArray(vao)
	g.GlPanicHelper("12")
	gl.EnableVertexAttribArray(0)
	g.GlPanicHelper("13")
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	g.GlPanicHelper("14")
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	g.GlPanicHelper("15")

	/// EBO SECTION
	// here we would store buffer name (aka descriptor or reference)
	var ebo uint32
	// get unique name, reserve it and store in variable
	gl.GenBuffers(1, &ebo)
	g.GlErrorHelper()
	// set type of buffer and bind it to current gl context. This allow to change buffer's internal state
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	// change buffer internal state (mutable storage)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, int(reflect.TypeOf(squareIdx[0]).Size())*len(squareIdx), gl.Ptr(squareIdx), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)

	pIndexes = gl.Ptr(squareIdx)

	err = g.GlErrorHelper()

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

func InitElements(program uint32) (err error) {
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

	return
}
