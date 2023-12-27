package p

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

func CreateProg() (program uint32, err error) {

	program = gl.CreateProgram()

	return
}
