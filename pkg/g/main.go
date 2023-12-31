package g

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unsafe"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func InitDebug(window *glfw.Window, name string) {
	defer glfw.DetachCurrentContext()
	window.MakeContextCurrent()

	gl.DebugMessageCallback(debugMessageCallback, nil)
}

func debugMessageCallback(
	source uint32,
	gltype uint32,
	id uint32,
	severity uint32,
	length int32,
	message string,
	userParam unsafe.Pointer) {
	fmt.Println("D:", message)
}

func getErrors() (err error) {
	errs := []string{}
	var e uint32
	for {
		e = gl.GetError()
		if e == 0 {
			break
		}
		errs = append(errs, fmt.Sprintf("%v", gl.GetString(e)))
	}

	if len(errs) != 0 {
		err = errors.New(strings.Join(errs, ";"))
	}

	return
}

func GlErrorHelper() (err error) {
	err = getErrors()
	return
}

func GlPanicHelper(s string) {
	err := getErrors()
	if err != nil {
		fmt.Printf("%s: %s", s, err.Error())
		os.Exit(1)
	}
}
