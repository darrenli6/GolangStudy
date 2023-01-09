package main

type ComputerBuilder interface {
	// 添加中央处理器
	AddCPU()
	// 添加内存
	AddMemory()
	// 添加主板
	AddMainboard()
	// 添加硬盘
	AddHardDisk()
	// 添加显卡
	AddGPU()

	// 获取电脑主机实例
	GetComputer() Computer
}

// 电脑结构， 主要描述电脑有哪些重要组件，没有方法
type Computer struct {
	CPU       string
	Memory    string
	Mainboard string
	HardDisk  string
	GPU       string
}
