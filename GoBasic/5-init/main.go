package main

import (
	// 匿名
	_ "goland/GolangStudy/5-init/lib1"
	// 起别名
	//mylib2 "goland/GolangStudy/5-init/lib2"
	// .  不建议推荐使用
	. "goland/GolangStudy/5-init/lib2"
)

func main() {

	//lib1.Lib1Test()

	//lib2.Lib2Test()
	// mylib2.Lib2Test()
	Lib2Test()
}
