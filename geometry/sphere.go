package geometry

import "fmt"

type Sphere struct {
    Center Vector3
    mat Material
    Radius float64
}

func (s Sphere) Intersect(r Ray , t float64) {
    fmt.Println("Checking for intersections")
}

