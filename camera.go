package goray

type Camera struct {
	eye Point
	w   int
	h   int
}

func (c *Camera) rayForPixel(x int, y int) Ray {

	dir := Vec3{0, 0, 1}
	eye := Point{float64(x), float64(y), -1000}

	return Ray{eye, dir, 0}
}
