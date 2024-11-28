package sphere

import (
	"math"

	"gamch1k.org/render/prefabs/3d/vector3"
)

type Sphere struct {
	vector3.Vector3
	Radius float64
}

func (s Sphere) Vector() *vector3.Vector3 {
	return &s.Vector3
}

func (s Sphere) Render(cam *vector3.Vector3, pixel *vector3.Vector3) float64 {

	ray := pixel.Substract(cam)
	ray_dir := ray.Position.Normalize()

	oc := cam.Substract(&s.Vector3)
	a := ray_dir.Dot(&ray_dir)
	b := 2.0 * oc.Dot(&ray_dir)
	c := oc.Position.Dot(&oc.Position) - s.Radius*s.Radius

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
