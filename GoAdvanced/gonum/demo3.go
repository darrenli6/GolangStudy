package main

import (
	"fmt"
	"github.com/gonum/matrix/mat64"
)

func main() {
	// 创建一个3x3的二维切片
	data := [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	// 将二维切片转换为矩阵
	mat := mat64.NewDense(len(data), len(data[0]), nil)
	for i := 0; i < len(data); i++ {
		mat.SetRow(i, data[i])
	}

	// 获取第2列的数据
	col := mat.ColView(1)

	// 将列数据转换为一维切片
	result := make([]float64, col.Len())
	for i := 0; i < col.Len(); i++ {
		result[i] = col.At(i, 0)
	}

	// 打印结果
	fmt.Println("Data:", data)
	fmt.Println("Result:", result)
}
