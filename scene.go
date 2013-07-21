package goray

type Light struct {
	origin Point
}

type Scene struct {
	lights []Light
	shapes []Shapes
}

func (s *Scene) RayTrace(r Ray) Colour {

}