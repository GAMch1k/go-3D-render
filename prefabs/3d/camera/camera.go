package camera

import (
	"fmt"
	"math"

	"gamch1k.org/render/prefabs/3d/vector3"
	"gamch1k.org/render/prefabs/screen"
	"gamch1k.org/render/utils"
)

type Camera struct {
	vector3.Vector3
	Screen       *screen.Screen
	ViewDistance float64
}

func (cam *Camera) Move(np vector3.Position) {
	cam.Position.Move(np)
	cam.Screen.Move(np)
}

// func (cam *Camera) Rotate(rotation vector3.Rotation) {
// 	// TO DO

// 	// I need to rotate screen around camera origin
// 	cam.Screen.Center.Rotation.Rotate(&rotation)
// 	cam.Screen.Center.Vector().Rotate(cam.Screen.Center.Rotation, cam.Position)
// 	x := (cam.Screen.Center.Vector().Position.X - float64(cam.Screen.Width) / 2) * cam.Screen.PixelWidth
// 	y := (cam.Screen.Center.Vector().Position.Y + float64(cam.Screen.Height) / 2) * cam.Screen.PixelHeight
// 	z := cam.Screen.Center.Vector().Position.Z

// 	cam.Screen.Position.X = x
// 	cam.Screen.Position.Y = y
// 	cam.Screen.Position.Z = z
// }

func (cam *Camera) Render(objects ...vector3.Object) {
	

	minimal := math.Inf(1)
	
	for scy := range cam.Screen.Shape {
		for scx := range cam.Screen.Shape[scy] {
			// if !utils.StringInSlice(cam.Screen.Get(scx, scy), cam.Screen.Shade) {
			// 	continue
			// }

			pixel := vector3.New(
				cam.Screen.Position.X+(float64(scx)*cam.Screen.PixelWidth),
				cam.Screen.Position.Y-(float64(scy)*cam.Screen.PixelHeight),
				cam.Screen.Position.Z,
				0, 0, 0,
			)

			tmin := math.Inf(1) 

			ray := pixel.Substract(&cam.Vector3)
			ray_dir := ray.Position.Normalize()

			for _, object := range objects {
				
				// if object.Vector().Position.Distance(&cam.Position) > cam.ViewDistance {
				// 	continue
				// }

				t := object.Render(&cam.Vector3, &ray_dir)
				if t < tmin && t >= 0 {
					tmin = t
					if t < minimal {
						minimal = t
					}
				}
			}

			shade := int(cam.ViewDistance / tmin * float64(len(cam.Screen.Shade)))

			if shade >= len(cam.Screen.Shade) {
				shade = len(cam.Screen.Shade) - 1
			} else if shade < 0 {
				shade = 0
			}
			
			cam.Screen.Set(scx, scy, cam.Screen.Shade[shade])

		}
	}

	utils.Print(cam.Screen, "Objects positions:")
	for _, object := range objects {
		utils.ShowPosition(cam.Screen, &object.Vector().Position)
	}

	utils.Print(cam.Screen, fmt.Sprintf("t minimal: %0.2f", minimal))
}

