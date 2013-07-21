package goray

import "math"

var Infinity float64 = float64(math.Inf(1))

type Vec3 struct {
	x, y, z float64
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

func (v Vec3) scalarMultiply( a float64 ) Vec3 {
	v.x *= a
	v.y *= a
	v.z *= a
	return v
}


func (v *Vec3) LengthSqrt() float64 {
	result := v.x * v.x
	result += v.y * v.y
	result += v.z * v.z
	return result
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSqrt())
}

func (v Vec3) Dot( v2 Vec3) float64 {
	result := v.x * v2.x
	result += v.y * v2.y
	result += v.z * v2.z
	return result
}

func normalize( v Vec3) Vec3 {
	return  v.scalarMultiply( 1.0 / math.Sqrt(v.Dot(v)))
}

type Point struct {
	x, y, z float64
}

func (p Point) Sub( a Point) Vec3 {
	x := p.x - a.x
	y := p.y - a.y
	z := p.z - a.z
	return Vec3{x, y, z}
}

func (p Point) Add( a Vec3 ) Point {
	x := p.x + a.x
	y := p.y + a.y
	z := p.z + a.z
	return Point{x,y,z}
}
