package models

import (
	"Prueba1/utils"
	"fmt"
	"math/rand"
	"sync"
	"time"
	"github.com/faiface/pixel"
)

type Moto struct {
	ParkingTime int
	Id          int
}

func NewMoto() *Moto {
	rand.Seed(time.Now().UnixNano())
	parkingTime := rand.Intn(10) + 15 
	return &Moto{ParkingTime: parkingTime}
}

func (c *Moto) GenerateMotos(n int, ch chan Moto) {
	for i := 1; i <= n; i++ {
		moto := NewMoto()
		moto.Id = i
		ch <- *moto
		randomSleep(1, 2)
	}
	close(ch)
	fmt.Println("Motocicletas Generados Con Exito")
}

func (c *Moto) Timer(pos int, pc *Parking, mut *sync.Mutex, spaces *[20]bool, chEntrance *chan int, sprite *pixel.Sprite, chWin chan utils.ImgMoto, coo pixel.Vec) {
	enterParkingLot(pos, pc, mut, chEntrance, sprite, chWin, coo)
	parkMoto(c, pos, pc, mut, spaces, chEntrance, sprite, chWin, coo)
	leaveParkingLot(c, pos, pc, mut, spaces, chEntrance, sprite, chWin, coo)
}

func randomSleep(minSeconds, maxSeconds int) {
	rand.Seed(time.Now().UnixNano())
	newTime := rand.Intn(maxSeconds-minSeconds+1) + minSeconds
	time.Sleep(time.Second * time.Duration(newTime))
}

func enterParkingLot(pos int, pc *Parking, mut *sync.Mutex, chEntrance *chan int, sprite *pixel.Sprite, chWin chan utils.ImgMoto, coo pixel.Vec) {
	mut.Lock()
	data := utils.NewImgMoto(sprite, pos, true, coo)
	chWin <- *data
	*chEntrance <- 0
	mut.Unlock()
}

func parkMoto(c *Moto, pos int, pc *Parking, mut *sync.Mutex, spaces *[20]bool, chEntrance *chan int, sprite *pixel.Sprite, chWin chan utils.ImgMoto, coo pixel.Vec) {
	mut.Lock()
	pc.nSpaces--
	fmt.Printf("Estacionamiento. %d: %d\n", c.Id, pos)
	fmt.Printf("Disponibles: %d\n", pc.nSpaces)
	mut.Unlock()

	time.Sleep(time.Second * time.Duration(c.ParkingTime))
}

func leaveParkingLot(c *Moto, pos int, pc *Parking, mut *sync.Mutex, spaces *[20]bool, chEntrance *chan int, sprite *pixel.Sprite, chWin chan utils.ImgMoto, coo pixel.Vec) {
	fmt.Printf("Salida: %d\n", c.Id)

	mut.Lock()
	data := utils.NewImgMoto(sprite, pos, false, coo)
	chWin <- *data
	pc.nSpaces++
	spaces[pos] = true
	fmt.Printf("Disponibles: %d\n", pc.nSpaces)
	mut.Unlock()

	mut.Lock()
	*chEntrance <- 1
	mut.Unlock()
}
