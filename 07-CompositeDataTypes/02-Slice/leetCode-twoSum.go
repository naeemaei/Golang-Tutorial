package main

import (
	"fmt"
	"reflect"
)

func main() {
	nums := []int{2, 7, 5, 11, 15, 5}
	target := 10
	result := twoSum(nums, target)
	fmt.Println(result)
	println(reflect.DeepEqual(result, []int{2, 5}))
}

func twoSum(nums []int, target int) []int {

	for i, item := range nums {
		complement := target - item
		for j := i + 1; j < len(nums); j++ {
			if i != j && nums[j] == complement {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}

}
