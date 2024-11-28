package point

import (
	"gamch1k.org/render/prefabs/screen"
)

type Point struct {
	X int
	Y int
	Char string
}


func (p *Point) Move(dx, dy int, sc screen.Screen) {
	p.X += dx
	p.Y += dy

	if p.X <= 0 {
		p.X = 0
	} else if p.X >= sc.Width {
		p.X = sc.Width - 1
	}

	if p.Y <= 0 {
		p.Y = 0
	} else if p.Y >= sc.Height {
		p.Y = sc.Height - 1
	}
}

func (p *Point) Draw(sc screen.Screen) {
	sc.Set(p.X, p.Y, p.Char)
}
