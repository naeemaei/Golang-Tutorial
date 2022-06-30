package main

func main() {
	i := 0
	//lst := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	lst1 := []int{12, 15, 14, 13, 16, 17, 18, 19, 20, 21}
	for {
		println("Hello, world!")
		break
	}

	for i < 10 {
		println("Hello, world! ", i)
		i++
	}

	for j := 0; j < 10; j++ {
		println("Hello, world! ", j)
	}

	for index, item := range lst1 {
		println("Hello, world! ", index, item)
	}
}

func twoSum(nums []int, target int) []int {

	for i, item := range nums {
		if item >= target {
			continue
		}
		compliment := target - item
		for j, itemJ := range nums {
			if itemJ == compliment && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}

}

// 11 * 11 = 121
// 10 * 11 / 2 = 55

// 110 * 110 = 12100
// 110 * 109 / 2 = 5995

func twoSum1(nums []int, target int) []int {

	// 7,11,15,45,22,10,2,8,10,12,14
	// 3,00,00,00,00,00,8,2,00,00,00

	for i, item := range nums {
		if item >= target {
			continue
		}
		compliment := target - item
		for j := i; j < len(nums); j++ {
			if nums[j] == compliment && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}

}

// [7,11,15,45,22,10,2,8,10,12,14]
// 10
// [3,2,41,7,6,9,8,110,14,20,25,23]
// 13
// [3,3]
// 6

func twoSum2(nums []int, target int) []int {

	// 7,11,15,45,22,10,2,8,10,12,14 // 121
	// 3,00,00,00,00,00,8,2,00,00,00 // 55
	keys := map[int]int{}
	for i, item := range nums {

		complement := target - item
		if _, exist := keys[complement]; exist {
			return []int{keys[complement], i}
		}

		keys[item] = i

	}
	return []int{0, 0}

}
