package main

import "fmt"

type ovoz interface {
	salomlash() string

}

type mushuk struct {
	mushuk_ovozi string
}

type kuchuk struct {

	kuchuk_ovozi string
}

func (m mushuk) salomlash() string {
	return m.mushuk_ovozi
}

func (k kuchuk) salomlash() string {
	return k.kuchuk_ovozi
}

func salom(o ovoz) {
	fmt.Println(o.salomlash())
}

func main() {

	m := mushuk{mushuk_ovozi: "meow"}
	k := kuchuk{kuchuk_ovozi: "vov-vov"}

	salom(m)
	salom(k)

}
