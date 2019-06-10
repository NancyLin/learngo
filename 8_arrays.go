package main

import "fmt"

// 八、数组
func main() {
    //定义数组，未初始值时，默认int类型为0 [0, 0, 0, 0, 0]
    var arr1 [5]int
    //可以用 := ，但必须给予初值
    arr2 := [3]int{1, 3, 5}
    //由编译器自己识别数组长度
    arr3 := [...]int{2, 4, 6, 8, 10}

    fmt.Println(arr1, arr2, arr3)

    // 定义二维数组
    var grid [4][5]int
    fmt.Println(grid)

    //遍历数组
    for i := 0; i < len(arr3); i++ {
        fmt.Println(arr3[i])
    }

    //遍历数组的方式2，使用range关键字
    for i := range arr3 {
        fmt.Println(arr3[i])
    }

    //遍历数组的键值和值
    for i, v := range arr3 {
        fmt.Println(i, v)
    }

    //如果只要值不要键值，可通过下划线来省略逗号，不仅range，任何地方都可以通过 _省略变量
    for _, v := range arr3 {
        fmt.Println(v)
    }

    //只能使用arr1 和 arr3作为参数传送进去，arr2长度为3，不符合参数长度
    fmt.Println("printArray(arr1)")
    printArray(arr1)
    //printArray(arr2) arr2 不能作为参数
    fmt.Println("printArray(arr3)")
    printArray(arr3)

    //指针传递数组
    fmt.Println("printArray2(arr1)");
    printArray2(&arr1)
    fmt.Println("printArray2(arr3)");
    printArray2(&arr3)
    fmt.Println("fmt.Println(arr1, arr3)");
    fmt.Println(arr1, arr3)
}

//数组作为函数传参，需要写明数组的长度，没有写明数组的长度的，代表的是切片，[5]int 代表数组，[]int 代表切片
func printArray(arr [5]int){
    for i, v := range arr {
        fmt.Println(i, v)
    }
}

//1. 数组是值类型
//2. [10]int 和 [20]int 是不同类型
//3. 调用func f(arr [10]int) 会拷贝数组
//4. 在go语言一般不直接使用数组，使用的是切片

// 可以使用指针参数进行地址传递，不过这种方法不常用
func printArray2(arr *[5]int){
    arr[0] = 100
    for i, v := range arr {
        fmt.Println(i, v)
    }
}