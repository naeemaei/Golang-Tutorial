package twoSum

import (
	"testing"
)

var testCases = [][]int{
	{5, 1, 6, 3, 10, 8, 12, 7, 9, 4},    // 5
	{11, 20, 1, 5, 9, 4, 3, 16, 18, 7},  // 13
	{16, 2, 9, 1, 7, 4, 18, 10, 3, 13},  // 9
	{81, 7, 13, 1, 4, 12, 12, 6, 9, 5},  // 15
	{4, 7, 2, 9, 1, 5, 8, 3, 6, 10},     // 18
	{3, 4, 15, 6, 2, 9, 11, 7, 10, 8},   // 7
	{19, 18, 15, 3, 4, 6, 2, 7, 10, -1}, // 11
	{10, 2, 3, 7, 4, 6, 1, 8, 5, 9},     // 3
	{13, 5, 1, 10, 6, 7, 19, 2, 14, 18}, // 14
	{7, 4, 51, 8, 10, 1, 3, 6, 2, 9},    // 6
}

var testTargets = []int{5, 13, 9, 15, 18, 7, 11, 3, 14, 6}

var testExpected = [][]int{
	{1, 9},
	{4, 5},
	{1, 4},
	{7, 8},
	{6, 9},
	{0, 1},
	{4, 7},
	{1, 6},
	{0, 2},
	{1, 8},
}

func BenchmarkTwoSumWithMakeMap(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i, item := range testCases {
			twoSumWithMakeMap(item, testTargets[i])
			// if result[0] != testExpected[i][0] || result[1] != testExpected[i][1] {
			// 	fmt.Printf("item %d", i)
			// }

		}
	}
}

func BenchmarkTwoSumWithNewMap(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i, item := range testCases {
			twoSumWithNewMap(item, testTargets[i])
			// if result[0] != testExpected[i][0] || result[1] != testExpected[i][1] {
			// 	fmt.Printf("item %d", i)
			// }

		}
	}
}

func BenchmarkTwoSumWithMakeMapAndExtraAssign(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i, item := range testCases {
			twoSumWithMakeMapAndExtraAssign(item, testTargets[i])
			// if result[0] != testExpected[i][0] || result[1] != testExpected[i][1] {
			// 	fmt.Printf("item %d", i)
			// }

		}
	}
}

func BenchmarkTwoSumWithSlice1(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i, item := range testCases {
			twoSumWithSlice1(item, testTargets[i])
			// if result[0] != testExpected[i][0] || result[1] != testExpected[i][1] {
			// 	fmt.Printf("item %d", i)
			// }

		}
	}
}

func BenchmarkTwoSumWithSlice2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i, item := range testCases {
			twoSumWithSlice2(item, testTargets[i])
			// if result[0] != testExpected[i][0] || result[1] != testExpected[i][1] {
			// 	fmt.Printf("item %d", i)
			// }

		}
	}
}
