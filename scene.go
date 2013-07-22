package goray

type Light struct {
	origin Point
}

type Scene struct {
	lights []Light
	shapes []Shapes
}

func (s *Scene) RayTrace(r Ray) Colour {

	if float64(r.depth) >= Infinity {
		return Colour{0,0,0}
	}

	hit := Hit{Infinity, Point{0,0,0},nil}

	for _, shape := range s.shapes {
		h := shape.Intersect(r)
		if h.distance < hit.distance {
			hit = h
		}

	}
}