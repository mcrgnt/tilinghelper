package w

import (
	"fmt"
	"tilinghelper/src/g"
	"time"
	"unsafe"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func Terminate() {
	glfw.Terminate()
}

func GetWindows(count, width, height int, title string) (windows []*glfw.Window, err error) {
	windows = []*glfw.Window{}
	for i := 0; i < count; i++ {
		var window *glfw.Window
		window, err = createWindow(width, height, fmt.Sprintf("%s_%v", title, i))
		if err != nil {
			return
		}
		windows = append(windows, window)
	}

	return
}

func DrawSquare(window *glfw.Window, program uint32, vao uint32, pIndexes unsafe.Pointer) {
	fmt.Println("???")
	g.GlPanicHelper("66")
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	g.GlPanicHelper("6")
	gl.UseProgram(program)
	fmt.Println("?????")
	g.GlPanicHelper("5")
	gl.BindVertexArray(vao)
	fmt.Println("????")
	g.GlPanicHelper("4")
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, pIndexes)
	fmt.Println("??")
	g.GlPanicHelper("3")
	window.SwapBuffers()
	g.GlPanicHelper("2")
	fmt.Println("?")
	glfw.PollEvents()
	g.GlPanicHelper("1")
	fmt.Println("!!!")
}

func DrawTriangle(window *glfw.Window, program uint32, vao uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	//g.GlPanicHelper("6")

	gl.UseProgram(program)
	//g.GlPanicHelper("5")

	gl.BindVertexArray(vao)
	//g.GlPanicHelper("4")

	gl.DrawArrays(gl.TRIANGLES, 0, int32(3))
	//g.GlPanicHelper("3")

	window.SwapBuffers()
	//g.GlPanicHelper("2")

	glfw.PollEvents()
	//g.GlPanicHelper("1")
}

func createWindow(w, h int, title string) (window *glfw.Window, err error) {
	if err = glfw.Init(); err != nil {
		return
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	//glfw.WindowHint(glfw.Maximized, glfw.True)
	glfw.WindowHint(glfw.OpenGLDebugContext, glfw.True)

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
