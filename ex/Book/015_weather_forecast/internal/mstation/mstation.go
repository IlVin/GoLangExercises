package mstation

import (
	"math/rand"
	"time"
)

func WeatherForecast() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(70) - 30
}
