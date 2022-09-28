package algorithm

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "testOne", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9}, want: 4},
		{name: "testTwo", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 5}, want: 3},
		{name: "testThree", args: args{nums: []int{-1, 0, 3, 5, 9, 12, 18, 20, 300}, target: 20}, want: 7},
		{name: "testFour", args: args{nums: []int{-10, -8, -6, -3, -1, 0, 3, 5, 9, 12, 18, 20, 300}, target: 3}, want: 6},
		{name: "testFive", args: args{nums: []int{-10, -8, -6, -3, -1, 0, 3, 5, 9, 12, 18, 20, 300}, target: -3}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_guessNumber(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "One", args: args{100}, want: gv.value}, // 注意：guessValue是运行中的变量，此时赋值会使want恒为0，导致测试结果异常
		{name: "Two", args: args{333}, want: gv.value},
		{name: "Three", args: args{555}, want: gv.value},
		{name: "Four", args: args{200}, want: gv.value},
		{name: "Five", args: args{999}, want: gv.value}}
	concurrentGuessNumber()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := guessNumber(tt.args.num); got != gv.value {
				t.Errorf("guess() = %v, want %v", got, gv.value)
			}
		})
	}
}

// 并发测试封装
func concurrentGuessNumber() {
	var count = 10000
	wait := sync.WaitGroup{}
	wait.Add(count)
	for i := 0; i < count; i++ {
		go conGuessNumber(&wait)
	}
	wait.Wait()
	fmt.Println("concurrentGuessNumber execute over...")
}

func conGuessNumber(wait *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	guessNumber(1000)
	wait.Done()
}
