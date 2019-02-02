package geometry


type Sphere struct {
	Center Vector3
        Radius float64
        Mat Material
}

func (s Sphere) Intersect(r Ray) bool {
     L := r.Origin.Subtract(s.Center)
     a := r.Direction.Dot(r.Direction)
     b := 2 * r.Direction.Dot(L)
     c := L.Dot(L) - s.Radius * s.Radius
     if b * b - 4 * a * c < 0  {
             return false
     }
     return true
}

