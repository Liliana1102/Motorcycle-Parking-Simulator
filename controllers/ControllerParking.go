package controllers

import (
	"Prueba1/models"
	"Prueba1/utils"
	"Prueba1/views"
	"sync"

	"github.com/faiface/pixel/pixelgl"
)

type ControllerParking struct {
	model *models.Parking
	view  *views.ViewParking
	mut    *sync.Mutex
}

func NewControllerParking(win *pixelgl.Window, mut *sync.Mutex) *ControllerParking {
	model := models.NewParking()
	view := views.NewViewParking(win)
	return &ControllerParking{
		model: model,
		view:  view,
		mut:    mut,
	}
}

func (cp *ControllerParking) PaintParking() {
	cp.view.PaintParking()
}

func (cp *ControllerParking) Park(chMoto *chan models.Moto, entranceController *EntranceController, ControllerMoto *ControllerMoto, chEntrance *chan int, chWin chan utils.ImgMoto) {
	go cp.ChangingState(chEntrance, entranceController)

	for moto := range *chMoto {
		pos := cp.model.FindSpaces()
		if pos != -1 {
			coo := cp.view.GetCoordinates(pos)
			ControllerMoto.view.SetSprite()
			sprite := ControllerMoto.view.PaintMoto(coo)

			state := entranceController.model.GetState()
			if state == "Parado" || state == "Entrando" {
				go moto.Timer(pos, cp.model, cp.mut, cp.model.GetAllSpaces(), chEntrance, sprite, chWin, coo)
			} else {
				*chEntrance <- 0
				go moto.Timer(pos, cp.model, cp.mut, cp.model.GetAllSpaces(), chEntrance, sprite, chWin, coo)
			}
		}
	}
}

func (cp *ControllerParking) ChangingState(chEntrance *chan int, entranceController *EntranceController) {
	for change := range *chEntrance {
		entranceController.model.SetState(change)
	}
}
