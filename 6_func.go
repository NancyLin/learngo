package main

import (
    "fmt"
    "reflect"
    "runtime"
    "math"
)

// 六、 函数
//定义类似 func eval(a, b int, op string) int
func eval(a, b int, op string) int {
    switch op {
    case "+":
        return a + b
    case "-":
        return a - b
    case "*":
        return a * b
    case "/":
        return a / b
    default:
        panic("unsupported operation:" + op)
    }
}

//函数也可以返回多个值
func div(a, b int) (int, int){
    return a / b, a % b
}

//也可以给返回值取名字
func div2(a, b int) (q, r int){
    //这种方式，如果函数体比较长，这种方式不方便读代码，不推荐使用
    q = a / b
    r = a % b
    return
}

//多返回值一般用来返回错误，返回一个值和一个error
func eval2(a, b int, op string) (int, error) {
    switch op {
    case "+":
        return a + b, nil
    case "-":
        return a - b, nil
    case "*":
        return a * b, nil
    case "/":
        return a / b, nil
    default:
        return 0, fmt.Errorf("unsupported operation: %s", op)
    }
}

//函数的参数也可以是一个函数
func apply(op func(int, int) int, a, b int ) int {
    //反射得到函数的指针
    p := reflect.ValueOf(op).Pointer()
    //得到函数名
    opName := runtime.FuncForPC(p).Name()
    fmt.Printf("calling function %s with args (%d, %d)", opName, a, b )

    return op(a, b)
}

func pow(a, b int) int {
    return int(math.Pow(float64(a), float64(b)))
}

//函数有可变参数列表
func sum(numbers ...int) int {
    sum := 0
    for i := range numbers {
        sum += numbers[i]
    }

    return sum
}

func main() {
    fmt.Println(eval(3, 4, "*"))
    fmt.Println(div(13, 3))

    q, r := div2(13, 4)
    fmt.Println(q, r)

    //如果只想读取返回的1个值，可以使用如下
    a, _ := div2(15, 2)
    fmt.Println(a)

    //多个返回值
    fmt.Println(eval2(3, 4, "X"))

    //如果要判断错误，调用方式可以
    if result, err := eval2(3, 4, "X"); err != nil {
        fmt.Println("Error: ", err)
    }else {
        fmt.Println(result)
    }

    fmt.Println(apply(pow, 3, 4))

    //还可以直接写匿名函数
    //返回 calling function main.main.func1 with args (3, 4)81
    fmt.Println(apply(
        func(a, b int) int {
            return int(math.Pow(float64(a), float64(b)))
        }, 3, 4))

    fmt.Println(sum(1, 2, 3, 4))

    a, b := 3, 4
    swap(&a, &b)
    fmt.Println(a, b)

    a2, b2 := 3, 4
    a2, b2 = swap2(a2, b2)
    fmt.Println(a2, b2)
}


//总结
// 函数可返回多个值
// 函数返回多个值可以起名字， 但对于调用者而言没有区别，仅用于非常简单的函数
// 函数作为参数
// 可变参数列表

// 语法要点
// 返回参数写在最后面
// 可返回多个值
// 函数作为参数
// 没有默认参数，可选参数，也没有函数重载，只有可变参数列表


// 七、 指针
//交换两个变量的值
func swap(a, b *int){
    *a, *b = *b, *a
}

//上述交换变量值写法也可以是
func swap2(a, b int) (int, int){
    return b, a
}

// 上述交换变量的写法也可以是
