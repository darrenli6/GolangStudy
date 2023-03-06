package main

import (
	"fmt"
	"github.com/gonum/stat"
	"math"
)

func Returns(prices []float64) []float64 {
	var returns []float64
	for i := 1; i < len(prices); i++ {
		ret := (prices[i] - prices[i-1]) / prices[i-1]
		returns = append(returns, ret)
	}
	return returns
}

func CumulativeReturns(returns []float64) []float64 {
	var cumRet []float64
	cumRet = append(cumRet, 1.0)
	for i := 1; i < len(returns); i++ {
		cumRet = append(cumRet, (1+returns[i])*cumRet[i-1])
	}
	return cumRet
}

func MaxDrawdown(cumRet []float64) float64 {
	var maxRet float64 = 1.0
	var maxDrawdown float64 = 0.0
	for _, ret := range cumRet {
		if ret > maxRet {
			maxRet = ret
		} else {
			drawdown := (maxRet - ret) / maxRet
			if drawdown > maxDrawdown {
				maxDrawdown = drawdown
			}
		}
	}
	return maxDrawdown
}

func AnnualizedReturn(cumRet []float64, days float64) float64 {
	ret := (cumRet[len(cumRet)-1] - 1) / (days / 365)
	return ret
}

func SharpeRatio(returns []float64, riskFreeRate float64) float64 {
	var excessRet []float64
	for _, ret := range returns {
		excessRet = append(excessRet, ret-riskFreeRate)
	}
	mean, std := stat.MeanStdDev(excessRet, nil)
	sharpe := mean / std * math.Sqrt(252)
	return sharpe
}

func main() {
	prices := []float64{100.0, 110.0, 120.0, 130.0, 125.0, 135.0, 140.0, 150.0, 160.0, 170.0}
	returns := Returns(prices)
	cumRet := CumulativeReturns(returns)
	maxDrawdown := MaxDrawdown(cumRet)
	annualizedReturn := AnnualizedReturn(cumRet, float64(len(prices)))
	sharpeRatio := SharpeRatio(returns, 0.0)

	// 打印结果
	fmt.Println("Prices:", prices)
	fmt.Println("Returns:", returns)
	fmt.Println("Cumulative Returns:", cumRet)
	fmt.Println("Max Drawdown:", maxDrawdown)
	fmt.Println("Annualized Return:", annualizedReturn)
	fmt.Println("Sharpe Ratio:", sharpeRatio)
}
