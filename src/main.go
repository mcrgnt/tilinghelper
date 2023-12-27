package main

import (
	"os"
	"runtime"
	"tilinghelper/src/e"
	"tilinghelper/src/g"
	"tilinghelper/src/p"
	"tilinghelper/src/w"
	"time"
)

const (
	wCount  int = 4
	wWidth  int = 800
	wHeight int = 600
	wTitle      = "TilingHelper"
)

func panicError(source string, err error) {
	println(source, err.Error())
	os.Exit(1)
}

func main() {
	go func() {
		return
		t := time.NewTicker(time.Millisecond * 500)
		for _ = range t.C {
			os.Exit(0)
		}
	}()

	runtime.LockOSThread()

	windows, err := w.GetWindows(wCount, wWidth, wHeight, wTitle)
	if err != nil {
		panicError("windows", err)
	}
	defer w.Terminate()

	err = g.InitGL()
	if err != nil {
		panicError("g.InitGL", err)
	}

	program, err := p.CreateProg()
	if err != nil {
		panicError("program", err)
	}

	// var params int32
	// gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &params)
	println(">>>>>>>>", program)

	err = e.InitElements(program)
	if err != nil {
		panicError("e.InitElements", err)
	}

	vao, err := e.GetVao()
	if err != nil {
		panicError("vao", err)
	}

	// square, pIndexes, err := e.GetSquare()
	// if err != nil {
	// 	panicError("square", err)
	// }

	go func() {
		return
		var (
			err error
			i   int
		)

		for {
			i++
			err = g.GlErrorHelper()
			if err != nil {
				println("ERR:", err.Error())
			}
		}
	}()

	for {
		for i, window := range windows {
			if window.ShouldClose() {
				window.Destroy()
				windows = append(windows[:i], windows[i+1:]...)
				continue
			}
			window.MakeContextCurrent()
			//g.GlPanicHelper("9")
			//w.DrawSquare(window, program, square, pIndexes)
			w.DrawTriangle(window, program, vao)
			//g.GlPanicHelper("8")
		}
		if len(windows) == 0 {
			break
		}
	}
}

func init() {

}
