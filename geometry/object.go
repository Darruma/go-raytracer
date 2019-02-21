package geometry

type Object interface {
	Intersect(Ray) (bool,Vector3,Vector3)
	GetMaterial() Vector3
    GetID() string
}
