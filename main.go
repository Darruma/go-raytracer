package main

import (
	"fmt"
	. "go-raytracer/geometry"
	"math"
	"os"
	"strconv"
)

func main() {
	sphere := Sphere{Vector3{-1, -0.5, -10}, 3, Vector3{123, 25, 60}}
	light := Light{Vector3{-20, 20, 20}, 0.6}
	os := []Object{sphere}
	lights := []Light{light}
	render(os, 1024, 768, lights)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func cast_ray(ray Ray, o []Object, lights []Light) Vector3 {
	object, intersect, hit, normal := intersection(ray, o)
	if intersect {
		var diffuse_lighting float64
		for j := 0; j < len(lights); j++ {
			lightDirection := lights[j].Position.Subtract(hit).Normalise()
			lightNormal := lightDirection.Dot(normal)
            diffuse_lighting += lights[j].Intensity * math.Abs(lightNormal)
		}
		fmt.Println(diffuse_lighting)
		return object.GetMaterial().Scale(diffuse_lighting)
	}
	return Vector3{10, 10, 10}

}

func intersection(ray Ray, objects []Object) (Object, bool, Vector3, Vector3) {
	var closest_dist float64 = math.MaxFloat64
	var closest_object Object
	var normal Vector3
	var v0 Vector3
	var intersect bool
	for i := 0; i < len(objects); i++ {
		intersect, v0, normal = objects[i].Intersect(ray)
		if intersect {
			dist_i := v0.Distance(ray.Origin)
			if dist_i < closest_dist {
				closest_object = objects[i]
				closest_dist = dist_i
			}
		}
	}
	return closest_object, closest_dist < 1000, v0, normal
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
