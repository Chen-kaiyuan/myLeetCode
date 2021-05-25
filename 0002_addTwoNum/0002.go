package main

import "fmt"

type Object interface{}

type Node struct {
	Val  int   // 定义数据域
	Next *Node // 定义地址域
}

type List struct {
	headNode *Node // 头节点
}

func main() {
	testlist1 := List{}
	testlist1.Add(2)
	testlist1.Append(4)
	testlist1.Append(3)
	testlist1.Append(4)
	testlist2 := List{}
	testlist2.Add(2)
	testlist2.Append(4)
	testlist2.Append(3)
	result := addTwoNumbers(testlist1.headNode, testlist2.headNode)
	fmt.Println(result.Next)

}

func addTwoNumbers(l1, l2 *Node) (head *Node) {
	var tail *Node
	carry := 0
	sum := 0
	for l1 != nil || l2 != nil {

		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		sum, carry = carry%10, carry/10
		if head == nil {
			head = &Node{Val: sum}
			tail = head
		} else {
			tail.Next = &Node{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &Node{Val: carry}
	}
	return
}

// IsEmpty 判断链表是否为空
func (l *List) IsEmpty() bool {
	if l.headNode == nil {
		return true
	} else {
		return false
	}
}

// Length 获取链表长度
func (l *List) Length() int {
	// 获取链表头节点
	cur := l.headNode
	// 计数器
	count := 0

	// 如果头节点不为空，则移动位置并且计数器+1，直到头节点为空
	for cur != nil {
		count++
		cur = cur.Next
	}

	return count
}

// Add 从链表头部添加一个数据
func (l *List) Add(data int) *Node {
	node := &Node{Val: data}
	node.Next = l.headNode
	l.headNode = node
	return node
}

// Append 从链表尾部添加一个数据
func (l *List) Append(data int) {
	node := &Node{Val: data}
	if l.IsEmpty() { // 判断链表是否为空，为空的话，直接写入数据到头节点
		l.headNode = node
	} else { // 如果链表不为空，则把数据写入尾结点
		cur := l.headNode
		for cur.Next != nil { // 判断尾结点是否为空,如果地址域为空，则是尾节点
			cur = cur.Next // 如果地址域不为空，先把尾节点数据位移到头节点
		}
		cur.Next = node // 直接把数据加到尾结点
	}
}

// ShowList 遍历链表所以节点并打印
func (l *List) ShowList() {
	if !l.IsEmpty() { // 判断链表是否为空
		cur := l.headNode
		i := 1
		for { // 遍历链表
			fmt.Printf("第%v个节点数据是:%v\n", i, cur.Val)
			i++
			if cur.Next != nil { //  判断尾结点是否为空
				cur = cur.Next // ,如果地址域不为空,则把尾节点数据位移到头节点
			} else { //  如果地址域为空，则跳出循环
				break
			}
		}
	}
}

// Get 取某一个位置的数字
func (l *List) Get(index int) Object {
	cur := l.headNode
	if index == 0 {
		return cur.Val
	} else if index > l.Length() {
		return nil
	} else {
		i := 0
		for i != (index-1) && cur.Next != nil {
			i++
			cur = cur.Next
		}
		return cur.Val
	}
	return nil
}
