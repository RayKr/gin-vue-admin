package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 数据类型转换
	// string -> int
	var str string = "123"
	i, _ := strconv.Atoi(str)
	fmt.Printf("%T\n", i) // %T可以打印数据类型

	// string -> i64
	i64, _ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("%T\n", i64)

	// int -> string
	i1 := 10
	s1 := strconv.Itoa(i1)
	fmt.Printf("%T\n", s1)

	// int64 -> string
	var i2 int64 = 12343
	s2 := strconv.FormatInt(i2, 10)
	fmt.Printf("%T\n", s2)

	// string -> float32/float64
	s3 := "3.1415926"
	f32, _ := strconv.ParseFloat(s3, 32)
	fmt.Printf("%T %v\n", f32, f32)
	f64, _ := strconv.ParseFloat(s3, 64)
	fmt.Printf("%T %v\n", f64, f64)

	// int64 -> int
	i644 := 450
	i4 := int(i644)
	fmt.Printf("%T\n", i4)

	// int -> int64
	i5 := 32
	i645 := int64(i5)
	fmt.Printf("%T\n", i645)

	// 数组
	// [元素长度]元素类型{元素1， 元素2， ...}
	a := [3]int{0, 1, 2}
	b := [...]int{1, 3, 342, 34545, 4, 456, 454, 45, 325, 2345, 0, 2354}
	c := new([10]int)
	c[5] = 3
	fmt.Println(a, b, c)

	// 数组循环
	zoom := [...]string{"狗子", "猫", "大象"}
	for i := 0; i < len(zoom); i++ {
		fmt.Println(zoom[i] + "跑")
	}
	for i, v := range zoom {
		fmt.Println(i, v)
	}
	fmt.Println(len(zoom), cap(zoom)) // len长度，cap容器大小

	// 二维数组
	arr2 := [3][3]int{
		{0, 1, 2},
		{1, 2, 3},
		{2, 3, 4},
	}
	fmt.Println(arr2)

	// 切片
	as := [3]int{0, 1, 2}
	cl := as[2:] // 2到末尾，切片用:, 前闭后开原则，前面能取到，后面取不到
	fmt.Println(cl)
	cl[0] = 5 // 尝试给切片重新复制，可以看到打印后的原数组的对应位置也变了，说明切片是数组的一部分，是在原数组的基础上做改动
	fmt.Println(cl, as)

	// 数组不能append，但是切片可以append
	// 切片append之后再修改值，就不会改变原数组的值了

	//声明切片的几种方法
	var aa []int          // 不给初始化值，就是个切片
	aaa := make([]int, 5) // make方式创建切片
	fmt.Println(aa, aaa)  // [] [0 0 0 0 0]
	// 两种方法创建切片的区别：aa是不初始化各个元素的，aaa是初始化的

	// map

}
