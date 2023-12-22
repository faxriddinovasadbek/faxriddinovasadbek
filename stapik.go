package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	soat, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	t, _ := strconv.Atoi(soat[11:13])

	d , _:= strconv.Atoi(soat[8:10])
	if t < 13{
		d++
		fmt.Println("2020-05-15 08:00:00")
		}else{
			fmt.Printf("%s%v%s\n",soat[:8],d, soat[10:]) 
	}

}
