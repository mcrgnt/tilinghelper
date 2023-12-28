package p

import (
	"fmt"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func CreateProg() (program uint32) {
	program = gl.CreateProgram()
	return
}

func Validate(program uint32) {
	var params int32 = 0
	gl.GetProgramiv(program, gl.VALIDATE_STATUS, &params)
	fmt.Println("Validate:", program, params)
}

func LinkStatus(program uint32) {
	var isLinked int32 = 0
	gl.GetProgramiv(program, gl.LINK_STATUS, &isLinked)
	fmt.Println("LinkStatus:", program, isLinked)

	var maxLength int32 = 0
	gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &maxLength)

	// The maxLength includes the NULL character
	var infoLog uint8 = 0
	gl.GetProgramInfoLog(program, maxLength, &maxLength, &infoLog)
	fmt.Println("Log:", program, infoLog)

	// The program is useless now. So delete it.
	gl.DeleteProgram(program)
}

func GetPrograms(window *glfw.Window) (programs []uint32) {
	defer glfw.DetachCurrentContext()
	window.MakeContextCurrent()

	programs = []uint32{}
	programs = append(programs, gl.CreateProgram())

	return
}
