package goray

type Shapes interface {
	Intersect(r Ray) Hit
}

type Sphere struct {
	center Point
	radius float32
}

func (s *Sphere) Intersect(r Ray) Hit {
	v := s.center.sub( r.orig)
	b := v.Dot( r.dir )

}