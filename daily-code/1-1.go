package dailycode

// 2022. 将一维数组转变成二维数组
// 给你一个下标从 0 开始的一维整数数组 original 和两个整数 m 和  n 。你需要使用 original 中 所有 元素创建一个 m 行 n 列的二维数组。

// original 中下标从 0 到 n - 1 （都 包含 ）的元素构成二维数组的第一行，下标从 n 到 2 * n - 1 （都 包含 ）的元素构成二维数组的第二行，依此类推。

// 请你根据上述过程返回一个 m x n 的二维数组。如果无法构成这样的二维数组，请你返回一个空的二维数组。

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return make([][]int, 0)
	}

	ans := make([][]int, m)

	for i := range ans {
		ans[i] = make([]int, n)
	}

	row, column := 0, 0
	for _, v := range original {
		ans[row][column] = v
		column += 1
		if column == n {
			row += 1
			column = 0
		}
	}
	return ans
}

// left -> right f(x)
// right -> lest b(x)
// 1) x == 1 f(x)==b(x)=1
// 2) b(x) +f(x) = n+1
// 3) f(x) = 2 * b(n/2)
func lastRemaining(n int) int {
	if n == 1 {
		return 1
	}
	return 2 * (n/2 + 1 - lastRemaining(n/2))
}
