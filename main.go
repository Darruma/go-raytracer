package main


import ("fmt"
        ."go-raytracer/geometry"
)
func main() {
     var width int = 1920;
     var height int= 1080;
     buffer := make([]Vector3,width * height)
     fmt.Println(buffer)
     render()
}

func render() {
  fmt.Println("Rendering")
}
