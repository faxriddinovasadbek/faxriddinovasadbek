package main

import (
	"fmt"
	"strconv"
)

func maximum69Number (num int) int {
    str_num := strconv.Itoa(num)
	str := ""
	for i := 0; i < len(str_num); i++ {
		if str_num[i] == 54  {
			str += str_num[:i]
			str += "9"
			str += str_num[i+1:]
			str1, _ := strconv.Atoi(str)
			return str1
		}
	}
	return num	
}

func main(){
	fmt.Println(maximum69Number(9669))
}
