package w

import (
	"tilinghelper/src/g"
	"time"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	initGLFW bool
	initGL   bool
)

var (
	iter int = 0

	colors = [][]float32{
		{0.1, 0.3, 1.0, 0.5},
		{0.5, 0.6, 0.7, 1.0},
		{0.1, 0.2, 0.3, 1.0},
	}
)

func DrawTriangle(clear bool, swap bool, window *glfw.Window, program uint32, vao uint32) (err error) {

	defer glfw.PollEvents()

	if clear {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		err = g.GlErrorHelper()
		if err != nil {
			return
		}
		gl.ClearColor(0.1, 0.5, 1.0, 0.5)
	}

	gl.UseProgram(program)
	err = g.GlErrorHelper()
	if err != nil {
		return
	}

	location := gl.GetUniformLocation(program, gl.Str("inDiffuseColor\x00"))
	err = g.GlErrorHelper()
	if err != nil {
		return
	}

	gl.Uniform4f(location, colors[iter][0], colors[iter][1], colors[iter][2], colors[iter][3])
	err = g.GlErrorHelper()
	if err != nil {
		return
	}

	iter++
	if iter == 3 {
		iter = 0
	}

	gl.BindVertexArray(vao)
	err = g.GlErrorHelper()
	if err != nil {
		return
	}

	gl.DrawArrays(gl.TRIANGLES, 0, int32(3))
	err = g.GlErrorHelper()
	if err != nil {
		return
	}

	if swap {
		window.SwapBuffers()
		err = g.GlErrorHelper()
	}

	return
}

func DrawRectangle(clear bool, swap bool, window *glfw.Window, program uint32, ebo uint32) (err error) {
	defer glfw.PollEvents()

	if clear {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		err = g.GlErrorHelper()
		if err != nil {
			return
		}
		gl.ClearColor(0.1, 0.5, 1.0, 0.5)
	}

	gl.UseProgram(program)
	err = g.GlErrorHelper()
	if err != nil {
		return
	}

	location := gl.GetUniformLocation(program, gl.Str("inDiffuseColor\x00"))
	err = g.GlErrorHelper()
	if err != nil {
		return
	}

	gl.Uniform4f(location, colors[iter][0], colors[iter][1], colors[iter][2], colors[iter][3])
	err = g.GlErrorHelper()
	if err != nil {
		return
	}

	iter++
	if iter == 3 {
		iter = 0
	}

	err = g.GlErrorHelper()
	if err != nil {
		return
	}

	gl.BindVertexArray(ebo)
	err = g.GlErrorHelper()
	if err != nil {
		return
	}

	gl.DrawElements(gl.TRIANGLES, int32(6), gl.UNSIGNED_INT, nil)
	err = g.GlErrorHelper()
	if err != nil {
		return
	}

	if swap {
		window.SwapBuffers()
		err = g.GlErrorHelper()
	}

	return
}

func GetWindow(width, height int, title string) (window *glfw.Window, err error) {
	window, err = createWindow(width, height, title)
	return
}

func createWindow(w, h int, title string) (window *glfw.Window, err error) {
	if !initGLFW {
		err = glfw.Init()
		if err != nil {
			return
		}
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	//glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	//glfw.WindowHint(glfw.Maximized, glfw.True)
	glfw.WindowHint(glfw.OpenGLDebugContext, glfw.True)
	//glfw.WindowHint(glfw.Blend, glfw.True)

	window, err = glfw.CreateWindow(w, h, title, nil, nil)
	if err != nil {
		return
	}

	window.SetMouseButtonCallback(mouseCallback)
	window.SetCursorEnterCallback(cursorCallBack)
	window.SetPosCallback(posCallback)
	window.SetDropCallback(dropCallback)
	window.SetScrollCallback(scrollCallback)
	window.SetSizeCallback(sizeCallback)

	//window.SetAspectRatio(16, 9)
	//SetFramebufferSizeCallback

	window.MakeContextCurrent()
	if !initGL {
		err = gl.Init()
		if err != nil {
			return
		}
		gl.Enable(gl.DEPTH_TEST)
	}
	glfw.DetachCurrentContext()

	return
}

func sizeCallback(window *glfw.Window, width int, height int) {
	println("new size:", width, height)
}

func scrollCallback(window *glfw.Window, xoff float64, yoff float64) {
	ctrlState := window.GetKey(glfw.KeyLeftControl)
	x, y := window.GetPos()
	println("scroll old pos:", x, y)
	switch {
	case ctrlState == glfw.Press:
		if yoff > 0 {
			x += 1
			break
		}
		if yoff < 0 {
			x -= 1
		}
	default:
		if yoff > 0 {
			y += 1
			break
		}
		if yoff < 0 {
			y -= 1
		}
	}
	window.SetPos(x, y)
}

func dropCallback(window *glfw.Window, names []string) {
	for _, name := range names {
		println("drop:", name)
	}
}

func posCallback(window *glfw.Window, xpos int, ypos int) {
	println("new pos:", xpos, ypos)
}

var (
	cursorCallBackHelper = map[bool]string{true: "enter:", false: "exit:"}
)

func cursorCallBack(window *glfw.Window, entered bool) {
	x, y := window.GetCursorPos()
	println(cursorCallBackHelper[entered], int(x), int(y))
}

var (
	mouseCallbackHelper = map[glfw.Action]string{0: "release:", 1: "capture:"}
)

func mouseCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	switch button {
	case 0:
		switch action {
		case 0:
		case 1:
			window.Hide()
			go func() {
				time.Sleep(time.Second * 1)
				window.Show()
			}()
		}
	case 1:
		switch action {
		case 1:
			if window.GetAttrib(glfw.Decorated) == 1 {
				window.SetAttrib(glfw.Decorated, glfw.False)
				window.SetAttrib(glfw.Floating, glfw.True)
			} else {
				window.SetAttrib(glfw.Decorated, glfw.True)
				window.SetAttrib(glfw.Floating, glfw.False)
			}
		}
	}
	println("click:", button, mouseCallbackHelper[action], mod)
}

func Terminate() {
	glfw.Terminate()
}
