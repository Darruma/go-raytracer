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

func cast_ray (ray Ray , s Sphere) Material {
  if s.Intersect(ray) == false {
      return [40,40,40]
  }
  return [0,0,0]

}
func render() {
 var width int = 1920
 var height int = 1080
 var buffer = make([]Material, width*height)
 sp := Sphere{Vector3{3, 3, 3}, 5}
 file, err := os.OpenFile("output.ppm", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
 if err != nil {
  fm.Println(err)
}
 defer file.Close()
 for i:=0 ; i < height; i++ {
    for j:=0; j < width; j++ {
       var ray_x float64 =  (2 * ((i + 0.5)/float64(width) -1) * math.Tan(45) * width / float64(height)
       var ray_y float64 = -(2 * ((j + 0.5)/float64(height -1) * math.Tan(45))
       var direction Vector3 = Vector3{ray_x,ray_y,-1}.Normalise()
       buffer[i + j * width] = cast_ray(Ray{Vector3{0,0,0},direction}, sp)
       _,err = file.Write(buffer[i+j * width]); 
       if err != nil {
         fm.Println(err)
       }
     }
 }

 
 
}