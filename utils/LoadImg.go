package utils

import (
	"image"
	"os"

	"github.com/faiface/pixel"
)


type Utils struct{}


type ImgMoto struct {
	sprite     *pixel.Sprite
	ID         int
	entering   bool
	position   pixel.Vec
}

//nueva instancia de ImgMoto.
func NewImgMoto(sprite *pixel.Sprite, ID int, state bool, position pixel.Vec) *ImgMoto {
	return &ImgMoto{
		sprite:   sprite,
		ID:       ID,
		entering: state,
		position: position,
	}
}

func (ic *ImgMoto) GetSprite() *pixel.Sprite {
	return ic.sprite
}

func (ic *ImgMoto) GetPosition() pixel.Vec {
	return ic.position
}

func (ic *ImgMoto) GetID() int {
	return ic.ID
}

//devuelve si la motocicleta está ingresando o no.
func (ic *ImgMoto) IsEntering() bool {
	return ic.entering
}

//devuelve una copia de la estructura ImgMoto.
func (ic *ImgMoto) GetData() *ImgMoto {
	return ic
}

func (u *Utils) LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func (u *Utils) NewSprite(picture pixel.Picture, form pixel.Rect) *pixel.Sprite {
	return pixel.NewSprite(picture, form)
}
