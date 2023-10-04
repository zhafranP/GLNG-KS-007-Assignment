package main

import (
	"fmt"
	"math/rand"
	"time"
)

type weather struct {
	wind  int
	water int
}

func main() {
	weather := weather{}

	for true {
		weather.wind = rand.Intn(100)
		weather.water = rand.Intn(100)
		fmt.Printf("%+v\n", weather)
		fmt.Println("status water", waterStatus(weather.water))
		fmt.Println("status wind", waterStatus(weather.wind))

		time.Sleep(15 * time.Second)
	}

}

func windStatus(w int) string {
	// status wind :
	// x <= 6 -> aman
	// 7 <= x <= 15 siaga
	// x > 15 bahaya

	if w <= 6 {
		return "aman"
	}
	if w >= 7 && w <= 15 {
		return "siaga"
	}
	return "bahaya"
}

func waterStatus(w int) string {
	// status water :
	// x <= 5 aman
	// 6 <= x <= 8 siaga
	// x > 8 bahaya
	if w <= 5 {
		return "aman"
	}
	if w >= 6 && w <= 8 {
		return "siaga"
	}
	return "bahaya"
}
