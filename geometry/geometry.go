package geometry

import "math"

type Vector3 struct {
     X float64
     Y float64
     Z float64
}

func (v Vector3) Dot(w Vector3) float64 {
  return v.X * w.X + v.Y * w.Y + v.Z  * w.Y
}

func (v Vector3) Length() float64 {
  return math.Sqrt(v.Dot(v))
}

func (v Vector3) Add(w Vector3) Vector3 {
  return Vector3{X:v.X + w.X,Y:v.Y + w.Y,Z: v.Z + w.Z}
}

func (v Vector3) Subtract(w Vector3) Vector3 {
  return Vector3{X:v.X - w.X,Y:v.Y - w.Y,Z: v.Z - w.Z}
}

func (v Vector3) Scale(a float64) Vector3 {
  return Vector3{X:a * v.X,Y: a * v.Y, Z: a * v.Z}
}


