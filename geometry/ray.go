package geometry

type Ray struct {
	Origin    Vector3
	Direction Vector3
}

func (r Ray) Point(t float64) Vector3 {
	return r.Origin.Add(r.Direction.Scale(t))
}
