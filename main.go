package main

import (
	"fmt"
	. "go-raytracer/geometry"
)

func main() {
	var width int = 1920
	var height int = 1080
	buffer := make([]Vector3, width*height)
	sp := Sphere{Vector3{3, 3, 3}, 5}
	fmt.Println(buffer)
	render()
}

func render() {
	fmt.Println("Rendering")

}
