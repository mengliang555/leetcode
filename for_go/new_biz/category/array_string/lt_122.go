package array_string

//https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/?envType=study-plan-v2&envId=top-interview-150

//给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
//在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
//返回 你能获得的 最大 利润 。

func maxProfit2(prices []int) int {
	maxPro, minN := 0, prices[0]
	for _, v := range prices {
		if v == minN {
			continue
		} else if v > minN {
			maxPro += v - minN
			minN = v
		} else {
			minN = v
		}
	}
	return maxPro
}
