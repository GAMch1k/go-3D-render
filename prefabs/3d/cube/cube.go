package cube

import (
	"math"

	"gamch1k.org/render/prefabs/3d/vector3"
)

type Cube struct {
	vector3.Vector3
	Size float64
	Vertices [8]Vertice
	Sides [12][3]int
}

type Vertice struct {
	vector3.Vector3
}

func New(v3 vector3.Vector3, size float64) *Cube {
	c := &Cube{
		Vector3: v3,
		Size: size,
		Sides: [12][3]int{
			{0, 1, 2}, {2, 3, 0},
			{1, 0, 4}, {4, 5, 1},
			{2, 3, 7}, {7, 6, 2},
			{5, 4, 7}, {7, 6, 5},
			{2, 1, 5}, {5, 6, 2},
			{0, 4, 7}, {7, 3, 0},
		},
	}

	c.GenerateVertices()
	return c
}

func (c *Cube) Vector() *vector3.Vector3 {
	return &c.Vector3
}

func (c *Cube) RenderOld(cam *vector3.Vector3, pixel *vector3.Vector3) float64 {

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
	
	return math.Inf(1)
}

func (c *Cube) Render(cam *vector3.Vector3, ray_dir *vector3.Position) float64 {
	c.GenerateVertices()

	t_min := math.Inf(1)
	for _, side := range c.Sides {
		t := c.RenderHalfSide(cam, *ray_dir, side[0], side[1], side[2])
		if t < t_min {
			t_min = t
		}
	}

	return t_min
	

}

func (cube *Cube) RenderHalfSide(cam *vector3.Vector3, ray_dir vector3.Position, a, b, c int) float64 {
	v0v1 := cube.Vertices[b].Substract(&cube.Vertices[a].Vector3)
	v0v2 := cube.Vertices[c].Substract(&cube.Vertices[a].Vector3)

	n := v0v1.CrossProduct(&v0v2.Position)
	nDotRayDir := n.DotProduct(&ray_dir)
	if math.Abs(nDotRayDir) <= 0.0001 {
		return math.Inf(1)
	}

	d := n.DotProduct(&cube.Vertices[a].Position) * -1

	t := -1 * (n.DotProduct(&cam.Position) + d) / nDotRayDir
	
	if t < 0 {
		return math.Inf(1)
	}

	ray_dir_vector := vector3.Vector3{
		Position: ray_dir,
	}
	t_vector := vector3.New(t, t, t, 0, 0, 0)
	ray_t_mult := ray_dir_vector.Multiply(&t_vector)

	p := cam.Add(&ray_t_mult)

	v0p := p.Substract(&cube.Vertices[a].Vector3)
	ne := v0v1.CrossProduct(&v0p.Position)
	if n.DotProduct(&ne) < 0 {
		return math.Inf(1)
	}


	v2v1 := cube.Vertices[c].Substract(&cube.Vertices[b].Vector3)
	v1p := p.Substract(&cube.Vertices[b].Vector3)
	ne = v2v1.CrossProduct(&v1p.Position)
	if n.DotProduct(&ne) < 0 {
		return math.Inf(1)
	}

	v2v0 := cube.Vertices[a].Substract(&cube.Vertices[c].Vector3)
	v2p := p.Substract(&cube.Vertices[c].Vector3)
	ne = v2v0.CrossProduct(&v2p.Position)
	if n.DotProduct(&ne) < 0 {
		return math.Inf(1)
	}

	return t
}

func (c *Cube) GenerateVertices() {

	c.Vertices[0] = Vertice {
		Vector3: vector3.Vector3{
			Position: vector3.Position {
				X: c.Position.X,
				Y: c.Position.Y,
				Z: c.Position.Z,
			},
		},
	}
	c.Vertices[1] = Vertice {
		Vector3: vector3.Vector3{
			Position: vector3.Position {
				X: c.Position.X,
				Y: c.Position.Y + c.Size,
				Z: c.Position.Z,
			},
		},
	}
	c.Vertices[2] = Vertice {
		Vector3: vector3.Vector3{
			Position: vector3.Position {
				X: c.Position.X + c.Size,
				Y: c.Position.Y + c.Size,
				Z: c.Position.Z,
			},
		},
	}
	c.Vertices[3] = Vertice {
		Vector3: vector3.Vector3{
			Position: vector3.Position {
				X: c.Position.X + c.Size,
				Y: c.Position.Y,
				Z: c.Position.Z,
			},
		},
	}
	c.Vertices[4] = Vertice {
		Vector3: vector3.Vector3{
			Position: vector3.Position {
				X: c.Position.X,
				Y: c.Position.Y,
				Z: c.Position.Z + c.Size,
			},
		},
	}
	c.Vertices[5] = Vertice {
		Vector3: vector3.Vector3{
			Position: vector3.Position {
				X: c.Position.X,
				Y: c.Position.Y + c.Size,
				Z: c.Position.Z + c.Size,
			},
		},
	}
	c.Vertices[6] = Vertice {
		Vector3: vector3.Vector3{
			Position: vector3.Position {
				X: c.Position.X + c.Size,
				Y: c.Position.Y + c.Size,
				Z: c.Position.Z + c.Size,
			},
		},
	}
	c.Vertices[7] = Vertice {
		Vector3: vector3.Vector3{
			Position: vector3.Position {
				X: c.Position.X + c.Size,
				Y: c.Position.Y,
				Z: c.Position.Z + c.Size,
			},
		},
	}

	center := vector3.Position{
        X: c.Position.X + c.Size/2,
        Y: c.Position.Y + c.Size/2,
        Z: c.Position.Z + c.Size/2,
    }

	for i := range c.Vertices {
		c.Vertices[i].Rotate(c.Rotation, center)
	}

	// c.RotateOld(c.Rotation)

}

func (c *Cube) RotateOld(rotation vector3.Rotation) {
	empty := vector3.Rotation{}
	if rotation == empty {
		return
	}

    // Convert angles to radians
    radX := rotation.X * math.Pi / 180
    radY := rotation.Y * math.Pi / 180 
    radZ := rotation.Z * math.Pi / 180

    // Store cube center
    center := vector3.Position{
        X: c.Position.X + c.Size/2,
        Y: c.Position.Y + c.Size/2,
        Z: c.Position.Z + c.Size/2,
    }

    // Rotation matrices
    for i := range c.Vertices {
        // Translate to origin
        x := c.Vertices[i].Position.X - center.X
        y := c.Vertices[i].Position.Y - center.Y 
        z := c.Vertices[i].Position.Z - center.Z

        // Rotate around X axis
        newY := y*math.Cos(radX) - z*math.Sin(radX)
        newZ := y*math.Sin(radX) + z*math.Cos(radX)
        y = newY
        z = newZ

        // Rotate around Y axis
        newX := x*math.Cos(radY) + z*math.Sin(radY)
        newZ = -x*math.Sin(radY) + z*math.Cos(radY)
        x = newX
        z = newZ

        // Rotate around Z axis
        newX = x*math.Cos(radZ) - y*math.Sin(radZ)
        newY = x*math.Sin(radZ) + y*math.Cos(radZ)
        x = newX
        y = newY

        // Translate back
        c.Vertices[i].Position.X = x + center.X
        c.Vertices[i].Position.Y = y + center.Y
        c.Vertices[i].Position.Z = z + center.Z
    }
}


