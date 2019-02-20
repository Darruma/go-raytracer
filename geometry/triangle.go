package geometry

import "math"

type Triangle struct {
	V0  Vector3
	V1  Vector3
	V2  Vector3
	Mat Vector3
}

func (t Triangle) Intersect(ray Ray) (bool, Vector3) {
	N := t.GetNormal()
    Ndr := N.Dot(ray.Direction)
    // if ray and normal are parallel , they dont intersect
	if math.Abs(Ndr) < 0.001 {
		return false,Vector3{0,0,0}
	}
    D := N.Dot(V0)
    t = (N.Dot(ray.Origin) + D) / Ndr
    if t < 0 {
        return false, Vector3{0,0,0}
    }
    Vector3 hit := ray.Point(t)
	return true,Vector3{0,0,0}
}
func (t Triangle) GetNormal() Vector3 {
	A := t.V1.Subtract(t.V0)
	B := t.V2.Subtract(t.V0)
	return A.Cross(B)
}

func(t Triangle) GetMaterial() Vector3 {
	return t.Mat
}
