package dailycode

func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs3(n int) int {
	if n <= 3 {
		return n
	}
	val := [2]int{1, 2}
	for i := 3; i <= n; i++ {
		temp := val[1]
		val[1] = val[1] + val[0]
		val[0] = temp
	}
	return val[1]
}

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	one, two := 1, 2
	for i := 3; i <= n; i++ {
		two, one = one+two, two
	}
	return two
}

func climbStairs2(n int) int {
	if n <= 3 {
		return n
	}
	val := make([]int, n+1)
	val[1], val[2] = 1, 2
	for i := 3; i <= n; i++ {
		val[i] = val[i-1] + val[i-2]
	}
	return val[n]
}

func makeEqual(words []string) bool {
	charCount := make(map[rune]int)

	for _, word := range words {
		for _, character := range word {
			charCount[character]++
		}
	}

	for _, count := range charCount {
		if count%len(words) != 0 {
			return false
		}
	}

	return true
}

//
//
//1 2 3 5 8
//      3 4  5
//one 1 3 8  21
//two 2 5 13 33
