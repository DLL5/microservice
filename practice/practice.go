package practice

import (
	"fmt"
	"time"
)

func InitPractice() {
	//chanPractice()
	//mainTest()
	//mutexDeadLockTest()
	//waitGroupPanicTest()
	//in := "sas oooHelloolleHworld"
	//BackSubString(in)
	SumMapBegin()
}

func chanPractice() {
	c := make(chan int, 10)
	end := make(chan bool)
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(100 * time.Millisecond)
	}

	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-end:
				close(c)
				return
			}
		}

	}()
	time.Sleep(time.Second)
	fmt.Println("hello world")
	end <- true
	close(end)
}

//var c = make(chan int)
//var a int
//
//func f() {
//	a = 1
//	<-c
//}
//func mainTest() {
//	go f()
//	c <- 0
//	print(a)
//}

//type MyMutex struct {
//	count int
//	sync.Mutex
//}

//func mutexDeadLockTest() {
//	var mu MyMutex
//	mu.Lock()
//	var mu1 = mu
//	mu.count++
//	mu.Unlock()
//	mu1.Lock()
//	mu1.count++
//	mu1.Unlock()
//	fmt.Println(mu.count, mu1.count)
//	// 加锁后复制变量，会将锁的状态也复制，所以 mu1 其实是已经加锁状态，再加锁会死锁。
//}

//func waitGroupPanicTest() {
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go func() {
//		time.Sleep(time.Millisecond)
//		wg.Done()
//		wg.Add(1)
//	}()
//	//time.Sleep(time.Second)
//	wg.Wait()
//	// sync.WaitGroup在调用我Wait()方法后，再调用Add()方法会panic。
//}
