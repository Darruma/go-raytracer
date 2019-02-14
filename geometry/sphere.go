package geometry

import "math"

type Sphere struct {
	Center Vector3
	Radius float64
	Mat    Material
}

func (s Sphere) Intersect(r Ray) (bool, float64, float64) {
	L := r.Origin.Subtract(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(L)
	c := L.Dot(L) - s.Radius*s.Radius
	determinant := b*b - 4*a*c
	if determinant < 0 {
		return false, 0, 0
	}

	t0 := -b + math.Sqrt(determinant)/2*a
	t1 := -b - math.Sqrt(determinant)/2*a
	return true, t0, t1
}

func (s Sphere) GetMaterial() Material {
	return s.Mat

}
