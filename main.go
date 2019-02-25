package main

import (
	"fmt"
	. "go-raytracer/geometry"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println("rendering")
	light := Light{Vector3{-20, 20, 20}, 1.1}
  sphere := Sphere{Vector3{0,-5,-8},2.2,Vector3{60,160,59},"sphere1"}
  spher2 := Sphere{Vector3{-3, -7, -10}, 2, Vector3{193, 96, 29}, "sphere2"}
	os := []Object{sphere,spher2}
	lights := []Light{light}
	render(os, 1024, 768, lights)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func cast_ray(ray Ray, o []Object, lights []Light) Vector3 {
	color, intersect, hit, normal := intersection(ray, o)
	if intersect == true {
		var diffuse_lighting float64
		for j := 0; j < len(lights); j++ {
			lightDirection := lights[j].Position.Subtract(hit).Normalise()
			lightNormal := lightDirection.Dot(normal)
			diffuse_lighting += lights[j].Intensity * math.Abs(lightNormal)
		}
    fmt.Println(diffuse_lighting)
		return color.Scale(diffuse_lighting)
	} else {
		return Vector3{10, 10, 10}
	}

}

func intersection(ray Ray, objects []Object) (Vector3, bool, Vector3, Vector3) {
	object_distance := math.MaxFloat64
	var color Vector3
	var hit_final Vector3
	var normal_final Vector3
	for i := 0; i < len(objects); i++ {
		var dist_i float64 = 0
		intersected, hit, normal := objects[i].Intersect(ray)
		if intersected {
			if dist_i < object_distance {
				object_distance = dist_i
				color = objects[i].GetMaterial()
				hit_final = hit
				normal_final = normal
			}
		}
	}
	return color, object_distance < 1000, hit_final, normal_final
}

func render(objects []Object, width int, height int, lights []Light) {
	var fov = math.Pi / 2
	var buffer = make([]Vector3, width*height)
	file, err := os.OpenFile("output.ppm", os.O_CREATE|os.O_WRONLY, 0644)
	_, err2 := file.WriteString("P3\n" + strconv.Itoa(width) + " " + strconv.Itoa(height) + "\n255\n")
	check(err)
	check(err2)
	defer file.Close()
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var ray_y float64 = (2*(float64(i)+0.5)/float64(width) - 1) * math.Tan(fov/2) * float64(width) / float64(height)
			var ray_x float64 = -(2*(float64(j)+0.5)/float64(height) - 1) * math.Tan(fov/2)
			var direction Vector3 = Vector3{ray_x, ray_y, -1}.Normalise()
			var pixel Vector3 = cast_ray(Ray{Vector3{0, 0, 0}, direction}, objects, lights)
			buffer[j+i*width] = pixel
		}
	}
	for k := 0; k < height; k++ {
		for l := 0; l < width; l++ {
			var pixel Vector3 = buffer[l+k*width]
			_, err = file.WriteString(strconv.Itoa(int(pixel.X)) + " " + strconv.Itoa(int(pixel.Y)) + " " + strconv.Itoa(int(pixel.Z)) + "\n")
		}
	}
}
