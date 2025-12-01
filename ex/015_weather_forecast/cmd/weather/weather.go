package main

import (
	"fmt"
	"weather_forecast/internal/mstation"
)

func main() {
	for i := 0; i < 1000; i++ {
		fmt.Println(mstation.WeatherForecast())
	}
}
