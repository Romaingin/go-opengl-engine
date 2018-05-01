package engine

import (
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/rginestou/check"
)

func init() {
	// Register format here
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

// TextureFromFile ...
func TextureFromFile(filePath string) uint32 {
	file, err := os.Open(filePath)
	check.Panic(err)
	defer file.Close()

	img, _, err := image.Decode(file)
	check.Panic(err)

	rect := img.Bounds()
	rgba := image.NewRGBA(rect)
	draw.Draw(rgba, rect, img, rect.Min, draw.Src)

	// Create one OpenGL texture
	var textureID uint32
	println("heuyyyye")
	gl.GenTextures(1, &textureID)
	println("heuyyyye2")

	// "Bind" the newly created texture : all future texture functions will modify this texture
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, textureID)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	width := int32(rect.Max.X - rect.Min.X)
	height := int32(rect.Max.Y - rect.Min.Y)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))

	return textureID
}
