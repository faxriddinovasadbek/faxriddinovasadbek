package main

import (
	"fmt"
	
)

func singleNumber(nums []int) int {
	count := 0
    num := 0

    for _, i := range nums {
		for in, _ := range nums {
			if i == nums[in+i]{
				fmt.Println(nums[in])
				count = nums[in]
                num++
			}
			if num == 1{
				break
			}
		}
	}
	return count
}
func main() {
	fmt.Println(singleNumber([]int{1,0,1}))
}
