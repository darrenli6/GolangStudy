package main

// 办公型电脑主机
type OffceModelBuilder struct {
	Computer
}

// 实现ComputerBuilder的方法
func (om *OffceModelBuilder) AddCPU() {
	om.CPU = "Intel-i3"
}

func (om *OffceModelBuilder) AddMemory() {
	om.Memory = "8G"
}

func (om *OffceModelBuilder) AddMainboard() {
	om.Mainboard = "B460"
}

func (om *OffceModelBuilder) AddHardDisk() {
	om.HardDisk = "256G HDD"
}

func (om *OffceModelBuilder) AddGPU() {
	om.GPU = "Intel HD630"
}

func (om *OffceModelBuilder) GetComputer() Computer {
	return Computer{
		CPU:       om.CPU,
		Memory:    om.Memory,
		Mainboard: om.Mainboard,
		HardDisk:  om.HardDisk,
		GPU:       om.GPU,
	}
}
