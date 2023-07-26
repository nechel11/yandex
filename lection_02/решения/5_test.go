package first

import "testing"

func FindMinEven(arr []int) int {
	ans := -1

	for _, v := range arr {
		if ans == -1 && v%2 == 0 {
			ans = v
		} else if v%2 == 0 && v < ans {
			ans = v
		}
	}
	return ans
}

func TestFindMinEven(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{2, 4, 6, 8}, 2},      // Все числа четные
		{[]int{1, 3, 5, 7}, -1},     // Все числа нечетные
		{[]int{-1, -2, -3, -4}, -4}, // Все числа отрицательные
		{[]int{-1, -2, -3, 0}, -2},  // Имеется четное число 0
		{[]int{1, 2, 3, 4}, 2},      // Четное число 2
		{[]int{5, 3, 9, 1, 7}, -1},  // Все числа нечетные
	}

	for _, test := range tests {
		result := FindMinEven(test.arr)
		if result != test.expected {
			t.Errorf("For sequence %v, expected %d, but got %d", test.arr, test.expected, result)
		}
	}
}
