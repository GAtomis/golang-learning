/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-06-27 15:06:23
 * @LastEditTime: 2022-06-28 17:25:00
 * @LastEditors: Gavin
 */
package datatype

import (
	"fmt"
	"math"
	"unsafe"
)

func funA() {

}

func Enter() bool {
	var name string = "tom"
	age := 20
	b := true
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", b)
	//指针类型
	a := 100
	p := &a
	fmt.Printf("%T\n", p)
	//静态数组
	arr := [2]int{1, 2}
	fmt.Printf("%T\n", arr)
	//切片类型
	act := []int{1, 3, 4, 5, 6}
	fmt.Printf("%T\n", act)
	//函数类型
	fmt.Printf("%T\n", funA)
	return b

}
func GetNumber() {
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64
	fmt.Printf("%T %dB  %v~%v\n", i, unsafe.Sizeof(i), math.MaxInt, math.MinInt)
	fmt.Printf("%T %dB  %v~%v\n", i8, unsafe.Sizeof(i8), math.MaxInt8, math.MinInt8)
	fmt.Printf("%T %dB  %v~%v\n", i16, unsafe.Sizeof(i16), math.MaxInt16, math.MinInt16)
	fmt.Printf("%T %dB  %v~%v\n", i32, unsafe.Sizeof(i32), math.MaxInt32, math.MinInt32)
	fmt.Printf("%T %dB  %v~%v\n", i64, unsafe.Sizeof(i64), math.MaxInt64, math.MinInt64)

	fmt.Printf("%T %dB  %v~%v\n", u8, unsafe.Sizeof(u8), 0, math.MaxUint8)

	fmt.Printf("%T %dB  %v~%v\n", u16, unsafe.Sizeof(u16), 0, math.MaxUint16)
	fmt.Printf("%T %dB  %v~%v\n", u32, unsafe.Sizeof(u32), 0, math.MaxUint32)
	fmt.Printf("%T %dB  %v~%v\n", u64, unsafe.Sizeof(u64), 0, uint64(math.MaxUint64))

}
func GetNumberAndAdvance() {

	var i10 int = 10
	var i8 int = 077
	var i16 int = 0xff
	var f32 float32 = 3.12
	fmt.Printf("i10: %d\n", i10)
	fmt.Printf("i8: %b\n", i8)
	fmt.Printf("i16: %x\n", i16)
	fmt.Printf("f32: %v\n", f32)

}
func GetString() {
	// var str1 string = "hello world"
	// var str2 = "word"
	// s3 := "sdasd"
	// s4 := `dasdasdsdasda
	// dsadsads
	// sdasd`
	//字符串拼接
	// msg := fmt.Sprintf("%s,%s", s3, s4)
	//正常拼接
	// msg1 := s3 + s4
	//转译字符
	// s := "hellow"
	// print(s)
	// print(s)
	// print("\n***********\n")
	// print(s + "\n")
	// print(s)
	// print("c:\\demo\\file.pdf")
	n := 3 //索引
	m := 5
	var s string = "hellow goto"
	fmt.Println(s[n])                  //打印索引3的原始字符
	fmt.Println(s[n:])                 //后续n以后的片段
	fmt.Println(s[n:m])                //n-m的片段
	fmt.Println(s[:m])                 //从头到m的片段
	fmt.Printf("s(len): %v\n", len(s)) //获取长度
}
