package main

import (
	"fmt"
	"github.com/gonum/stat"
)

// MA  EMA
func MA(data []float64, window int) []float64 {
	var result []float64

	for i := 0; i < len(data); i++ {
		if i < window {
			result = append(result, 0.0)
		} else {
			slice := data[i-window : i]
			mean := stat.Mean(slice, nil)
			result = append(result, mean)
		}
	}

	return result
}

func EMA(data []float64, window int) []float64 {
	var result []float64
	var esf float64 = 2 / float64(window+1)
	var ema float64 = data[0]

	for i := 1; i < len(data); i++ {
		ema = esf*data[i] + (1-esf)*ema
		result = append(result, ema)
	}

	return result
}

func main() {
	prices := []float64{10.0, 12.0, 15.0, 13.0, 11.0, 10.0, 9.0, 8.0, 10.0, 12.0}
	ma := MA(prices, 5)
	ema := EMA(prices, 5)

	// 打印结果
	fmt.Println("Prices:", prices)
	fmt.Println("MA:", ma)
	fmt.Println("EMA:", ema)
}
