package geometry

import "math"

type Sphere struct {
	Center Vector3
	Radius float64
}

func (s Sphere) Intersect(r Ray) bool {
	var origin_to_center Vector3 = s.Center.Subtract(r.Origin)
	var origin_to_center_middle float64 = origin_to_center.Dot(r.Direction)
	if origin_to_center_middle < 0 {
		return false
	}
	otc_squared := origin_to_center.Dot(origin_to_center)
	otcm_squared := origin_to_center_middle * origin_to_center_middle
  ctcm := math.Sqrt(otc_squared - otcm_squared)
	if ctcm < 0 {
		return false
	}
	//intersect_to_cm := math.Sqrt(s.Radius * s.Radius - ctcm * ctcm)
  //t_intersect_1 := origin_to_center_middle - intersect_to_cm
	//t_intersect_2 := origin_to_center_middle + intersect_to_cm
	//fmt.Println(t_intersect_1);
	//fmt.Println(t_intersect_2);
	return true
}
