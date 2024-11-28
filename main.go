package main

import (
	"time"

	"golang.org/x/term"

	"gamch1k.org/render/prefabs/3d/camera"
	"gamch1k.org/render/prefabs/3d/sphere"
	"gamch1k.org/render/prefabs/3d/vector3"
	"gamch1k.org/render/prefabs/screen"
	"gamch1k.org/render/utils"
)

func errHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	screen_w, screen_h, err := term.GetSize(0)
	errHandler(err)

	var screen screen.Screen
	screen.Init(screen_w, screen_h, " ", 0.1, 0.2)

	var camera camera.Camera = camera.Camera{
		Vector3: vector3.New(0, 0, -5, 0, 0, 0),
		Screen:  &screen,
		ViewDistance: 40,
	}

	sphere1 := sphere.Sphere{
		Vector3: vector3.New(0, 0, 20, 0, 0, 0),
		Radius:  3,
	}

	sphere2 := sphere.Sphere{
		Vector3: vector3.New(10, 8, 8, 0, 0, 0),
		Radius:  8,
	}

	keyreader := utils.KeyReader()
	_frames := 0
	_frames_counter := 0
	_frames_counter_timestamp := time.Now().Unix()

	for {
		utils.ShowFps(&screen, _frames)

		camera.Render(sphere1, sphere2)



		utils.ShowPosition(&screen, &camera.Position, &screen.Position)
		
		
		// Showing rendered image
		screen.Draw()
		screen.Clear()

		// Frames counter stuff
		_frames_counter++

		if time.Now().Unix() - _frames_counter_timestamp >= 1 {
			_frames = _frames_counter
			_frames_counter = 0
			_frames_counter_timestamp = time.Now().Unix()
		}

		// Camera movement
		select {
		case key := <-*keyreader:
			switch key {
			case "w":
				camera.Move(vector3.Position{
					X: 0,
					Y: 0,
					Z: 0.1,
				})
			case "s":
				camera.Move(vector3.Position{
					X: 0,
					Y: 0,
					Z: -0.1,
				})
			case "a":
				camera.Move(vector3.Position{
					X: -0.1,
					Y: 0,
					Z: 0,
				})
			case "d":
				camera.Move(vector3.Position{
					X: 0.1,
					Y: 0,
					Z: 0,
				})
			case "q":
				camera.Move(vector3.Position{
					X: 0,
					Y: -0.1,
					Z: 0,
				})
			case "e":
				camera.Move(vector3.Position{
					X: 0,
					Y: 0.1,
					Z: 0,
				})
				
			default:
				continue
			}
		default:
		}
		
		// FPS
		time.Sleep(time.Second / 100)
	}
	
}
