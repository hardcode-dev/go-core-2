package main

import "fmt"

type point struct {
	lat float64
	lon float64
}

type object struct {
	name  string
	point // встраивание, обращение по имени типа
}

func main() {
	o := object{}
	o.name = "Object"
	o.point.lat = 10
	o.point.lon = 20
	o.PrintPoint()
}

func (p point) PrintPoint() {
	fmt.Printf("lat:lon - %f, %f\n", p.lat, p.lon)
}
