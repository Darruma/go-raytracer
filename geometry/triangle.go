package geometry

type Triangle struct {
	V0 Vector3
	V1 Vector3
	V2 Vector3
	mat Material
}

func (t Triangle) RayIntersect( ay Ray) Vector3 {
	// Find point P where Ray intersects the Triangle
}
func (t Triangle) Normal() {
	A := t.V1.Subtract(V0)
	B := t.V2.Subtract(V0)
	return A.Cross(B)
}