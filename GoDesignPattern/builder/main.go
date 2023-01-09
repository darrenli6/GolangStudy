package main

import "fmt"

func main() {

	// 办公型电脑建造者
	offceModelBuilder := &OffceModelBuilder{}

	// 实例化一个指挥者，并传入具体建造者
	director := Director{offceModelBuilder}
	computer1 := director.Build()

	fmt.Printf("办公型主机配置：%v \n", computer1)

	// 游戏型电脑建造者
	gameModelBuilder := &GameModelBuilder{}

	// 这里直接赋值，也证明指挥者 Director不直接依赖具体建造者
	director.builder = gameModelBuilder
	computer2 := director.Build()

	fmt.Printf("游戏型主机配置：%v \n", computer2)

}
