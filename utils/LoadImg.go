package utils

import (
	"image"
	"os"

	"github.com/faiface/pixel"
)

//estructura que contiene funciones de utilidad.
type Utils struct{}

// ImgMoto es la estructura que representa una img.
type ImgMoto struct {
	sprite     *pixel.Sprite
	ID         int
	entering   bool
	position   pixel.Vec
}

// NewImgMoto crea una nueva instancia de ImgMoto.
func NewImgMoto(sprite *pixel.Sprite, ID int, state bool, position pixel.Vec) *ImgMoto {
	return &ImgMoto{
		sprite:   sprite,
		ID:       ID,
		entering: state,
		position: position,
	}
}

// GetSprite devuelve el sprite asociado a ImgMoto.
func (ic *ImgMoto) GetSprite() *pixel.Sprite {
	return ic.sprite
}

// GetPosition devuelve la posición del motocicleta.
func (ic *ImgMoto) GetPosition() pixel.Vec {
	return ic.position
}

// GetID devuelve el ID del motocicleta.
func (ic *ImgMoto) GetID() int {
	return ic.ID
}

// IsEntering devuelve si la motocicleta está ingresando o no.
func (ic *ImgMoto) IsEntering() bool {
	return ic.entering
}

// GetData devuelve una copia de la estructura ImgMoto.
func (ic *ImgMoto) GetData() *ImgMoto {
	return ic
}

// LoadPicture carga una img desde el archivo especificado.
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

// NewSprite crea un nuevo sprite con la imagen y forma especificadas.
func (u *Utils) NewSprite(picture pixel.Picture, form pixel.Rect) *pixel.Sprite {
	return pixel.NewSprite(picture, form)
}
