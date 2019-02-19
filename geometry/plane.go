package geometry

import "math"

type Plane struct {
	Center Vector3
	Normal Vector3
}

func (p Plane) Intersect(r Ray) (bool, Vector3, Vector3) {
	denom := p.Normal.Normalise().Dot(r.Direction.Normalise())
	if math.Abs(denom) > 0.000001 {
		t := p.Center.Subtract(r.Origin).Dot(p.Normal) / denom
		if t > 0 {
			return true, r.Point(t), Vector3{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64}
		}
	}

	return false, Vector3{0, 0, 0}, Vector3{0, 0, 0}
}


func (p Plane) GetType() string {
	return "Plane"
}

func (p Plane) GetNormal() Vector3 {
	return p.Normal
}
