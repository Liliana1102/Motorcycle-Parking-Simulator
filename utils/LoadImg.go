package utils

import (
	"image"
	"os"

	"github.com/faiface/pixel"
)

type ImageUtils struct{}

type ImgMoto struct {
	sprite   *pixel.Sprite
	ID       int
	entering bool
	position pixel.Vec
}

func NewImgMoto(sprite *pixel.Sprite, ID int, state bool, position pixel.Vec) *ImgMoto {
	return &ImgMoto{
		sprite:   sprite,
		ID:       ID,
		entering: state,
		position: position,
	}
}

func (im *ImgMoto) GetSprite() *pixel.Sprite {
	return im.sprite
}

func (im *ImgMoto) GetPosition() pixel.Vec {
	return im.position
}

func (im *ImgMoto) GetID() int {
	return im.ID
}

func (im *ImgMoto) IsEntering() bool {
	return im.entering
}

func (im *ImgMoto) GetData() *ImgMoto {
	return im
}

func (u *ImageUtils) LoadPicture(path string) (pixel.Picture, error) {
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

func (u *ImageUtils) NewSprite(picture pixel.Picture, form pixel.Rect) *pixel.Sprite {
	return pixel.NewSprite(picture, form)
}
