package f

import (
	"image"
	_ "image/png"
	"io"
	"os"
	"tilinghelper/pkg/h"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func AddTexture(window *glfw.Window, path string) (texture uint32, err error) {
	h.ShowThreadId()

	defer glfw.DetachCurrentContext()
	window.MakeContextCurrent()

	//image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	fd, err := os.Open(path)
	if err != nil {
		return
	}
	defer fd.Close()

	data, width, height, err := getData(fd)
	if err != nil {
		return
	}

	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, width, height, int32(0), gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&data[0]))
	gl.BindTexture(gl.TEXTURE_2D, 0)

	return
}

func getData(fd io.Reader) (data []byte, width, height int32, err error) {
	img, _, err := image.Decode(fd)
	if err != nil {
		return
	}

	bounds := img.Bounds()
	width, height = int32(bounds.Max.X), int32(bounds.Max.Y)

	for y := range make([]int, height) {
		for x := range make([]int, width) {
			r, g, b, _ := img.At(x, y).RGBA()
			data = append(data, uint8(r))
			data = append(data, uint8(g))
			data = append(data, uint8(b))
		}
	}

	return
}
