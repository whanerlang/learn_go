package variable

import (
	. "fmt"
)

/*
数组Array

定义数组的格式：var <varName> [n]<type>，n>=0
数组长度也是类型的一部分，因此具有不同长度的数组为不同类型
注意区分指向数组的指针和指针数组
数组在Go中为值类型
数组之间可以使用==或!=进行比较，但不可以使用<或>
可以使用new来创建数组，此方法返回一个指向数组的指针
Go支持多维数组
*/

func ArrayVar() {
	Println("---------------------ArrayVar--------------------------")
	array1()
	array2()
	array3()
	array4()
	array5()
	array6()
	array7()
	array8()
	array9()
	array10()
}

func array1() {
	Println("..............array1..........")
	a := [20]int{15: 15} //长度也是精妙数组的一部分，长度不同，类型也不同
	b := [...]int{1, 2, 3, 4, 5}
	c := [...]int{0: 3, 18: 2}
	d := [5]int{}
	Println("a=", a)
	Println("b=", b)
	Println("c=", c)
	Println("d=", d)
	Println("d==b:", d == b)
}

func array2() {
	Println("..............array2..........")
	var arr = [10]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		n := arr[i]
		Println("arr[", i, "]=", n)
	}
	Println()
	Println("arr[m:n]=", arr[1:6])

}

func array3() {
	Println("..............array3..........")
	a := [...]int{1, 2, 3, 4, 5}
	var pa *[5]int = &a //指向数组指针 pb= &[1 2 3 4 5]
	x, y := 12, 2
	var b = [...]*int{&x, &y} //保存*int的数组 d= [0xc042054408 0xc042054410]

	Println("a=", a)
	Println("b=", b)
	Println("*b[0]=", *b[0])
	Println("*b[1]=", *b[1])
	Println("pb=", pa)
	// Println("a==c",a==c)
	// Println("b==*pb",b==*pb)
	// Println("&b==pb",&b==pb)
	Println()
}

func array4() {
	Println("..............array4..........")
	i := [2][3]int{ //行数可以自动计算  i:=[...][3]int 列不行
		{0: 1, 2, 3},
		{10, 20, 2: 30}}
	Println("i=", i)
}

func array5() {
	Println("..............array5..........")
	//new 创建指向数组的指针
	var p = new([5]int)
	// var p =[5]int{}
	len := len(p)
	Println("len(p)=", len)
	for a, i := p, 0; i < len; i++ {
		a[i] = 20 + i
		Println(a)
	}
	Println(p)
}

func array6() {
	Println("..............array6..........")
	s1 := make([]int, 3, 6)
	Printf("%v,%p \n", s1, s1)
	s1 = append(s1, 1, 2, 3)
	Printf("%v %p\n", s1, s1)

	a := []int{1, 2, 3, 4, 5}
	s2 := a[2:5]
	s3 := a[1:3]
	Printf("%v,%p \n", s2, s2)
	Printf("%v,%p \n", s3, s3)

	s3[1] = 100
	Println(s2, "\n", s3, "\n", a[2])
}

func array7() {
	Println("..............array7..........")
	var b *[3]int = &[3]int{1, 2, 3}
	s(b)
}
func s(p *[3]int) {
	Println(*p)
}

//--------------------------------------
func array8() {
	Println("..............array8..........")
}

func array9() {
	Println("..............array9..........")
}

func array10() {
	Println("..............array10..........")
}

// 以下为一些常规的数组声明方法：
// [32]byte // 长度为32的数组，每个元素为一个字节
// [2*N] struct { x, y int32 } // 复杂类型数组
// [1000]*float64 // 指针数组
// [3][5]int // 二维数组
// [2][2][2]float64 // 等同于[2]([2]([2]float64))
// 从以上类型也可以看出，数组可以是多维的，比如[3][5]int就表达了一个3行5列的二维整
// 型数组，总共可以存放15个整型元素。

// 在Go语言中，数组长度在定义后就不可更改，在声明时长度可以为一个常量或者一个常量
// 表达式（常量表达式是指在编译期即可计算结果的表达式）。数组的长度是该数组类型的一个内
// 置常量，可以用Go语言的内置函数len()来获取。下面是一个获取数组arr元素个数的写法：
// arrLength := len(arr)

//第一种
//var <数组名称> [<数组长度>]<数组元素>
// var arr [2]int
//     arr[0]=1
//     arr[1]=2
//
// //第二种
// //var <数组名称> = [<数组长度>]<数组元素>{元素1,元素2,...}
// var arr = [2]int{1,2}
// //或者
// arr := [2]int{1,2}
//
// //第三种
// //var <数组名称> [<数组长度>]<数组元素> = [...]<元素类型>{元素1,元素2,...}
// var arr = [...]int{1,2}
// //或者
// arr := [...]int{1,2}
//
// //第四种
// //var <数组名称> [<数组长度>]<数组元素> = [...]<元素类型>{索引1:元素1,索引2:元素2,...}
// var arr = [...]int{1:1,0:2}
// //或者
// arr := [...]int{1:1,0:2}
