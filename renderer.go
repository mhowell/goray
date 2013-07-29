package goray

import (
	"image"
)

type Ray struct {
	orig  Point
	dir   Vec3
	depth int
}

type Hit struct {
	distance float64
	pos      Point
	shape    Shapes
}

type Renderer struct {
	scene *Scene
	cam   *Camera
}

func (renderer *Renderer) renderScene() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, renderer.cam.h, renderer.cam.w))

	for y := 0; y < renderer.cam.h; y++ {
		for x := 0; x < renderer.cam.w; x++ {
			ray := renderer.cam.rayForPixel(x, y)
			colour := renderer.scene.RayTrace(ray)
			img.Set(x, y, colour)
		}
	}
	return img
}
