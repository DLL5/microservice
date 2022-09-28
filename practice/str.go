package practice

import (
	"fmt"
	"strconv"
)

// BackSubString 字符串，回文子串
func BackSubString(in string) string {
	inByte := []byte(in)
	// 低效率写法,只返回找到的第一个最长回文子串
	var (
		left, right int
		length      int
	)
	for i := 0; i < len(inByte)-1; i++ {
		for j := i + length; j < len(inByte); j += 1 {
			if isBackBytes(inByte[i:j]) {
				left, right = i, j
				length = j - i
			}
		}
	}
	fmt.Println(string(inByte[left:right]))
	return string(inByte[left:right])
}

// 输入数据不能为空
func isBackBytes(by []byte) bool {
	length := len(by)
	if length == 0 {
		return false
	}
	for i := 0; i < length/2; i++ {
		if by[i] != by[length-i-1] {
			return false
		}
	}
	return true
}

//type MapValue interface {
//	int | int32 | int64
//}

type sumMapType map[string]any

var SumMapType sumMapType

func InitSumMapType() {
	//sum := make(sumMapType, 10)
	//SumMapType = sum
	SumMapType = make(map[string]any, 10)
	return
}

func SumMap(name string, m any) {
	if _, ok := SumMapType[name]; !ok {
		SumMapType[name] = m
	}
	return
}

func SumMapPrint() {
	for i, v := range SumMapType {
		fmt.Println("Index: ", i, "Value: ", v)
	}
}

func SumMapBegin() {
	InitSumMapType()
	count := 0
	for i := 0; i < 10000; i++ {
		SumMap(strconv.Itoa(i), i)
		count++
	}
	SumMapPrint()
	if k, ok := SumMapType["10"]; ok {
		fmt.Println("Random value: ", k)

	}
	fmt.Println("count: ", count)
}
