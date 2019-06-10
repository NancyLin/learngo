package main

import "fmt"

//一、怎么定义go语言的变量
func variable() {
    //（1）定义变量，var关键字，变量名写在前面，类型在后面
    // 定义变量后如果不赋值，则会有默认初值， int 默认为 0, string 默认为空串
    var a1 int
    var s1 string
    fmt.Println(a1, s1)
    // 空串打印不出来，可以用fmt.Printf加格式，%q可以打出空串（%s无法打出空串，它就是空串的格式）
    fmt.Printf("%d %q\n", a1, s1)

    //（2）变量也可以赋初值
    var a2 int = 3
    var s2 string = "abc"
    fmt.Println(a2, s2)

    //（3）变量也可以定义多个
    var a3, a3_1 int = 3, 4
    fmt.Println(a3, a3_1)

    //（4）go可以根据赋值，推断变量的类型，因此也可以不用显示地写明变量类型
    // 如果不写明类型，不同的类型也可以写在一行中
    var a4, a4_1, b4, s4 = 3, 4, true, "def"
    fmt.Println(a4, a4_1, b4, s4)

    //（5）上述的写法可能比较繁琐，go还有简单的写法 := 用来定义赋值，在函数外部是不能使用 := 简单方式来定义赋值变量
    a5, a5_1, b5, s5 := 3, 4, true, "def"
    fmt.Println(a5, a5_1, b5, s5)

}

//（6）函数外部定义变量不能使用 :=， 每个函数外部都需要以关键字开头，例如var，func
var a6 = 3
var s6 = "KKK"
//a6_1 := true 是错误的，由于没有关键字开头，所以无法使用

//（7）函数外部的变量并不是全局变量，只是包内部的变量

// (8) 上述的函数定义还有简便的写法
var (
    a8 = 4
    s8 = "KKK"
    a8_1 = true
)

//注意点：变量定义一定要用到
// 在函数中定义的变量只能函数使用
// 在函数的外部不能使用 := 来定义赋值变量

func main(){
    fmt.Println("Hello World")

    variable()
    //使用包变量
    fmt.Println(a6, s6, a8, s8, a8_1)
}

/*
变量定义总结：
（1）使用var关键字
 var a, b, c bool
 var s1, s2 string = "hello", "world"
 可以放在函数内，或直接放在包内
 使用var()集中定义变量
（2）让编译器自动决定类型
 var a, b, i, s1, s2 = true, false, 3, "hello", "world"
（3）使用 := 定义变量
 a, b, i, s1, s2 := true, false, 3, "hello", "world"
 只能在函数内使用
 */
