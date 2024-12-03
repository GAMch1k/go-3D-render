package sphere

import (
	"math"

	"gamch1k.org/render/prefabs/3d/vector3"
)

type Sphere struct {
	vector3.Vector3
	Radius float64
}

func New(v3 vector3.Vector3, r float64) *Sphere {
	return &Sphere{v3, r}
}

func (s *Sphere) Vector() *vector3.Vector3 {
	return &s.Vector3
}

func (s *Sphere) Render(cam *vector3.Vector3, ray_dir *vector3.Position) float64 {

	oc := cam.Substract(&s.Vector3)
	a := ray_dir.DotProduct(ray_dir)
	b := 2.0 * oc.DotProduct(ray_dir)
	c := oc.Position.DotProduct(&oc.Position) - s.Radius*s.Radius

	discr := b*b - 4*a*c

	if discr >= 0 {
		t1 := -b - math.Sqrt(discr) / (2 * a)
		t2 := -b + math.Sqrt(discr) / (2 * a)


		if t1 <= t2 {
			return t1
		} 

		return t2

	}

	return math.Inf(1)
}
