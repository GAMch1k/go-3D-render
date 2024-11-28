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
	Render(cam *Vector3, pixel *Vector3) float64
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

func (p *Position) Dot(p2 *Position) float64 {
	return p.X*p2.X + p.Y*p2.Y + p.Z*p2.Z
}

func (p *Position) Distance(p2 *Position) float64 {
	return math.Sqrt(math.Pow(p.X-p2.X, 2) + math.Pow(p.Y-p2.Y, 2) + math.Pow(p.Z-p2.Z, 2))
}


func (r *Rotation) Rotate(nr Rotation) {
	r.X += nr.X
	r.Y += nr.Y
	r.Z += nr.Z

	if r.X > 360 {
		r.X -= 360
	}
	if r.Y > 360 {
		r.Y -= 360
	}
	if r.Z > 360 {
		r.Z -= 360
	}

	if r.X < 0 {
		r.X += 360
	}
	if r.Y < 0 {
		r.Y += 360
	}
	if r.Z < 0 {
		r.Z += 360
	}
}