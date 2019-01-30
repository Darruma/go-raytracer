package main

import (
  "math"
	"fmt"
	. "go-raytracer/geometry"
  "os"
)

func main() {
  
	render()
}

func cast_ray (ray Ray , s Sphere) Vector3 {
  if s.Intersect(ray) == false {
      return Vector3{40,40,40}
  }
  return Vector3{0,0,0}

}
func render() {
 var width int = 1920
 var height int = 1080
 var buffer = make([]Vector3, width*height)
 sp := Sphere{Vector3{3, 3, 3}, 5}
 file,err := os.Create("output.ppm");
 for i:=0 ; i < height; i++ {
    for j:=0; j < width; j++ {
       var ray_x float64 =  (2 * ((i + 0.5)/width) -1) * math.Tan(45) * width / height
       var ray_y float64 = -(2 * ((j + 0.5)/height -1) * math.Tan(45))
       var direction Vector3 = Vector3{ray_x,ray_y,-1}.Normalise()
       buffer[i + j * width] = cast_ray(Ray{Vector3{0,0,0},direction}, sp)
     }
 }
 fmt.Println(buffer)
}