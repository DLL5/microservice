package dataStructure

// 链表
type node struct {
	val        int
	prev, next *node
}

type MyLinkedList struct {
	head, tail *node
	size       int
}

func Constructor() MyLinkedList {
	var head, tail node
	tail.prev = &head
	head.next = &tail
	return MyLinkedList{head: &head, tail: &tail, size: 0}
}

func (l *MyLinkedList) Get(index int) int {
	if !l.checkIndex(index) {
		return -1
	}

	var result = l.head
	for i := 0; i <= index; i++ {
		result = result.next
	}
	return result.val
}

func (l *MyLinkedList) AddAtHead(val int) {
	var addHe node
	addHead := &addHe
	addHead.val = val
	addHead.next = l.head.next
	addHead.prev = l.head
	l.head.next.prev = addHead
	l.head.next = addHead
	l.size++
	return
}

func (l *MyLinkedList) AddAtTail(val int) {
	var addTa node
	addTail := &addTa
	addTail.val = val

	addTail.next = l.tail
	addTail.prev = l.tail.prev
	l.tail.prev.next = addTail
	l.tail.prev = addTail

	l.size++
	return
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		l.AddAtHead(val)
		return
	}
	if index == l.size {
		l.AddAtTail(val)
		return
	}
	if index > l.size {
		return
	}
	var lis, newNo node
	list, newNode := &lis, &newNo
	newNode.val = val
	list = l.head
	for i := 0; i <= index; i++ {
		list = list.next
	}

	newNode.next = list
	newNode.prev = list.prev
	list.prev.next = newNode
	list.prev = newNode

	l.size++
	return
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if !l.checkIndex(index) {
		return
	}
	var list = l.head

	for i := 0; i <= index; i++ {
		list = list.next
	}

	list.prev.next = list.next
	list.next.prev = list.prev
	list.prev = nil
	list.next = nil
	l.size--
}

func (l *MyLinkedList) checkIndex(index int) bool {
	if index < 0 {
		return false
	}
	if index >= l.size {
		return false
	}
	return true
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

//
//type node struct {
//	val        int
//	next, prev *node
//}
//
//type MyLinkedList struct {
//	head, tail *node
//	size       int
//}
//
//func Constructor() MyLinkedList {
//	head := &node{}
//	tail := &node{}
//	head.next = tail
//	tail.prev = head
//	return MyLinkedList{head, tail, 0}
//}
//
//func (l *MyLinkedList) Get(index int) int {
//	if index < 0 || index >= l.size {
//		return -1
//	}
//	var curr *node
//	if index+1 < l.size-index {
//		curr = l.head
//		for i := 0; i <= index; i++ {
//			curr = curr.next
//		}
//	} else {
//		curr = l.tail
//		for i := 0; i < l.size-index; i++ {
//			curr = curr.prev
//		}
//	}
//	return curr.val
//}
//
//func (l *MyLinkedList) AddAtHead(val int) {
//	l.AddAtIndex(0, val)
//}
//
//func (l *MyLinkedList) AddAtTail(val int) {
//	l.AddAtIndex(l.size, val)
//}
//
//func (l *MyLinkedList) AddAtIndex(index, val int) {
//	if index > l.size {
//		return
//	}
//	index = max(0, index)
//	var pred, succ *node
//	if index < l.size-index {
//		pred = l.head
//		for i := 0; i < index; i++ {
//			pred = pred.next
//		}
//		succ = pred.next
//	} else {
//		succ = l.tail
//		for i := 0; i < l.size-index; i++ {
//			succ = succ.prev
//		}
//		pred = succ.prev
//	}
//	l.size++
//	toAdd := &node{val, succ, pred}
//	pred.next = toAdd
//	succ.prev = toAdd
//}
//
//func (l *MyLinkedList) DeleteAtIndex(index int) {
//	if index < 0 || index >= l.size {
//		return
//	}
//	var pred, succ *node
//	if index < l.size-index {
//		pred = l.head
//		for i := 0; i < index; i++ {
//			pred = pred.next
//		}
//		succ = pred.next.next
//	} else {
//		succ = l.tail
//		for i := 0; i < l.size-index-1; i++ {
//			succ = succ.prev
//		}
//		pred = succ.prev.prev
//	}
//	l.size--
//	pred.next = succ
//	succ.prev = pred
//}
//
//func max(a, b int) int {
//	if b > a {
//		return b
//	}
//	return a
//}
//
//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/design-linked-list/solution/she-ji-lian-biao-by-leetcode-solution-abix/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func bestSeqAtIndex(height []int, weight []int) int {
	end := len(height)
	if end == 0 || end == 1 {
		return end
	}
	var user []users
	user = make([]users, end, end)
	for i := 0; i < end; i++ {
		user[i].height = height[i]
		user[i].weight = weight[i]
	}

	for i := 0; i < end; i++ {
		for j := 0; j < i; j++ {
			if user[j].height > user[i].height {
				user[j], user[i] = user[i], user[j]
			}
		}
	}
	count := 0
	// 求体重的最长上升子序列
	
	return count
}

type users struct {
	height int
	weight int
}
