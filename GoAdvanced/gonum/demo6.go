package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
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
	returns := make(plotter.XYs, len(slices))
	cumulativeReturns := make(plotter.XYs, len(slices))
	cumReturn := 0.0
	for i, slice := range slices {
		close := slice[0]
		open := slice[2]
		r := (close - open) / open
		cumReturn += r
		returns[i].X = float64(i)
		returns[i].Y = r
		cumulativeReturns[i].X = float64(i)
		cumulativeReturns[i].Y = cumReturn
	}

	// 创建绘图
	p := plot.New()
	p.Title.Text = "收益率曲线"
	p.X.Label.Text = "时间段"
	p.Y.Label.Text = "收益率"
	p.Add(plotter.NewGrid())

	// 绘制收益率曲线和累积收益率曲线
	err := plotutil.AddLinePoints(p,
		"收益率", returns,
		"累积收益率", cumulativeReturns)
	if err != nil {
		panic(err)
	}

	// 保存图像到文件
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "returns.png"); err != nil {
		panic(err)
	}
}
