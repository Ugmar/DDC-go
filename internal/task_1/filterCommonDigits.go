package task1

import "math"

func CountDigit(a int, arr []int) (len int) {
	for a > 0 {
		arr[a % 10] = 1
		a /= 10
		len++
	}

	return len
}

func CreateNewDigit(a int, len int, arr_1 []int, arr_2 []int) int {
	var digit, x int
	var divider int = int(math.Pow(10, float64(len - 1)))

	for i := 0; i < len; i++ {
		digit = (a / divider) % 10

		if arr_1[digit] * arr_2[digit] == 0 {
			x += digit
			x *= 10
		}

		divider /= 10
	}

	return x / 10
}

func FilterCommonDigits(a, b int) (int, int, error){
	if a <= 0 || b < 0 {
		return 0, 0, ErrNegNums
	}

	arr_digits_a := make([]int, 10)
	len_a := CountDigit(a, arr_digits_a)

	arr_digits_b := make([]int, 10)
	len_b := CountDigit(b, arr_digits_b)

	var new_a, new_b int

	new_a = CreateNewDigit(a, len_a, arr_digits_a, arr_digits_b)
	new_b = CreateNewDigit(b, len_b, arr_digits_a, arr_digits_b)

	if new_a * new_b == 0 {
		return 0, 0, ErrEmptyNum
	}

	return new_a, new_b, nil
}
