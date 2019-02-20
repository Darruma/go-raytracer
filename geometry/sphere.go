package geometry

import "math"

type Sphere struct {
	Center Vector3
	Radius float64
	Mat    Vector3
    Normal Vector3
}

func (s Sphere) Intersect(r Ray) (bool, Vector3) {
	L := r.Origin.Subtract(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(L)
	c := L.Dot(L) - s.Radius*s.Radius
	determinant := b*b - 4*a*c
	if determinant < 0 {
		return false, Vector3{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64}
	}
	t0 := -b + math.Sqrt(determinant)/2*a
	t1 := -b - math.Sqrt(determinant)/2*a
	if r.Point(t0).Distance(r.Origin) <  r.Point(t1).Distance(r.Origin) {
		return true, r.Point(t0)
	} else {
		return true,r.Point(t1)
}
}

func (s Sphere) GetMaterial() Vector3 {
	return s.Mat
}

func (s Sphere) GetNormal() Vector3 {
    return s.Normal
}

func (s Sphere) GetType() string {
	return "Sphere"
}
