package goray

type Vec3 struct {
	x, y, z float64
}

func (v *Vec3) add( a Vec3) {
	v.x += a.x
	v.y += b.y
	v.z += b.z
}