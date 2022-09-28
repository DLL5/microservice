package algorithm

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target,
// 写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
// 二分查找（简单）
func binarySearch(nums []int, target int) int {
	var left, right, middle int
	right = len(nums) - 1
	if right < 0 {
		return -1
	}

	for {
		if left > right {
			return -1
		}
		middle = (left + right) / 2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] < target {
			left = middle + 1
		}
		if nums[middle] > target {
			right = middle - 1
		}
	}
}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

//var guessValue int
type guessValue struct {
	lock  sync.Mutex
	value int
}

var gv = guessValue{}

func guessNumber(n int) int {
	// guess 函数，返回猜测的值与生成的随机数值的大小
	guess := func(num int) int {
		if num > gv.value {
			return -1
		}
		if num < gv.value {
			return 1
		}
		return 0
	}

	var (
		val         int
		left, right = 1, n
	)

	defer gv.lock.Unlock()
	seed := int64(time.Now().Nanosecond()) // 随机数种子
	rand.Seed(seed)
	gv.lock.Lock()
	gv.value = int(rand.Int31n(int32(n))) + 1 // 生成随机数并修改gv的值
	check := gv.value
	fmt.Println("随机值为： ", gv.value)

	var count int
	// 二分查找，知道找到结果
	for {
		count++
		val = (left + right) / 2
		if guess(val) == 0 {
			break
		}
		if guess(val) == -1 {
			right = val - 1
		}
		if guess(val) == 1 {
			left = val + 1
		}
	}
	// 检查是否支持并发运行
	if check != gv.value {
		log.Fatal("the target is not equal value, please check the program! ")
	}
	fmt.Println("最终猜测的值为： ", val)
	fmt.Println("查找范围最大值为： ", n, "，查找次数count为： ", count)
	return val
}

func bestSeqAtIndex(height []int, weight []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	person := make([][2]int, n)
	for i := range person {
		person[i] = [2]int{height[i], weight[i]}
	}
	sort.Slice(person, func(i, j int) bool {
		if person[i][0] == person[j][0] {
			return person[i][1] < person[j][1]
		}
		return person[i][0] > person[j][0]
	})
	var f []int
	for i := 0; i < n; i++ {
		if len(f) == 0 || person[f[len(f)-1]][0] > person[i][0] && person[f[len(f)-1]][1] > person[i][1] {
			f = append(f, i)
		} else {
			l, r := 0, len(f)-1
			for l < r {
				mid := (l + r) >> 1
				if person[f[mid]][0] > person[i][0] && person[f[mid]][1] > person[i][1] {
					l = mid + 1
				} else {
					r = mid
				}
			}
			f[r] = i
		}
	}
	return len(f)
}

//作者：himuralee
//链接：https://leetcode.cn/problems/circus-tower-lcci/solution/go-zui-chang-di-zeng-zi-xu-lie-de-er-wei-ban-jian-/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
