package screen

import (
	"fmt"
	"strings"

	"gamch1k.org/render/prefabs/3d/vector3"
)

type Screen struct {
	vector3.Vector3
	Center vector3.Vector3
	Width int
	Height int
	Shape [][]string
	Background string
	PixelWidth float64
	PixelHeight float64
	Shade []string

	StatsLast int
}

func (s *Screen) Clear() {
	for i := range s.Shape {
		s.Shape[i] = make([]string, s.Width)
		for j := range s.Shape[i] {
			s.Shape[i][j] = s.Background
		}
	}
	s.StatsLast = 0
}

func (s *Screen) Draw() {
	draw := ""

	for row := range s.Shape {
		draw += strings.Join(s.Shape[row], "")
	}

	fmt.Println(draw)
}

func (s *Screen) Init(w, h int, bg string, pixelwidth, pixelheight float64) {
	s.Width = w
	s.Height = h
	s.Background = bg
	s.PixelWidth = pixelwidth
	s.PixelHeight = pixelheight
	s.Center = vector3.New(0, 0, 0, 0, 0, 0)
	s.Vector3 = vector3.New(
		float64(0 - w / 2) * pixelwidth,
		float64(h / 2) * pixelheight,
		0, 0, 0, 0,
	)
	s.Shade = []string{" ", "░", "▒", "▓", "█"}
	// s.Shade = []string{" ", ".", ",", "_", "-", "=", "+", "*", "%", "#", "@"}

	

	s.Shape = make([][]string, h)
	s.Clear()
}

func (s *Screen) Set(x, y int, c string) {
	s.Shape[y][x] = c
}

func (s *Screen) Get(x, y int) string {
	return s.Shape[y][x]
}

