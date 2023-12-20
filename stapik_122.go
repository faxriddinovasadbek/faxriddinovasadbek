package main

import (
	"fmt"
	"strings"
)

type Battery struct {
	char string
}

func (a *Battery) string() string {
	nol := strings.Count(a.char, "0")
	bir := strings.Count(a.char, "1")

	return "[" + strings.Repeat(" ", nol) + strings.Repeat("X", bir) + "]"
}

func main(){

	var s string
	fmt.Scan(&s)
	batteryForTest := &Battery{char: s}
}
