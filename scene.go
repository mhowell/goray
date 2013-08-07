package goray

import (
	"image/color"
	"log"
)

type Light struct {
	origin Point
}

type Scene struct {
	lights []Light
	shapes []Shapes
}

func (s *Scene) RayTrace(r Ray) color.RGBA {

	if float64(r.depth) >= Infinity {
		return color.RGBA{0, 0, 0, 0}
	}

	hit := Hit{Infinity, Point{0, 0, 0}, nil}

	for _, shape := range s.shapes {
		h := shape.Intersect(r)
		if h.distance < hit.distance {
			hit = h
			log.Print(h)
		}

	}
	if hit.distance < Infinity {
		return color.RGBA{255, 0, 255, 0}
	}
	return color.RGBA{0, 0, 0, 0}
}
