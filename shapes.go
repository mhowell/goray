package goray

import (
	"math"
	"log"
)

type Shapes interface {
	Intersect(r Ray) Hit
}

type Sphere struct {
	center Point
	radius float64
}

func (s Sphere) Intersect(r Ray) Hit {
	dist := s.center.Sub(r.orig)
	b := dist.Dot(r.dir)
	d := b*b - dist.Dot(dist) + s.radius*s.radius
	if d < 0.0 {
		return Hit{Infinity, Point{0, 0, 0}, nil}
	}
	d = math.Sqrt(d)
	t1 := b - d
	t2 := b + d

	log.Print(t1)
	log.Print(t2)

	if t1 > 0.1 {
		return Hit{t1, r.orig.Add(s.center.SubVec3(r.dir.scalarMultiply(t1))), Shapes(s)}
	}

	if t2 > 0.1 {
		return Hit{t2, r.orig.Add(s.center.SubVec3(r.dir.scalarMultiply(t2))), Shapes(s)}
	}
	return Hit{Infinity, Point{0, 0, 0}, nil}
}
