package main

import (
	"fmt"
	"math"
)

func main() {
	// 假设收盘价、时间和开盘价分别为 10.0, 1.0, 9.0
	slice1 := []float64{10.0, 1.0, 9.0}
	// 假设收盘价、时间和开盘价分别为 11.0, 2.0, 10.0
	slice2 := []float64{11.0, 2.0, 10.0}
	// 假设收盘价、时间和开盘价分别为 12.0, 3.0, 11.0
	slice3 := []float64{12.0, 3.0, 11.0}

	slices := [][]float64{slice1, slice2, slice3}

	// 计算每个时间段的收益率
	returns := make([]float64, len(slices))
	for i, slice := range slices {
		close := slice[0]
		open := slice[2]
		returns[i] = (close - open) / open
	}

	// 计算复合收益率和年化收益率
	// 计算复合收益率和年化收益率
	totalReturn := 1.0
	for _, r := range returns {
		totalReturn *= (r + 1.0)
	}
	totalReturn -= 1.0
	annualReturn := math.Pow(totalReturn+1.0, 1.0/float64(len(returns))) - 1.0
	// 输出结果
	fmt.Printf("年化收益率为%.2f%%\n", annualReturn*100)
}
