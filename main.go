package main

import (
	"awesomeProject/practice"
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
)

func main() {
	practice.InitPractice()
	// 路由分组
	// mysql
	// redis
	// cookie
	// session
	// token
	// mysql读写分离，主从同步
	// 分页
	// 配置文件 读取（配置中心？）
	// 其他中间件 跨域等
	// 微服务
	// etcd
	// 负载均衡
	// 代理（反向代理）
	// 消息队列 （ActiveMQ(apache开源)、RabbitMQ、Kafka、Rocket MQ(阿里开源 纯java)、）
	// 搜索引擎（Elastic Search）
	// k8s
	//maxSum()
	//InitProgram()
	//Begin()
	//tt("2020-05-16 19:20:34|user.login|name=Charles&location=Beijing&device=iPhone")
	//FindAbsMin([]int{-3, -2, -1, 0, 3, 4, 5, 6, 7})
	//middleData([]int{-3, -2, -1, 0, 3, 4, 5, 7})
	//for i := 0; i <= 5; i++ {
	//	c := cCount(i, 5)
	//	fmt.Println(i, "  ", c)
	//}
	//FindPathSum(81)

	//{
	//	wait := sync.WaitGroup{}
	//	wait.Add(100)
	//	for i := 0; i < 100; i++ {
	//		go FindPathSum(i, &wait)
	//	}
	//	wait.Wait()
	//}
	//a := 5_3_7
	//fmt.Println(a)
	//var e ldl
	//s, ok := e.(interface{ Ldl() })
	//fmt.Println(s, "   ", ok)
	//result := algorithm.ClimbStairs3(4)
	//fmt.Println(result)
}

func maxSum() {
	var vv int
	_, _ = fmt.Scanf("%d", &vv)
	fmt.Println(vv)
}

type temp struct{}

func (t *temp) Add(elem int) *temp {
	fmt.Println(elem)
	return &temp{}
}

// LogFormat 格式化日志
func LogFormat(s string) {
	info := strings.Split(s, "|")
	res := info[len(info)-1]
	arr := strings.Split(res, "&")
	inf := &Info{}
	for _, v := range arr {
		if strings.HasPrefix(v, "name=") {
			inf.Name = strings.TrimLeft(v, "name=")
		}
		if strings.HasPrefix(v, "location=") {
			inf.Location = strings.TrimLeft(v, "location=")
		}
		if strings.HasPrefix(v, "device=") {
			inf.Device = strings.TrimLeft(v, "device=")
		}
	}

	infoByte, err := json.MarshalIndent(inf, "", "")
	if err != nil {
		return
	}
	fmt.Println(string(infoByte))
}

//定义一个日志解析函数，接收一个字符串类型的参数，能够把字符串按如下格式返回。可以用任何语
//言实现。
//输入：2020-05-16 19:20:34|user.login|name=Charles&location=Beijing&device=iPhone
//输出：
//{
//name:”Charles”, location:”Beijing”, device:”iPhone” }

// Info 用户登陆信息结构体
type Info struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Device   string `json:"device"`
}

// FindAbsMin 查找绝对值最小的值
func FindAbsMin(in []int) {
	left, right := 0, len(in)-1
	if right == 0 {
		return
	}
	var min = math.MaxInt
	// 如果大于0，找到其中的最小值，比较记录值
	for {
		if left == right {
			min = compare(min, in[left])
			break
		}
		if in[left]*in[right] == 0 {
			min = 0
			break
		}
		if in[left]*in[right] > 0 {
			if in[left] > 0 {
				min = compare(min, in[left])
			} else {
				min = compare(min, in[right])
			}
			break
		}
		if in[left]*in[right] < 0 {
			middle := (left + right) / 2

			if in[left]*in[middle] == 0 {
				min = 0
				break
			}
			if in[left]*in[middle] > 0 {
				if in[left] > 0 {
					min = compare(min, in[left])
				} else {
					min = compare(min, in[middle])
				}
				left = middle
			} else {
				if in[middle] > 0 {
					min = compare(min, in[middle])
				} else {
					min = compare(min, in[right])
				}
				right = middle
			}
		}

	}
	fmt.Println(min)
}

func compare(i, j int) int {
	l, r := i, j
	if l*r == 0 {
		return 0
	}
	if l < 0 {
		l = -1 * l
	}
	if r < 0 {
		r = -1 * r
	}
	if l < r {
		return i
	}
	return j
}

func middleData(in []int) {
	quickSort(in, 0, len(in)-1)
	length := len(in)
	if length%2 == 1 {
		fmt.Println(in[length/2])
	} else {
		fmt.Println(float32(in[length/2]+in[length/2-1]) / 2)
	}

}

func quickSort(nums []int, l, r int) {
	if l < r {
		m := partition(nums, l, r)
		quickSort(nums, l, m-1)
		quickSort(nums, m+1, r)
	}
}

func partition(nums []int, l int, r int) int {
	key := nums[r]
	i := l
	j := l
	for j < r {
		if nums[j] < key {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}

//【选做题】有一座高度是 k 级台阶的楼梯，从下往上走，每跨一步只能向上 1 级或者 3 级台阶。请用
//你认为最优的方法求一共有多少种走法。例如给定楼梯台阶数 k 为 3，一共有 2 种走法。要求算法的
//时间复杂度需要小于 O(2n)，可以用任何语言实现。
//输入：楼梯台阶数 k。
//输出：走法总数。

// FindPathSum 查找路径和函数
func FindPathSum(n int, wait *sync.WaitGroup) {
	defer wait.Done()
	// n对3求余，余1、余2、余0情况分别讨论
	// n = 1: result = 1 = 1 + 0
	// n = 2: result = 1 = 1 + 0
	// n = 3: result = 2 = 1 + 1
	// n = 4: result = 3 = 1 + 2
	// n = 5: result = 4 = 1 + 3 = 1 + (n - 2)
	// n = 6: result = 6 = 1 + 6 - 2 + 1 = 1 + (n - 2) + 1
	// n = 7: result = 1 + 7 - 2 + 3 = 9 = 1 + C(1 (n-2)) + C(2 (n-4))
	// n = 8: result = 1 + 8 - 2 + 6 = 13
	// n = k: result = C(0 k) + C(1 (k - 2)) + C(2 (k - 2*2)) + C(3 (k - 2*3)) + C(4 (k - 2*4)) + ...
	if n == 0 {
		fmt.Println(n, " ", "0")
		return
	}
	res := n / 3
	var result int
	for i := 0; i <= res; i++ {
		result += cCount(i, n-2*i)
	}

	if result < 0 {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic info:", err)
			}
			time.Sleep(1 * time.Second)
			fmt.Println("recover success")
		}()
		panic(fmt.Sprintf("FindPathSum函数当传入的参数为 %d 时，计算的结果数字 %d 越界!", n, result))
	}
	fmt.Println(n, "   ", result)
}

func cCount(start, value int) int {
	var result = 1
	var mid = 1
	for i := start; i > 0; i-- {
		result *= value
		value--
		mid *= i
	}
	return result / mid
}
