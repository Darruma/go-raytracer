package geometry

type Object interface {
	Intersect(Ray) (bool, float64, float64)
	GetMaterial() Material
}
