package main

import (
    "fmt"
    "strconv"
    "os"
    "bufio"
)

//五、循环语句

func convertToBin(n int) string {
    result := ""
     //for的条件里可以省略初始条件，结束条件，递增表达式
    for ; n > 0; n /= 2 {
        lsb := n % 2
        result = strconv.Itoa(lsb) + result
    }
    return result
}

//打开文件，一行行读取文件
func printFile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }

    //scanner才能按行读取文件
    scanner := bufio.NewScanner(file)

    //如果没有初始条件和结束条件，只有递增条件，也可以不加分号
    //就相当于其他语言的while语句，go语言没有while语句
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}

func forever() {
    //for还可以省略递增条件，所有的条件都不加就是死循环
    //为什么把死循环设计得这么好写，是由于go语言经常要使用到，并发编程
    for {
        fmt.Println("abc")
    }
}

func main() {

    //for的条件里不需要括号
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }

    fmt.Println(
        convertToBin(5),
        convertToBin(13),
        convertToBin(72387885),
        convertToBin(0),
    )

    printFile("abc.txt");
   
}

//基本语法要点回顾
// for, if后面的条件没有括号
// if 条件里也可以定义变量
// 没有while
// switch不需要break, 也可以直接switch多个条件