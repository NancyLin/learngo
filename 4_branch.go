package main

import "fmt"
import "io/ioutil"

//四、条件语句  if结构
func bounded(v int) int {
    //if 条件里不需要括号
    if v > 100 {
        return 100
    }else if v < 0 {
        return 0
    }else {
        return v
    }
}

func main() {
    const filename = "abc.txt"
    //读取文件内容，可以使用ioutil.ReadFile, 函数返回[]byte（文件内容）, error（出错信息）
    //content, err := ioutil.ReadFile(filename)
    //没有错误的，error返回的是nil
    //if err != nil {
    //    fmt.Println(err)
    //}else{
    //    fmt.Printf("%s\n", content)
    //}

    //上述还能使用其他写法
    if content, err := ioutil.ReadFile(filename); err != nil {

        fmt.Println(err)

    }else{

        fmt.Printf("%s\n", content)
    }
    //但上述这个写法，content就无法在除上述的代码块中使用，例如下面，会提示undefined: content
    //fmt.Printf("%s\n", content)

    fmt.Println(
        grade(0),
        grade(59),
        grade(60),
        grade(82),
        grade(99),
        grade(100),
        grade(101),//这个会根据panic报错
    )
}

//总结
// if 的条件里可以赋值，赋值后再判断，if条件里赋值的变量作用域就在这个if语句里

// switch 语句
func eval(a, b int, op string) int {
    var result int
    // switch 会自动break，除非使用fallthrough
    switch op {
        case "+":
            result = a + b
        case "-":
            result = a - b
        case "*":
            result = a * b
        case "/":
            result = a / b
        default:
            panic("unsupported operator:" + op)
    }

    return result
}

//另外一种switch的用法
func grade(score int) string {
    g := ""
    //switch后可以没有表达式，将条件放到case中
    switch {
        case score < 0 || score > 100:
            panic(fmt.Sprintf("Wront score: %d", score))//panic会中断程序，直接报错
        case score < 60:
            g = "F"
        case score < 80:
            g = "C"
        case score < 90:
            g = "B"
        case score <= 100:
            g = "A"
    }
    return g
}