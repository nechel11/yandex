package first

import "testing"

func FindMaxAndSecondMax(arr []int) (int, int) {
	max_1, max_2 := maxMin(arr[0], arr[1])

	for i := 2; i <= len(arr)-1; i++ {
		if arr[i] > max_2 {
			max_1 = max_2
			max_2 = arr[i]
		} else if arr[i] > max_1 {
			max_1 = arr[i]
		}
	}
	return max_2, max_1
}

func maxMin(a, b int) (int, int) {
	if a >= b {
		return b, a
	}
	return a, b
}

func TestFindMaxAndSecondMax(t *testing.T) {
	tests := []struct {
		arr               []int
		expectedMax       int
		expectedSecondMax int
	}{
		{[]int{2, 4, 6, 8}, 8, 6},
		{[]int{-1, -2, -3, -4}, -1, -2},
		{[]int{0, 0, 0, 0}, 0, 0},
		{[]int{1, 1, 1, 1}, 1, 1},
		{[]int{5, 3, 9, 1, 7}, 9, 7},
	}

	for _, test := range tests {
		resultMax, resultSecondMax := FindMaxAndSecondMax(test.arr)
		if resultMax != test.expectedMax || resultSecondMax != test.expectedSecondMax {
			t.Errorf("For sequence %v, expected max=%d and second max=%d, but got max=%d and second max=%d", test.arr, test.expectedMax, test.expectedSecondMax, resultMax, resultSecondMax)
		}
	}
}
