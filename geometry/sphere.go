package geometry

import "math"

type Sphere struct {
	Center Vector3
	Radius float64
	Mat    Vector3
    ID string
}

func (s Sphere) Intersect(r Ray) (bool, Vector3,Vector3) {
	L := r.Origin.Subtract(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(L)
	c := L.Dot(L) - s.Radius*s.Radius
	determinant := b*b - 4*a*c
	if determinant < 0 {
		return false, Vector3{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},Vector3{0,0,0}
	}
	t0 := -b + math.Sqrt(determinant)/2*a
	t1 := -b - math.Sqrt(determinant)/2*a
	v0 := r.Point(t0)
	v1 := r.Point(t1)
	var hit Vector3
	if v0.Distance(r.Origin) < v1.Distance(r.Origin) {
		hit = v0
	} else {
		hit = v1
	}
    return true,hit,hit.Subtract(s.Center).Normalise()
}

func (s Sphere) GetMaterial() Vector3 {
	return s.Mat
}

func (s Sphere) GetID() string {
	return s.ID
}
