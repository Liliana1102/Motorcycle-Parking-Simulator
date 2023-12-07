package views

import (
	"Prueba1/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type ViewMoto struct {
	win    *pixelgl.Window
	utils  *utils.ImageUtils
	sprite *pixel.Sprite
}

type SpriteMoto struct {
	img *pixel.Sprite
	Id  int
}

func NewViewMoto(win *pixelgl.Window) *ViewMoto {
	return &ViewMoto{
		win:   win,
		
	}
}

func (mv *ViewMoto) SetSprite() {
	motoSprite := mv.loadMotoSprite()
	mv.sprite = motoSprite
}

func (mv *ViewMoto) PaintMoto(pos pixel.Vec) *pixel.Sprite {
	mv.sprite.Draw(mv.win, pixel.IM.Moved(pos))
	return mv.sprite
}

func (mv *ViewMoto) loadMotoSprite() *pixel.Sprite {
	picMoto, _ := mv.utils.LoadPicture("./assets/bikeman2.png")
	return mv.utils.NewSprite(picMoto, picMoto.Bounds())
}

func NewImgMoto(spr *pixel.Sprite, Id int) *SpriteMoto {
	return &SpriteMoto{
		img: spr,
		Id:  Id,
	}
}
