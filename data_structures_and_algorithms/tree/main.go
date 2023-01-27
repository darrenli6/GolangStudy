package main

import "fmt"

// Node 二叉树节点

type Node struct {
	Left  *Node
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
func (n *Node) SetData(data interface{}) { n.Data = data }
func (n *Node) PrintBT()                 { PrintBT(n) }
func (n *Node) Depth() int               { return Depth(n) }
func (n *Node) LeafCount() int           { return LeafCount(n) }
func (n *Node) PreOrder()                { PreOrder(n) }
func (n *Node) InOrder()                 { InOrder(n) }
func (n *Node) PostOrder()               { PostOrder(n) }
func (n *Node) BreadthTravel()           { BreadthTravel(n) }

// 底层函数的实现

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

// printBT 用于输出一个给定的二叉树  采用递归的算法
// 根节点 -- 左子树 -- 右子树
func PrintBT(n *Node) {
	if n != nil {
		fmt.Println("%v ", n.Data)

		if n.Left != nil || n.Right != nil {
			fmt.Printf("(")
			PrintBT(n.Left)
			if n.Right != nil {
				fmt.Printf(",")
			}
			PrintBT(n.Right)
			fmt.Printf(")")

		}
	}
}

// Depth 用于计算二叉树的深度 采用递归算法
// 如果二叉树为空 则深度为0

func Depth(n *Node) int {
	var depthleft, depthright int
	if n == nil {
		return 0
	}
	depthleft = Depth(n.Left)
	depthright = Depth(n.Right)
	if depthleft > depthright {
		return depthleft + 1
	}
	return depthright + 1

}

// LeefCount 用于统计一个二叉树叶子节点数 采用递归算法
// 如果一个儿茶为空 则叶子节点数为0

func LeafCount(n *Node) int {
	if n == nil {
		return 0
	} else if n.Left == nil && n.Right == nil {
		return 1
	} else {
		return LeafCount(n.Left) + LeafCount(n.Right)
	}
}

// PreOrder 先序遍历 采用递归算法
// 根节点 -- 左子树 -- 右子树

func PreOrder(n *Node) {

	if n != nil {
		fmt.Printf("%v ", n.Data)
		PreOrder(n.Left)
		PreOrder(n.Right)
	}
}
