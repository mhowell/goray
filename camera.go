package goray 

type Camera struct {
	eye Point
	w int
	h int
}

func (c *Camera) rayForPixel( x int, y int) Ray {
	dirX := float64(x) - float64(c.w) * float64(0.5)
	dirY := float64(y) - float64(c.h) * float64(0.5)
	dirZ := float64(c.w)

	dir := normalize(Vec3{dirX, dirY, dirZ})

	return Ray{c.eye, dir, 0}
}