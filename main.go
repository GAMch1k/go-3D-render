package main

import (
	"time"

	"golang.org/x/term"

	"gamch1k.org/render/prefabs/3d/camera"
	"gamch1k.org/render/prefabs/3d/sphere"
	"gamch1k.org/render/prefabs/3d/cube"
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
		ViewDistance: 15,
	}

	sphere1 := sphere.New(
		vector3.New(0, 0, 20, 0, 0, 0),
		3,
	)

	sphere2 := sphere.New(
		vector3.New(10, 8, 20, 0, 0, 0),
		8,
	)

	cube1 := cube.New(
		vector3.New(-20, -5, 8, 0, 0, 0),
		15,
	)

	keyreader := utils.KeyReader()
	_frames := 0
	_frames_counter := 0
	_frames_counter_timestamp := time.Now().Unix()

	for {

		camera.Render(sphere1, sphere2, cube1)
		// camera.Render(cube1)

		utils.Print(&screen, "Cube Vertices:")
		for _, vertice := range cube1.Vertices {
			utils.ShowPosition(&screen, &vertice.Position)
		}
		utils.Print(&screen, "Cube Rotation:")
		utils.ShowRotation(&screen, &cube1.Rotation)

		utils.Print(&screen, "Screen Rotation:")
		utils.ShowRotation(&screen, &screen.Center.Rotation)

		utils.Print(&screen, "Screen Position:")
		utils.ShowPosition(&screen, &screen.Center.Position)
		utils.ShowPosition(&screen, &screen.Position)


		cube1.Rotation.Rotate(&vector3.Rotation{
			X: 0.3,
			Y: 0.2,
			Z: 0.1,
		})



		// utils.ShowPosition(&screen, &camera.Position, &screen.Position)
		
		utils.ShowFps(&screen, _frames)
		
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

			// case "i":
			// 	camera.Rotate(vector3.Rotation{
			// 		X: 0.1,
			// 		Y: 0,
			// 		Z: 0,
			// 	})
			// case "k":
			// 	camera.Rotate(vector3.Rotation{
			// 		X: -0.1,
			// 		Y: 0,
			// 		Z: 0,
			// 	})
			// case "j":
			// 	camera.Rotate(vector3.Rotation{
			// 		X: 0,
			// 		Y: -0.1,
			// 		Z: 0,
			// 	})
			// case "l":
			// 	camera.Rotate(vector3.Rotation{
			// 		X: 0,
			// 		Y: 0.1,
			// 		Z: 0,
			// 	})
			
				
			default:
				continue
			}
		default:
		}
		
		// FPS
		// time.Sleep(time.Second / 100)
	}
	
}
