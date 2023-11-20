package main

import (
    "github.com/faiface/pixel/pixelgl"
    "Prueba1/simulation"
)

func main() {
    pixelgl.Run(func() {
        sim := simulation.NewSimulation()
        sim.Init()
        sim.Run()
    })
}
