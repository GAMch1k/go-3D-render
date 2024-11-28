package square

import (
	"gamch1k.org/render/prefabs/screen"
)

type Square struct {
	X int
	Y int
	Width int
	Height int
	Char string
}


func (s *Square) Move(dx, dy int, sc screen.Screen) {
	s.X += dx
	s.Y += dy

	if s.X <= 0 {
		s.X = 0
	}
	if s.X + s.Width >= sc.Width - 1 {
		s.X = sc.Width - s.Width - 1
	}

	if s.Y < 0 {
		s.Y = 0
	}
	if s.Y + s.Height >= sc.Height - 1 {
		s.Y = sc.Height - s.Height - 1
	}
}

func (s *Square) Draw(sc screen.Screen) {
	for i := s.X; i <= s.X + s.Width; i++ {
		for j := s.Y; j <= s.Y + s.Height; j++ {
			sc.Set(i, j, s.Char)
		}
	}
}