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

func (renderer *Renderer) CreateScene() {

	s1 := new(Sphere)
	s1.center = Point{-4, 0, 10}
	s1.radius = 2

	s2 := new(Sphere)
	s2.center = Point{4, 0, 10}
	s2.radius = 3

	spheres := make([]Shapes, 2)

	spheres[0] = s1
	spheres[1] = s2

	l1 := Light{Point{0.0, 5, 5}}

	lights := make([]Light, 1)
	lights[0] = l1

	renderer.scene = &Scene{lights, spheres}

}

func (renderer *Renderer) CreateCamera(){
	renderer.cam = &Camera{Point{0, 0, -10}, 100, 100}
}

func (renderer *Renderer) RenderScene() *image.RGBA {
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
