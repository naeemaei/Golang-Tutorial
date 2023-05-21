package twoSum

func twoSumWithMakeMap(nums []int, target int) []int {
	checked := make(map[int]int)
	for index, num := range nums {
		if index2, ok := checked[target-num]; ok {
			return []int{index2, index}
		}
		checked[num] = index
	}
	return []int{}
}

func twoSumWithNewMap(nums []int, target int) []int {
	checked := map[int]int{}
	for index, num := range nums {
		if index2, ok := checked[target-num]; ok {
			return []int{index2, index}
		}
		checked[num] = index
	}
	return []int{}
}

func twoSumWithMakeMapAndExtraAssign(nums []int, target int) []int {
	checked := make(map[int]int)
	for index, num := range nums {
		complete := target - num
		if _, ok := checked[complete]; ok {
			return []int{checked[complete], index}
		}
		checked[num] = index
	}
	return []int{}
}

func twoSumWithSlice1(nums []int, target int) []int {
	// 7,11,15,45,22,10,2,8,10,12,14 // 121
	// 3,00,00,00,00,00,8,2,00,00,00 // 55

	for i, item := range nums {

		compliment := target - item
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == compliment {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}

}

func twoSumWithSlice2(nums []int, target int) []int {
	// 7,11,15,45,22,10,2,8,10,12,14 // 121
	// 3,00,00,00,00,00,8,2,00,00,00 // 55

	for i, item := range nums {

		compliment := target - item
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == compliment && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}

}
