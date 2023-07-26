package first

import "testing"

func FindLastOccurrence(arr []int, x int) int {
	ans := -1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == x {
			ans = i
			break
		}
	}
	return ans
}

func TestFindLastOccurrence(t *testing.T) {
	tests := []struct {
		arr      []int
		x        int
		expected int
	}{
		{[]int{2, 4, 6, 8, 4}, 4, 4},
		{[]int{-1, -2, 3, 4, 3}, 3, 4},
		{[]int{-1, -2, -3, -4, 5}, 5, 4},
		{[]int{0, 1, 2, 3}, 0, 0},
		{[]int{1, 2, 3, 4}, 5, -1},
		{[]int{-1, -2, -3, 0}, 0, 3},
		{[]int{1, 3, 5, 7}, 1, 0},
		{[]int{1, 1, 1, 1}, 1, 3},
		{[]int{1, 2, 3, 2, 1}, 2, 3},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]int{4, 3, 2, 1}, 5, -1},
	}

	for _, test := range tests {
		result := FindLastOccurrence(test.arr, test.x)
		if result != test.expected {
			t.Errorf("For sequence %v and x = %d, expected %d, but got %d", test.arr, test.x, test.expected, result)
		}
	}
}
