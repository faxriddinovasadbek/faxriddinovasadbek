package main

import "fmt"

func twoSum(nums []int, target int) []int{

	var index []int
	check := false

	for in, i := range nums{
		for j := 1; j < len(nums); j++ {
			if (i + nums[j]) == target{
				index = append(index, in)
				index = append(index, in+1)
				check = true
				break
			}
		}
		if check {
			break
		}
	}

	return index
}

func main(){

	fmt.Println(twoSum([]int{3,2,3}, 6))
}
