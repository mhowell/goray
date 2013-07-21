package goray 

type Colour struct {
	red float32
	green float32
	blue float32
}

type Camera struct {
	eye Point
	w int
	h int
}

func (c *Camera) rayForPixel( x int, y int) Ray {
	dirX := float32(x) - float32(c.w) * float32(0.5)
	dirY := float32(y) - float32(c.h) * float32(0.5)
	dirZ := float32(c.w)

	dir := normalize(Vec3{dirX, dirY, dirZ})

	return Ray{c.eye, dir}
}