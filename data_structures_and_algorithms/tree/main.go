package main

// Node 二叉树节点

type Node struct {
	left  *Node
	Data  interface{}
	Right *Node
}

/*
接口定义包括
新节点的创建 初始化
二叉树的输出 度的高度 叶子节点的统计
二叉树的中后续遍历 广度遍历
*/

// Initer 接口进行初始化

type Initer interface {
	SetData(data interface{})
}

type Operater interface {
	PrintBT()
	Depth() int
	LeafCount() int
}

// Order 接口提供了2中二叉树遍历的方式
// 深度遍历 先序遍历 中序遍历 后序遍历
// 广度遍历
type Order interface {
	PreOrder()      // 先序遍历
	InOrder()       // 中序遍历
	PostOrder()     // 后序遍历
	BreadthTravel() // 广度遍历

}

/*
接口的实现 通过底层函数的实现
*/
