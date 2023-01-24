package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func ShowNode(p *Node) {
	if p != nil {
		fmt.Println(*p)
		p = p.next
	}
	fmt.Println()
}

// 是否存在元素obj
func Find(p *Node, obj int) int {

	head := p

	if head == nil {
		return -1
	}

	for i := 0; head != nil; i++ {
		if head.data == obj {
			return i
		}
		head = head.next
	}

	return -1
}

func Update(p *Node, val, index int) bool {
	head := p
	if head == nil {
		return false
	}

	for i := 1; i <= index; i++ {

		head = head.next
		if head == nil {
			return false
		}
	}
	head.data = val
	return true
}

func linkList(length int) *Node {
	head := Node{data: 0}

	var tail *Node

	tail = &head

	for i := 1; i <= length; i++ {
		// 声明node
		node := Node{data: i}
		// 头部指向下一个地址
		tail.next = &node
		// 尾部是下一个node
		tail = &node
	}

	return &head
}
func GetPreNode(p *Node, index int) *Node {
	head := p
	if head == nil || index < 0 {
		return head
	}

	// 找到前驱节点
	// 倒序
	for ; index-1 > 0; index-- {
		if head.next != nil {
			head = head.next
		} else {
			return nil
		}
	}

	return head
}

func Insert(p *Node, val, index int) bool {
	head := p
	if head == nil || index < 0 {
		return false
	}
	preNode := GetPreNode(head, index)
	if preNode == nil {
		return false
	}

	node := Node{data: val, next: preNode.next}
	preNode.next = &node

	return true

}

const LENGTH = 3

func main() {

	head := linkList(LENGTH)

	ShowNode(head)

	fmt.Println("Find0", Find(head, 0))
	fmt.Println("Find0", Find(head, 2))
	fmt.Println("Find0", Find(head, 4))

	if Insert(head, 10, 4) {
		fmt.Println("insert successfully")
		ShowNode(head)
	} else {
		fmt.Println("插入失败")
	}

}
