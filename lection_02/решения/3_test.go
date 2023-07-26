package first

import "testing"

func FindMaxNumber(arr []int) int {
	max := 0

	for i, v := range arr {
		if v >= arr[max] {
			max = i
		}
	}
	return arr[max]
}

func TestFindMaxNumber(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{2, 4, 6, 8}, 8},
		{[]int{-1, -2, -3, -4}, -1},
		{[]int{0, 0, 0, 0}, 0},
		{[]int{1}, 1},
		{[]int{-1}, -1},
		{[]int{-1, 0, 1}, 1},
		{[]int{5, 3, 9, 1, 7}, 9},
	}

	for _, test := range tests {
		result := FindMaxNumber(test.arr)
		if result != test.expected {
			t.Errorf("For sequence %v, expected %d, but got %d", test.arr, test.expected, result)
		}
	}
}
