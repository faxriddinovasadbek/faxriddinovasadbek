package main

import (
	"fmt"
	"time"
)

type car struct {
	name  string
	year  int
	price float32
	color string
}

func (a *car) change_price(num float32) {
	a.price = num
}

func (a *car) how_year() {
	t := time.Now().Year()
	
	fmt.Println("Mashina", int(t) - a.year, "yillik")
}

func main() {

	c := car{
		name: "Tahoe",
		year:  2023,
		price: 70000,
		color: "black",
	}


	fmt.Println(c.price)
	c.change_price(50000)
	fmt.Println(c.price)

	c.how_year()

	
}
