package cube

import (
	"math"

	"gamch1k.org/render/prefabs/3d/vector3"
)

type Cube struct {
	vector3.Vector3
	Size float64
	Vertices [8]Vertice
}

type Vertice struct {
	Position vector3.Vector3
}

func (c Cube) Vector() *vector3.Vector3 {
	return &c.Vector3
}

func (c Cube) Render(cam *vector3.Vector3, pixel *vector3.Vector3) float64 {

	halfSize := c.Size / 2.0
	
	cube_min := vector3.Position{
		X: c.Position.X - halfSize,
		Y: c.Position.Y - halfSize,
		Z: c.Position.Z - halfSize,
	}

	cube_max := vector3.Position{
		X: c.Position.X + halfSize,
		Y: c.Position.Y + halfSize,
		Z: c.Position.Z + halfSize,
	}

	ray := pixel.Substract(cam)
	ray_dir := ray.Position.Normalize()

	var t_min_x, t_max_x, t_min_y, t_max_y, t_min_z, t_max_z float64

	t1_x := (cube_min.X - cam.Position.X) / ray_dir.X
	t2_x := (cube_max.X - cam.Position.X) / ray_dir.X
	t_min_x = math.Min(t1_x, t2_x)
	t_max_x = math.Max(t1_x, t2_x)
	


	t1_y := (cube_min.Y - cam.Position.Y) / ray_dir.Y
	t2_y := (cube_max.Y - cam.Position.Y) / ray_dir.Y
	t_min_y = math.Min(t1_y, t2_y)
	t_max_y = math.Max(t1_y, t2_y)
	

	t1_z := (cube_min.Z - cam.Position.Z) / ray_dir.Z
	t2_z := (cube_max.Z - cam.Position.Z) / ray_dir.Z
	t_min_z = math.Min(t1_z, t2_z)
	t_max_z = math.Max(t1_z, t2_z)

	t_min := math.Max(t_min_x, math.Max(t_min_y, t_min_z))
	t_max := math.Min(t_max_x, math.Min(t_max_y, t_max_z))

	if t_min <= t_max && t_max > 0 {
		return t_min
	}
	
	return math.Inf(-1)
}



