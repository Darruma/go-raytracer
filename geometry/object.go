package geometry

type Object interface {
	Intersect(Ray) (bool, Vector3)
	GetMaterial() Vector3
	GetType() string
	GetNormal() Vector3
}
