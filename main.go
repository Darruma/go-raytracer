package main

import (
	. "go-raytracer/geometry"
	"math"
	"os"
    "fmt"
	"strconv"
)

func main() {
	sphere := Sphere{Vector3{-3, 0, -16}, 1, Vector3{123, 25, 60}}
	sphere2 := Sphere{Vector3{-1, 0, -16}, 2, Vector3{70, 230, 4}}
    sphere3 := Sphere{Vector3{-5,4,-10},3,Vector3{56,232,67}}
    light := Light{Vector3{-6,6,7},100}
    os := []Object{sphere, sphere2,sphere3}
	lights := []Light{light}
	render(os, 1024, 768, lights)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func cast_ray(ray Ray, o []Object, lights []Light) Vector3 {
	object, intersect, hit := intersection(ray, o)

	if intersect {
	normal := object.GetNormal()
    fmt.Println(normal)
		var diffuse_lighting float64
		for j := 0; j < len(lights); j++ {
			lightDirection := lights[j].Position.Normalise().Subtract(hit).Normalise()
			fmt.Print("light_calculation ")
            fmt.Print(lightDirection.Dot(normal))
            diffuse_lighting = diffuse_lighting + (lights[j].Intensity * lightDirection.Dot(normal))
            fmt.Println(diffuse_lighting)
		}
		return object.GetMaterial().Scale(diffuse_lighting)
	}
	return Vector3{10, 10, 10}

}

func intersection(ray Ray, objects []Object) (Object, bool, Vector3) {
	var closest_dist float64 = math.MaxFloat64
	var closest_object Object
	var hit_vector Vector3
	for i := 0; i < len(objects); i++ {
		intersect, v0 := objects[i].Intersect(ray)
        hit_vector = v0
		if intersect {
			dist_i := v0.Distance(ray.Origin)
			if dist_i < closest_dist {
				closest_object = objects[i]
				closest_dist = dist_i
			}
		}
	}
	return closest_object, closest_dist < 1000, hit_vector
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
			_, err = file.WriteString(strconv.Itoa(int(pixel.X)) + " " +strconv.Itoa(int(pixel.Y)) + " " +strconv.Itoa(int(pixel.Z)) + "\n")
		}
	}
}
