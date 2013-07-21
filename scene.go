package goray

type Light struct {
	origin Point
}

type Scene struct {
	lights []Light
	shapes []Shapes
}

func (s *Scene) RayTrace(r Ray) Colour {

	if r.depth >= Infinity {
		return Colour{0,0,0}
	}

	for shape := range s.shapes {

	}
}