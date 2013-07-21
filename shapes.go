package goray

import (
	"math"
)

type Shapes interface {
	Intersect(r Ray) Hit
}

type Sphere struct {
	center Point
	radius float64
}

func (s *Sphere) Intersect(r Ray) Hit {
	dist := s.center.sub( r.orig )
	b := dist.Dot( r.dir )
	d := b*b - dist.Dot(dist) + s.radius * s.radius
	if d < 0.0 {
		return Hit{0, Point{0, 0, 0}, nil}
	}
	d = math.Sqrt(d)
	t1 := b - d
	t2 := b + d

	if t1 > 0.1 {
		return Hit{t1, r.orig.Add(r.dir.scalarMultiply(t1).Sub(s.center)), s}
	}

	if t2 > 0.1 {
		return Hit{t2, r.orig.Add(r.dir.scalarMultiply(t2).Sub(s.center)), s}
	}
}