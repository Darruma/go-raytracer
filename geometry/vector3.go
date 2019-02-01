package geometry

import "math"

type Vector3 struct {
     X float64
     Y float64
     Z float64
}

func (v Vector3) Dot(w Vector3) float64 {
  return v.X * w.X + v.Y * w.Y + v.Z  * w.Z
}

func (v Vector3) Cross(w Vector3) Vector3 {
  return Vector3{X:v.Y * w.Z - v.X * w.Y, Y:v.Z * w.X - v.X * w.Z,Z:v.X * w.Y - v.Y * w.X}
}

func (v Vector3) Normalise() Vector3 {
    return v.Scale(1/v.Length())

}
func (v Vector3) Distance(w Vector3) float64 {
    return v.Subtract(w).Length()
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

func (v Vector3) Project(w Vector3) Vector3 {
  return w.Scale(w.Dot(v)/w.Length())

}


