package goray 

type Ray struct {
	orig Point
	dir Vec3
	depth int
}

type Hit struct {
	distance float64
	pos Point
	shape Shapes
}

type Renderer struct {
	scene *Scene
	cam *Camera
}

func (renderer *Renderer) renderScene() (image Image) {
	for y := 0; y < renderer.cam.h; y++ {
		for x := 0; x < renderer.cam.w; x++ {
			ray := renderer.cam.rayForPixel(x, y)
			colour := renderer.scene.RayTrace(ray)
		}
	}
}