package main

import "fmt"

//九 切片
//半开半闭区间
// arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
// s := arr[2:6]
// s的值为[2, 3, 4, 5]
// 
func main() {
    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

    fmt.Println("arr[2:6] = ", arr[2:6])//[2 3 4 5]
    fmt.Println("arr[2:] = ", arr[2:])//[2 3 4 5 6 7]
    fmt.Println("arr[:6] = ", arr[:6])//[0 1 2 3 4 5]
    fmt.Println("arr[:] = ", arr[:])//[0 1 2 3 4 5 6 7]

    s1 := arr[2:]
    fmt.Println("s1 = ", s1)//[2 3 4 5 6 7]
    s2 := arr[:]
    fmt.Println("s2 = ", s2)//[0 1 2 3 4 5 6 7]

    fmt.Println("After updateSlice s1:")
    updateSlice(s1)
    fmt.Println("s1 = ", s1)//[100 3 4 5 6 7]
    fmt.Println("arr = ", arr)//[0 1 100 3 4 5 6 7]

    fmt.Println("After updateSlice s2:")
    updateSlice(s2)
    fmt.Println("s2 = ", s2)//[100 1 100 3 4 5 6 7]
    fmt.Println("arr = ", arr)//[100 1 100 3 4 5 6 7]
    
    //定义数组，未初始值时，默认int类型为0 [0, 0, 0, 0, 0]
    var arr1 [5]int
    //由编译器自己识别数组长度
    arr3 := [...]int{2, 4, 6, 8, 10}

    fmt.Println("printArray2(arr1)");
    printArray2(arr1[:])
    fmt.Println("printArray2(arr3)");
    printArray2(arr3[:])
    fmt.Println("fmt.Println(arr1, arr3)");
    fmt.Println(arr1, arr3)//[100 0 0 0 0] [100 4 6 8 10]

    fmt.Println("Reslice")
    fmt.Println(s2)//[100 1 100 3 4 5 6 7]
    s2 = s2[:5]
    fmt.Println(s2)//[100 1 100 3 4]
    s2 = s2[2:]
    fmt.Println(s2)//[100 3 4]

    fmt.Println("Extension Slice")
    arr[0], arr[2] = 0, 2
    fmt.Println("arr = ", arr)//[0 1 2 3 4 5 6 7]
    s1 = arr[2:6]
    fmt.Println("s1 = ", s1)//[2 3 4 5]
    s2 = s1[3:5]
    fmt.Println("s2 = ", s2)//[5 6], 为什么s1只有4个元素却能取出第5个元素，slice是对底层array的一个view，实际它还有保存了arr的下标

    //注意：Slice的实现
    //Slice中有3个组成部分，
    // ptr 指向slice开头的元素
    // len 声明slice长度，取值时只能取到len长度的值，取值下标大于等于len都会报错下标越界
    // cap 代表整个array从ptr开始到结束整个的长度

    //slice可以向后扩展，不可以向前扩展
    //s[i]不可以超越len(s), 向后扩展不可以超越底层数组cap(s)，即s2[3:7]是不允许的

    //取slice的长度已经cap长度
    fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))//s1 = [2 3 4 5], len(s1) = 4, cap(s1) = 6
    fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d\n", s2, len(s2), cap(s2))//s2 = [5 6], len(s2) = 2, cap(s2) = 3

    //slice 操作
    fmt.Println("arr = ", arr)//arr =  [0 1 2 3 4 5 6 7]
    s3 := append(s2, 10)
    s4 := append(s3, 11)
    s5 := append(s4, 12)
    fmt.Println("s3, s4, s5 =", s3, s4, s5)//s3, s4, s5 = [5 6 10] [5 6 10 11] [5 6 10 11 12]
    // s4 and s5 no longer view arr
    fmt.Println("arr = ", arr)//arr = [0 1 2 3 4 5 6 10]

    //注意：向slice添加元素
    // 添加元素时如果超越cap，系统会重新分配更大的底层数组
    // 由于值传递的关系，必须接收append的返回值
    // s = append(s, val)
}
//切片作为参数传递，改变里面的值，原本的值也会一起改变
func updateSlice(s []int){
    s[0] = 100
}

//对于上一节数组作为参数更改数组的元素值，可以改写为如下
func printArray2(arr []int){
    arr[0] = 100
    for i, v := range arr {
        fmt.Println(i, v)
    }
}

//注意
// Slice本身是没有数据的，是对底层array的一个view
