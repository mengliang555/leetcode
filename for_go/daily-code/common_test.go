package dailycode

import "testing"

func Test_maxLengthBetweenEqualCharacters(t *testing.T) {
	println(maxLengthBetweenEqualCharacters("aa"))
	println(maxLengthBetweenEqualCharacters("abca"))
	println(maxLengthBetweenEqualCharacters("cbzxy"))
	println(maxLengthBetweenEqualCharacters("mgntdygtxrvxjnwksqhxuxtrv"))
}

func Test_dayOfYear(t *testing.T) {
	println(dayOfYear("2019-01-09"))
	println(dayOfYear("2019-02-10"))
	println(dayOfYear("1900-05-02"))
}

func Test_diagonalSum(t *testing.T) {
	println(diagonalSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
