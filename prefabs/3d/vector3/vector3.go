package vector3

import(
	"math"
)

type Vector3 struct {
	Position
	Rotation
}

type Position struct {
	X float64
	Y float64
	Z float64
}

type Rotation struct {
	X float64
	Y float64
	Z float64
}

type Object interface {
	Vector() *Vector3
	Render(cam *Vector3, ray_dir *Position) float64
}


func (v *Vector3) Vector() *Vector3 {
	return v
}

func New(px, py, pz, rx, ry, rz float64) Vector3 {
	return Vector3{
		Position: Position{
			X: px,
			Y: py,
			Z: pz,
		},
		Rotation: Rotation{
			X: rx,
			Y: ry,
			Z: rz,
		},
	}
}

func (v *Vector3) Add(v2 *Vector3) Vector3 {
	return Vector3{
		Position: Position{
			X: v.Position.X + v2.Position.X,
			Y: v.Position.Y + v2.Position.Y,
			Z: v.Position.Z + v2.Position.Z,
		},
		Rotation: Rotation{
			X: v.Rotation.X + v2.Rotation.X,
			Y: v.Rotation.Y + v2.Rotation.Y,
			Z: v.Rotation.Z + v2.Rotation.Z,
		},
	}
}

func (v *Vector3) Substract(v2 *Vector3) Vector3 {
	return Vector3{
		Position: Position{
			X: v.Position.X - v2.Position.X,
			Y: v.Position.Y - v2.Position.Y,
			Z: v.Position.Z - v2.Position.Z,
		},
		Rotation: Rotation{
			X: v.Rotation.X - v2.Rotation.X,
			Y: v.Rotation.Y - v2.Rotation.Y,
			Z: v.Rotation.Z - v2.Rotation.Z,
		},
	}
}

func (v *Vector3) Multiply(v2 *Vector3) Vector3 {
	return Vector3{
		Position: Position{
			X: v.Position.X * v2.Position.X,
			Y: v.Position.Y * v2.Position.Y,
			Z: v.Position.Z * v2.Position.Z,
		},
		Rotation: Rotation{
			X: v.Rotation.X * v2.Rotation.X,
			Y: v.Rotation.Y * v2.Rotation.Y,
			Z: v.Rotation.Z * v2.Rotation.Z,
		},
	}
}



func (p *Position) Move(np Position) {
	p.X += np.X
	p.Y += np.Y
	p.Z += np.Z
}

func (p *Position) Magnitude() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

func (p *Position) Normalize() Position {
	magnitude := p.Magnitude()
	return Position{
		X: p.X / magnitude,
		Y: p.Y / magnitude,
		Z: p.Z / magnitude,
	}
}

func (p *Position) DotProduct(p2 *Position) float64 {
	return p.X*p2.X + p.Y*p2.Y + p.Z*p2.Z
}

func (p *Position) CrossProduct(p2 *Position) Position {
	return Position{
		X: p.Y*p2.Z - p.Z*p2.Y,
		Y: p.Z*p2.X - p.X*p2.Z,
		Z: p.X*p2.Y - p.Y*p2.X,
	}
}

func (p *Position) Distance(p2 *Position) float64 {
	return math.Sqrt(math.Pow(p.X-p2.X, 2) + math.Pow(p.Y-p2.Y, 2) + math.Pow(p.Z-p2.Z, 2))
}

func (r *Rotation) Add(r2 *Rotation) *Rotation {
	x := r.X + r2.X
	y := r.Y + r2.Y
	z := r.Z + r2.Z

	if x > 360 {
		x -= 360
	}
	if y > 360 {
		y -= 360
	}
	if z > 360 {
		z -= 360
	}

	if x < -360 {
		x += 360
	}
	if y < -360 {
		y += 360
	}
	if z < -360 {
		z += 360
	}

	return &Rotation{
		X: x,
		Y: y,
		Z: z,
	}
}

func (r *Rotation) Rotate(r2 *Rotation) {
	x := r.X + r2.X
	y := r.Y + r2.Y
	z := r.Z + r2.Z

	if x > 360 {
		x -= 360
	}
	if y > 360 {
		y -= 360
	}
	if z > 360 {
		z -= 360
	}

	if x < -360 {
		x += 360
	}
	if y < -360 {
		y += 360
	}
	if z < -360 {
		z += 360
	}

	r.X = x
	r.Y = y
	r.Z = z
}

	
func (v *Vector3) Rotate(r Rotation, center Position) {
	// if r.X == 0 && r.Y == 0 && r.Z == 0 {
	// 	return
	// }

	v.Position.X -= center.X
	v.Position.Y -= center.Y 
	v.Position.Z -= center.Z

    if r.X != 0 {
        // X rotation
        rx := r.X * math.Pi / 180
        cosX := math.Cos(rx)
        sinX := math.Sin(rx)
        y := v.Position.Y*cosX - v.Position.Z*sinX
        z := v.Position.Y*sinX + v.Position.Z*cosX
        v.Position.Y = y
        v.Position.Z = z
    }

    if r.Y != 0 {
        // Y rotation
        ry := r.Y * math.Pi / 180
        cosY := math.Cos(ry)
        sinY := math.Sin(ry)
        x := v.Position.X*cosY + v.Position.Z*sinY
        z := -v.Position.X*sinY + v.Position.Z*cosY
        v.Position.X = x
        v.Position.Z = z
    }

    if r.Z != 0 {
        // Z rotation
        rz := r.Z * math.Pi / 180
        cosZ := math.Cos(rz)
        sinZ := math.Sin(rz)
        x := v.Position.X*cosZ - v.Position.Y*sinZ
        y := v.Position.X*sinZ + v.Position.Y*cosZ
        v.Position.X = x
        v.Position.Y = y

    }
	
	v.Position.X += center.X
	v.Position.Y += center.Y
	v.Position.Z += center.Z
}	
