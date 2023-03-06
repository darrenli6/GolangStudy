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

	// 取出第2行第3列的元素
	element := mat.At(1, 2)

	// 将元素放入一维切片
	result := []float64{element}

	// 打印结果
	fmt.Println("Data:", data)
	fmt.Println("Element:", element)
	fmt.Println("Result:", result)
}
