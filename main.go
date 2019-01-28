package main

import "fmt"

type Vector3 struct {
     x float32
     y float32
     z float32
}

func main() {
     vec := Vector3{x:10.2,y:3.5,z:5.2}
     fmt.Println(vec.x)
     fmt.Println(vec.y)
     fmt.Println(vec.z)
}
