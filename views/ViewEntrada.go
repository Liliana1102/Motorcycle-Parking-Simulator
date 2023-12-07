package views

import (
	"Prueba1/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type ViewEntrada struct {
	win    *pixelgl.Window
	utils  *utils.ImageUtils
	states [3]pixel.Sprite
}

func NewViewEntrada(win *pixelgl.Window) *ViewEntrada {
	return &ViewEntrada{
		win:   win,
		
	}
}

func (ev *ViewEntrada) LoadStatesImages() [3]pixel.Sprite {
	openEntrance, openingEntrance, closeEntrance := ev.loadEntranceSprites()
	return [3]pixel.Sprite{openEntrance, openingEntrance, closeEntrance}
}

func (ev *ViewEntrada) loadEntranceSprites() (pixel.Sprite, pixel.Sprite, pixel.Sprite) {
	picEntranceOpen, _ := ev.utils.LoadPicture("./assets/bikeman2.png")
	picEntranceOpening, _ := ev.utils.LoadPicture("./assets/Entrada.png")
	picEntranceClose, _ := ev.utils.LoadPicture("./assets/Alarm.png")

	openEntrance := ev.utils.NewSprite(picEntranceOpen, picEntranceOpen.Bounds())
	openingEntrance := ev.utils.NewSprite(picEntranceOpening, picEntranceOpening.Bounds())
	closeEntrance := ev.utils.NewSprite(picEntranceClose, picEntranceClose.Bounds())

	return *openEntrance, *openingEntrance, *closeEntrance
}

func (ev *ViewEntrada) SetStateImages(imgs [3]pixel.Sprite) {
	ev.states = imgs
}

func (ev *ViewEntrada) PaintEntrance(img int) {
	entrancePos := pixel.V(920, 200)
	ev.states[img].Draw(ev.win, pixel.IM.Moved(entrancePos))
}
