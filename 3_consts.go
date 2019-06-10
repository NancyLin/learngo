package main

import "fmt"
import "math"

//三、常量的定义

func consts(){
    const filename = "abx.txt"
    //常量如果不定义类型，就是一个文本
    const a, b = 3,4 
    //上面的写法也可以定义为
    //const (
    //    filename = "abx.txt"
    //    a, b = 3, 4
    //)
    var c int
    //如果没有指定a,b是整型，由于Sqrt要求参数是float64，则需要将参数强制转为float64，如果没有指定是整型，则常量是文本格式，
    c = int(math.Sqrt(a * a + b * b))

    fmt.Println(c)
}

func  main() {
    consts()
    enums()
}

//总结
//定义 const filename = "abc.txt"
//const数值可作为各种类型使用，当它作为函数参数，会直接根据函数参数类型转换
//例如：const a, b = 3, 4
// var c int = int(math.Sqrt(a * a + b * b))

//使用常量定义枚举类型
//没有特殊的枚举关键字，用的const
func enums(){
    //普通枚举类型
    /*const(
        cpp    = 0
        java   = 1
        python = 2
        goland = 3
    )*/
    //也可以用itoa来表示这组枚举是自增值的,例如：
    /*const(
        cpp = iota
        java
        python
        goland
    )*/
    //自增值的枚举类型，如果想中间跳过一个变量，也可以
    const(
        cpp = iota
        _
        java
        python
        goland
        javascript
    )
    //定义 b, kb, mb, gb, tb, pb, iota可以用于表达式
    const(
        b = 1 << (10 * iota)
        kb
        mb
        gb
        tb
        pb
    )
    fmt.Println(cpp, javascript, python, goland)
    fmt.Println(b, kb, mb, gb, tb, pb)
}

//变量定义要点回顾
//变量类型写在变量名的后面
//编译器可推测变量类型
//没有char，只有rune，32位
//原生支持复数的类型
