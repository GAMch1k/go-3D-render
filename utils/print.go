package utils

import (
	"fmt"

	"gamch1k.org/render/prefabs/3d/vector3"
	"gamch1k.org/render/prefabs/screen"
)


func ShowPosition(s *screen.Screen, pos ...*vector3.Position) {
	for _, p := range pos {
		str := fmt.Sprintf("X: %.2f, Y: %.2f, Z: %.2f", p.X, p.Y, p.Z)
		Print(s, str)
			
	}
}

func Print(s *screen.Screen, str string, args ...int) {
	if args != nil {
		for x, ch := range str {
			s.Set(x, args[0], string(ch))
		}
		return 
	}

	for x, ch := range str {
		s.Set(x, s.StatsLast + 1, string(ch))
	}
	s.StatsLast ++
}

func ShowFps(s *screen.Screen, fps int) {
	str := fmt.Sprintf("fps: %d", fps)
	Print(s, str)
}

