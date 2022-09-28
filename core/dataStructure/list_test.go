package dataStructure

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {

	l := Constructor()
	l.AddAtHead(2)
	l.AddAtHead(1)
	l.AddAtTail(3)
	l.AddAtIndex(1, 7)
	l.DeleteAtIndex(2)
	var s *node
	s = l.head.next
	for i := 0; i < l.size; i++ {
		fmt.Println(s.val)
		s = s.next
	}
	fmt.Println()
}

func Test_bestSeqAtIndex(t *testing.T) {
	type args struct {
		height []int
		weight []int
	}
	height, weight := []int{2868, 5485, 1356, 1306, 6017, 8941, 7535, 4941, 6331, 6181}, []int{5042, 3995, 7985, 1651, 5991, 7036, 9391, 428, 7561, 8594}
	i := bestSeqAtIndex(height, weight)
	fmt.Println(i)
}
