package first

import "testing"

func FindFirstPositive(arr []int, x int) int {
	ans := -1

	if len(arr) == 0 || x < 1 {
		return ans
	}

	for i, v := range arr {
		if v == x {
			ans = i
			break
		}
	}
	return ans
}

func TestFindFirstPositive(t *testing.T) {
	tests := []struct {
		arr      []int
		x        int
		expected int
	}{
		{[]int{2, 4, 6, 8}, 4, 1},      // Второй элемент равен 4
		{[]int{-1, -2, 3, 4}, 3, 2},    // Третий элемент равен 3
		{[]int{-1, -2, -3, -4}, 5, -1}, // Числа X нет в последовательности
		{[]int{0, 1, 2, 3}, 0, -1},     // Ноль не является положительным числом
		{[]int{1, 2, 3, 4}, 5, -1},     // Числа X нет в последовательности
		{[]int{-1, -2, -3, 0}, 0, -1},  // Ноль не является положительным числом
		{[]int{1, 3, 5, 7}, 1, 0},      // Первый элемент равен 1
		{[]int{1, 1, 1, 1}, 1, 0},      // Первый элемент равен 1
		{[]int{1, 2, 3, 2, 1}, 2, 1},   // Второй элемент равен 2
		{[]int{1, 2, 3, 4, 5}, 6, -1},  // Числа X нет в последовательности
		{[]int{4, 3, 2, 1}, 5, -1},     // Числа X нет в последовательности
	}

	for _, test := range tests {
		result := FindFirstPositive(test.arr, test.x)
		if result != test.expected {
			t.Errorf("For sequence %v and x = %d, expected %d, but got %d", test.arr, test.x, test.expected, result)
		}
	}
}
