package controllers

import (
	"Prueba1/models"
	"Prueba1/views"
	"sync"

	"github.com/faiface/pixel/pixelgl"
)

type EntranceController struct {
	model *models.Entrada 
	view  *views.ViewEntrada
	mut    *sync.Mutex
}

func NewEntranceController(win *pixelgl.Window, mut *sync.Mutex) *EntranceController {
	return &EntranceController{
		model: models.NewEntrada(),
		view:  views.NewViewEntrada(win),
		mut:    mut,
	}
}

func (ec *EntranceController) LoadStates() {
	imgs := ec.view.LoadStatesImages()
	ec.view.SetStateImages(imgs)
}


func (ec *EntranceController) PaintEntrance(pos int) {
	ec.view.PaintEntrance(pos)
}
