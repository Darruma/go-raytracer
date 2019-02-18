package main

import (
	. "go-raytracer/geometry"
	"math"
	"os"
	"strconv"
)

func main() {
	sphere := Sphere{Vector3{-3, 0, -16}, 1, Material{123, 25, 60}}
	sphere2 := Sphere{Vector3{-1, 0, -16}, 2, Material{70, 230, 4}}
	os := []Object{sphere, sphere2}
	render(os, 1024, 768, lights)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func cast_ray(ray Ray, o []Object, lights []Light) Material {
	object, intersect, hit, normal := intersection(ray, o)
	if intersect {
		var diffuse_lighting float64
		for j := 0; j < len(lights); j++ {
			lightDirection := lights[i].Position.Subtract(hit).Normalise()
			diffuse_lighting = diffuse_lighting + lightDirection.Dot(normal)
		}
		return object.GetMaterial()
	}
	return Material{10, 10, 10}

}

func intersection(ray Ray, objects []Object) (Object, bool, Vector3) {
	var closest_dist float64 = math.MaxFloat64
	var closest_object Object
	var hit_vector Vector3
	var normal Vector3
	for i := 0; i < len(objects); i++ {
		intersect, t0, t1 := objects[i].Intersect(ray)
		//calculate hit points
		if intersect {
			v0 := ray.Point(t0)
			v1 := ray.Point(t1)
			v0_dist := v0.Distance(ray.Origin)
			v1_dist := v1.Distance(ray.Origin)
			dist_i := math.Min(v0_dist, v1_dist)
			if dist_i < closest_dist {
				closest_object = objects[i]
				closest_dist = dist_i
			}
			if v0_dist > v1_dist {
				hit_vector = v0
			} else {
				hit_vector = v1
			}
			normal = hit_vector.Subtract(objects[i].Center).Normalise()
			//calculate distance from hitpoints
		}
	}
	return closest_object, closest_dist < 1000, hit_vector, normal

}
func render(objects []Object, width int, height int, lights []Light) {
	var fov = math.Pi / 2
	var buffer = make([]Material, width*height)
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
			var pixel Material = cast_ray(Ray{Vector3{0, 0, 0}, direction}, objects, lights)
			buffer[j+i*width] = pixel
		}
	}
	for k := 0; k < height; k++ {
		for l := 0; l < width; l++ {
			var pixel Material = buffer[l+k*width]
			_, err = file.WriteString(strconv.Itoa(pixel.R) + " " + strconv.Itoa(pixel.G) + " " + strconv.Itoa(pixel.B) + "\n")
		}
	}
}
