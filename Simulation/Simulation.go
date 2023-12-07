package simulation

import (
	_ "image/png"
	"Prueba1/controllers"
	"Prueba1/models"
	"Prueba1/utils"
	"sync"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Simulation struct {
	win             *pixelgl.Window
	motoChannel      chan models.Moto
	entranceChannel chan int
	winChannel      chan utils.ImgMoto
	mut              *sync.Mutex
	parkingCtrl     *controllers.ControllerParking
	entranceCtrl    *controllers.EntranceController
	motoCtrl         *controllers.ControllerMoto
	motoSprites      []utils.ImgMoto
}

//ventana
func NewSimulation() *Simulation {
	cfg := pixelgl.WindowConfig{
		Title:  "Estacionamiento",
		Bounds: pixel.R(0, 0, 1024, 700),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	
	motoChannel := make(chan models.Moto, 100)
	entranceChannel := make(chan int)
	winChannel := make(chan utils.ImgMoto)
	mut := &sync.Mutex{}

	return &Simulation{
		win:             win,
		motoChannel:      motoChannel,
		entranceChannel: entranceChannel,
		winChannel:      winChannel,
		mut:              mut,
		parkingCtrl:     controllers.NewControllerParking(win, mut),
		entranceCtrl:    controllers.NewEntranceController(win, mut),
		motoCtrl:         controllers.NewControllerMoto(win, mut),
	}
}

func (s *Simulation) Init() {
	s.motoCtrl.LoadSprite()
	s.entranceCtrl.LoadStates()
}

func (s *Simulation) Run() {
	go s.parkingCtrl.Park(&s.motoChannel, s.entranceCtrl, s.motoCtrl, &s.entranceChannel, s.winChannel)
	go s.motoCtrl.GenerateMotos(100, &s.motoChannel)

	for !s.win.Closed() {
		s.win.Clear(colornames.Black)

		s.parkingCtrl.PaintParking()
		

		s.handleWinChannel()

		for _, value := range s.motoSprites {
			s.drawMotoSprite(value)
		}

		s.win.Update()
	}
}

func (s *Simulation) handleWinChannel() {
	select {
	case val := <-s.winChannel:
		if val.IsEntering() {
			s.motoSprites = append(s.motoSprites, val)
		} else {
			s.removeMotoSprite(val)
		}
	}
}

func (s *Simulation) removeMotoSprite(val utils.ImgMoto) {
	var arrAux []utils.ImgMoto
	for _, value := range s.motoSprites {
		if value.GetID() != val.GetID() {
			arrAux = append(arrAux, value)
		}
	}
	s.motoSprites = s.motoSprites[:0]
	s.motoSprites = append(s.motoSprites, arrAux...)
}
func (s *Simulation) drawMotoSprite(value utils.ImgMoto) {
	sprite := value.GetSprite()
	pos := value.GetPosition()
	sprite.Draw(s.win, pixel.IM.Moved(pos))
}
 