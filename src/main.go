package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"tilinghelper/src/e"
	"tilinghelper/src/g"
	"tilinghelper/src/p"
	"tilinghelper/src/w"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	wWidth  int = 800
	wHeight int = 600
	wTitle      = "TilingHelper"
)

var (
	err error

	wCount = make([]byte, 3)
	cc     = []*ctx{}
)

type ctx struct {
	W  *glfw.Window
	PP []uint32
	V  uint32
}

func panicError(source string, err error) {
	println(source, err.Error())
	os.Exit(1)
}

func exitOnTime() {
	time.Sleep(time.Millisecond * 1600)
	os.Exit(0)
}

func main() {
	//go exitOnTime()
	var prefix int
	for range wCount {
		name := fmt.Sprintf("window %v", prefix)
		prefix++

		c := &ctx{}
		c.W, err = w.GetWindow(wWidth, wHeight, name)
		if err != nil {
			panicError("get window", err)
		}
		g.InitDebug(c.W, name)

		c.PP = p.GetPrograms(c.W)
		for _, p := range c.PP {
			err = e.InitElements(c.W, p)
			if err != nil {
				panicError("elements", err)
			}
		}

		c.V, err = e.GetVao(c.W)
		if err != nil {
			panicError("get vao", err)
		}

		cc = append(cc, c)
	}

	for _, c := range cc {
		fmt.Printf("DEBUG: %+v\n", *c)
	}

	time.Sleep(time.Second)

	for {
		for i := 0; i < len(cc); i++ {
			if cc[i].W.ShouldClose() {
				cc[i].W.Destroy()
				cc = append(cc[:i], cc[i+1:]...)
				i--
				continue
			}

			cc[i].W.MakeContextCurrent()
			err = w.DrawTriangle(cc[i].W, cc[i].PP[0], cc[i].V)
			if err != nil {
				panicError("draw:", err)
			}
			glfw.DetachCurrentContext()

			//time.Sleep(time.Millisecond * 1000)
		}
		if len(cc) == 0 {
			break
		}
	}
}

func init() {
	runtime.LockOSThread()
}
