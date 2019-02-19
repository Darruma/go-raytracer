package geometry

type Triangle struct {
	V0  Vector3
	V1  Vector3
	V2  Vector3
	Mat Vector3
}

func (t Triangle) Intersect(ray Ray) bool {
	// Find point P where Ray intersects the Triangle
	return false
}
func (t Triangle) GetNormal() Vector3 {
	A := t.V1.Subtract(t.V0)
	B := t.V2.Subtract(t.V0)
	return A.Cross(B)
}

func(t Triangle) GetMaterial() Vector3 {
	return t.Mat
}