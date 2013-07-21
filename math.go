package goray

import "math"

type Vec3 struct {
	x, y, z float32
}

func (v Vec3) Add( a Vec3) Vec3 {
	v.x += a.x
	v.y += a.y
	v.z += a.z

	return v
}

func (v Vec3) Sub( a Vec3) Vec3 {
	v.x -= a.x
	v.y -= a.y
	v.z -= a.z
	return v
}

func (v Vec3) scalarMultiply( a float32 ) Vec3 {
	v.x *= a
	v.y *= a
	v.z *= a
	return v
}


func (v *Vec3) LengthSqrt() float32 {
	result := v.x * v.x
	result += v.y * v.y
	result += v.z * v.z
	return result
}

func (v *Vec3) Length() float32 {
	return math.Sqrt(v.LengthSqrt())
}

func (v Vec3) Dot( v2 Vec3) float32 {
	result := v.x * v2.x
	result += v.y * v2.y
	result += v.z * v2.z
	return result
}

func normalize( v Vec3) Vec3 {
	return  v.scalarMultiply( 1.0 / math.Sqrt(v.Dot(v)))
}

type Point struct {
	x, y, z float32
}

func (p Point) sub( a Point) Vec3 {
	result := p.x - a.x
	result = p.y - a.y
	result = p.z - a.z
	return result
}
