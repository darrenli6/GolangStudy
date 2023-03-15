package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"math"
)

func main() {
	// 假设切片数组如下所示
	slices := [][]float64{
		{10.0, 1.0, 9.0, 100.0},
		{11.0, 2.0, 10.0, 200.0},
		{12.0, 3.0, 11.0, 300.0},
	}

	// 将收盘价作为价格
	prices := make([][]float64, len(slices))
	for i, slice := range slices {
		close := slice[0]
		prices[i] = []float64{close}
	}

	// 计算每个时间段的收益率
	returns := make([]float64, len(prices)-1)
	for i := 1; i < len(prices); i++ {
		returns[i-1] = (prices[i][0] - prices[i-1][0]) / prices[i-1][0]
	}

	// 计算复合收益率、年化收益率和最大回撤
	totalReturn := 1.0
	for _, r := range returns {
		totalReturn *= 1 + r
	}
	annualReturn := math.Pow(totalReturn, 1.0/float64(len(returns))) - 1.0

	maxDrawdown := MaxDrawdown(prices)

	// 计算夏普比
	riskFreeRate := 0.0
	sharpeRatio := SharpeRatio(returns, riskFreeRate)

	// 输出结果
	fmt.Printf("年化收益率为%.2f%%\n", annualReturn*100)
	fmt.Printf("最大回撤为%.2f%%\n", maxDrawdown*100)
	fmt.Printf("夏普比为%.2f\n", sharpeRatio)
}

// 计算最大回撤
func MaxDrawdown(prices [][]float64) float64 {
	n := len(prices)
	drawdowns := make([]float64, n)
	peak := prices[0][0]
	for i := 1; i < n; i++ {
		if prices[i][0] > peak {
			peak = prices[i][0]
		}
		drawdowns[i] = (peak - prices[i][0]) / peak
	}

	maxValue := math.Inf(-1)
	for _, v := range drawdowns {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

// 计算夏普比
func SharpeRatio(returns []float64, riskFreeRate float64) float64 {
	meanReturn := stat.Mean(returns, nil)
	stdDev := stat.StdDev(returns, nil)
	return (meanReturn - riskFreeRate) / stdDev * math.Sqrt(252)
}
