package main

import (
	"fmt"
	"runtime"
	"time"

	"tilinghelper/src/e"
	"tilinghelper/src/g"
	"tilinghelper/src/h"
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
	E  uint32
}

func main() {
	var prefix int
	for range wCount {
		name := fmt.Sprintf("window %v", prefix)
		prefix++

		c := &ctx{}
		c.W, err = w.GetWindow(wWidth, wHeight, name)
		if err != nil {
			h.PanicError("get window", err)
		}
		g.InitDebug(c.W, name)

		c.PP = p.GetPrograms(c.W)
		for _, p := range c.PP {
			err = e.InitElements(c.W, p)
			if err != nil {
				h.PanicError("elements", err)
			}
		}

		c.V, err = e.GetVao(c.W)
		if err != nil {
			h.PanicError("get vao", err)
		}

		c.E, err = e.GetEbo(c.W)
		if err != nil {
			h.PanicError("get e", err)
		}

		cc = append(cc, c)
	}

	fmt.Println("check point")

	for {
		for i := 0; i < len(cc); i++ {
			if cc[i].W.ShouldClose() {
				cc[i].W.Destroy()
				cc = append(cc[:i], cc[i+1:]...)
				i--
				continue
			}

			counterChInc <- true

			cc[i].W.MakeContextCurrent()

			err = w.DrawRectangle(true, false, cc[i].W, cc[i].PP[0], cc[i].E)
			if err != nil {
				h.PanicError("draw rectangle:", err)
			}

			err = w.DrawTriangle(false, true, cc[i].W, cc[i].PP[0], cc[i].V)
			if err != nil {
				h.PanicError("draw triangle:", err)
			}

			// err = w.DrawTriangle(true, false, cc[i].W, cc[i].PP[0], cc[i].V)
			// if err != nil {
			// 	h.PanicError("draw triangle:", err)
			// }

			// err = w.DrawRectangle(false, true, cc[i].W, cc[i].PP[0], cc[i].E)
			// if err != nil {
			// 	h.PanicError("draw rectangle:", err)
			// }

			glfw.DetachCurrentContext()
		}
		if len(cc) == 0 {
			break
		}
	}
}

func init() {
	runtime.LockOSThread()
}

var (
	counter      int
	counterChInc = make(chan bool)
)

func init() {
	go func() {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-counterChInc:
				counter++
			case <-t.C:
				fmt.Println("FPS: ", counter)
				counter = 0
			}
		}
	}()
}
