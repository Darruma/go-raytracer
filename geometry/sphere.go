package geometry


type Sphere struct {
	Center Vector3
	Radius float64
}

func (s Sphere) Intersect(r Ray) bool {
     var vpc Vector3 = s.Center.Subtract(r.Origin)
     if vpc.Dot(r.Direction) < 0  {
        if vpc.Length() > s.Radius {
          return false
        } else {
          return true
        }
     } else {
      var pc = s.Center.Project(r.Origin)
       if ( s.Center.Subtract(pc).Length() > s.Radius ) {
            return false
       } else  {
          return true
       }

     }
}

