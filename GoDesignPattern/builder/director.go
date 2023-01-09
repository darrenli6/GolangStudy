package main

// 指挥者
type Director struct {
	builder ComputerBuilder
}

// 建造方法，定义步骤及顺序（类似于Template Method模版模式）
func (d *Director) Build() Computer {
	d.builder.AddCPU()
	d.builder.AddMemory()
	d.builder.AddMainboard()
	d.builder.AddHardDisk()
	d.builder.AddGPU()

	return d.builder.GetComputer()
}
