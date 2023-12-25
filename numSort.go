package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode"
)

func IsSorted(n []int, m []int) bool {
	for i := 0; i < len(n); i++ {
		if n[i] != m[i] {
			return false
		}
	}
	return true
}

func main() {
	str1, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// fmt.Println(str1)

	nums := ""
	arr := []string{}

	for _,  val := range str1 {
		if unicode.IsDigit(val){
			nums += string(val)
		}else{
			if nums != ""{
				arr = append(arr, nums)
				nums = ""
			}
		}
	}

	num_arr := []int{}	

	for _, val := range arr {
		n, _ :=  strconv.Atoi(string(val))
		num_arr = append(num_arr, n)
	}

	arrCopy := make([]int, len(num_arr))
	copy(arrCopy, num_arr)
	
	sort.Slice(arrCopy, func(i, j int) bool {
		return arrCopy[i] < arrCopy[j]	
	})

	if len(arrCopy) > 0{
		fmt.Println(IsSorted(arrCopy, num_arr))
	}else{
		fmt.Println(false)
	}
}
