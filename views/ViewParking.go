package views

import (
	"Prueba1/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type ViewParking struct {
	win            *pixelgl.Window
	utils          *utils.ImageUtils
	spaces         [20]pixel.Vec
	questionSpaces [20]pixel.Vec
}

func NewViewParking(win *pixelgl.Window) *ViewParking {
	return &ViewParking{
		win: win,
		spaces: [20]pixel.Vec{
			pixel.V(835, 300), pixel.V(335, 285), pixel.V(835, 465), pixel.V(835, 545), pixel.V(835, 625),
			pixel.V(685, 625), pixel.V(685, 545), pixel.V(685, 465), pixel.V(285, 185), pixel.V(685, 300),
			pixel.V(175, 625), pixel.V(355, 545), pixel.V(375, 465), pixel.V(175, 385), pixel.V(175, 300),
			pixel.V(325, 300), pixel.V(325, 385), pixel.V(325, 465), pixel.V(325, 545), pixel.V(325, 625),
		},
		questionSpaces: [20]pixel.Vec{
			pixel.V(975, 300), pixel.V(975, 385), pixel.V(975, 465), pixel.V(975, 545), pixel.V(975, 625),
			pixel.V(560, 625), pixel.V(560, 545), pixel.V(560, 465), pixel.V(560, 385), pixel.V(560, 300),
			pixel.V(50, 625), pixel.V(50, 545), pixel.V(50, 465), pixel.V(50, 385), pixel.V(50, 300),
			pixel.V(450, 300), pixel.V(450, 385), pixel.V(450, 465), pixel.V(450, 545), pixel.V(450, 625),
		},
	}
}

func (vp *ViewParking) PaintParking() {
	picParking, err := vp.utils.LoadPicture("./assets/estacionamiento.png")
	if err != nil {
		panic(err)
	}

	parking := vp.utils.NewSprite(picParking, picParking.Bounds())

	matrix := pixel.IM
	matrix = pixel.IM.Moved(pixel.V(512, 400))
	parking.Draw(vp.win, matrix)
}



func (vp *ViewParking) GetCoordinates(n int) pixel.Vec {
	return vp.spaces[n]
}
