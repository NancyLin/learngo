package main

import "fmt"

//十 切片的操作

func printSlice(s []int){
    fmt.Printf("%v, len = %d, cap = %d\n", s, len(s), cap(s))
}

func main() {

    fmt.Println("Creating slice")
    //slice 的创建
    var s []int//Zero value for slice is nil

    for i := 0; i < 10; i++{
        s = append(s, 2 * i + 1)
        printSlice(s)
    }

    /* 上述运行结果如下：
    [1], len = 1, cap = 1
    [1 3], len = 2, cap = 2
    [1 3 5], len = 3, cap = 4
    [1 3 5 7], len = 4, cap = 4
    [1 3 5 7 9], len = 5, cap = 8
    [1 3 5 7 9 11], len = 6, cap = 8
    [1 3 5 7 9 11 13], len = 7, cap = 8
    [1 3 5 7 9 11 13 15], len = 8, cap = 8
    [1 3 5 7 9 11 13 15 17], len = 9, cap = 16
    [1 3 5 7 9 11 13 15 17 19], len = 10, cap = 16
    [1 3 5 7 9 11 13 15 17 19]
    */

    fmt.Println(s)//[1 3 5 7 9 11 13 15 17 19]

    // slice创建：初始一个已知值的slice
    s1 := []int{2, 4, 6, 8}
    printSlice(s1)//[2 4 6 8], len = 4, cap = 4

    // slice创建：已知长度，但没有初始值
    s2 := make([]int, 16)
    printSlice(s2)//[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len = 16, cap = 16

    // 或者创建时，可以定义它的cap
    s3 := make([]int, 10, 32)
    printSlice(s3)//[0 0 0 0 0 0 0 0 0 0], len = 10, cap = 32
    
    // copy slice
    fmt.Println("Copying slice")
    copy(s2, s1)
    printSlice(s2)//[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0], len = 16, cap = 16

    //删除slice
    fmt.Println("Deleting elements from slice")
    //例如删除 8 ，由于append第二个参数是可变参数，即传参的方式为 0, 0, 0不能直接传入slice
    s2 = append(s2[:3], s2[4:]...)
    printSlice(s2)//[2 4 6 0 0 0 0 0 0 0 0 0 0 0 0], len = 15, cap = 16

    fmt.Println("Poping from front")
    front := s2[0]
    s2 = s2[1:]
    fmt.Println(front)//2
    printSlice(s2)//[4 6 0 0 0 0 0 0 0 0 0 0 0 0], len = 14, cap = 16

    fmt.Println("Poping from back")
    tail := s2[len(s2) - 1]
    s2 = s2[:len(s2) - 1]
    fmt.Println(tail)//0
    printSlice(s2)//[4 6 0 0 0 0 0 0 0 0 0 0 0], len = 13, cap = 16


}