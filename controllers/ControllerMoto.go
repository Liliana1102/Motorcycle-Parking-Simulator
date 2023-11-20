package controllers

import (
	"Prueba1/models"
	"Prueba1/views"
	"sync"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type ControllerMoto struct {
	model *models.Moto
	view  *views.ViewMoto
	mut    *sync.Mutex
}

func NewControllerMoto(win *pixelgl.Window, mut *sync.Mutex) *ControllerMoto {
	return &ControllerMoto{
		model: models.NewMoto(),
		view:  views.NewViewMoto(win),
		mut:    mut,
	}
}

//se genera las motos util el modelo
func (cm *ControllerMoto) GenerateMotos(n int, chMoto *chan models.Moto) {
	cm.model.GenerateMotos(n, *chMoto)
}

//cargando el sprite de la moto
func (cm *ControllerMoto) LoadSprite() {
	cm.view.SetSprite()
}

func (cm *ControllerMoto) PaintMoto(pos pixel.Vec) {
	cm.view.PaintMoto(pos)
}
