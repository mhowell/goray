package goray

import "math"

type Vec3 struct {
	x, y, z float32
}

func (v *Vec3) add( a Vec3) {
	v.x += a.x
	v.y += b.y
	v.z += b.z
}


func (v *Vec3) LengthSqrt() float32 {
	result := v.x * v.x
	result += v.y * v.y
	result += v.z * v.z
	return result
}

func (v *Vec3) Length() float32 {
	return math.sqrt(v.LenghtSqrt())
}

func (v *Vec3) Dot( v2 *Vec3) flaot32 {
	result := v.x * v2.x
	result += v.y * v2.y
	result += v.z * v2.z
	return result
}

type Point struct {
	x, y, z float32
}