package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToInt(str string) float32 {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, ",", ".")
	st := strings.Split(str, ";")
	n1, _ := strconv.ParseFloat(st[0], 32)
	n2, _ := strconv.ParseFloat(st[1], 32)

	num := n1 / n2

	return float32(num)
}

func main() {

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Printf("%.4f", romanToInt(text))
}
