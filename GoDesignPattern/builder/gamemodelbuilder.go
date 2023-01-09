package main

// 游戏型电脑主机
type GameModelBuilder struct {
	Computer
}

// 实现ComputerBuilder的方法
func (gm *GameModelBuilder) AddCPU() {
	gm.CPU = "Intel-i9"
}

func (gm *GameModelBuilder) AddMemory() {
	gm.Memory = "32G"
}

func (gm *GameModelBuilder) AddMainboard() {
	gm.Mainboard = "X460"
}

func (gm *GameModelBuilder) AddHardDisk() {
	gm.HardDisk = "512G SSD，1T HDD"
}

func (gm *GameModelBuilder) AddGPU() {
	gm.GPU = "NVIDIA GTX3090"
}

func (gm *GameModelBuilder) GetComputer() Computer {
	return Computer{
		CPU:       gm.CPU,
		Memory:    gm.Memory,
		Mainboard: gm.Mainboard,
		HardDisk:  gm.HardDisk,
		GPU:       gm.GPU,
	}

}
