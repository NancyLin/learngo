// 十五、 封装
// 名字一般使用类似CamelCase的形式，多个单词首字母大写，没有下划线的形式
// 如果是首字母大写，则代表是 pulic
// 如果是首字母小写，则代表是 private

// public 和 private 都是针对包的

// 每个目录一个包，但包名不一定和目录名一样
// main包包含可执行入口
// 为结构定义的方法必须放在同一个包内，但可以是不同的文件

package main

import (
    "./tree"
    "fmt"
)

func main() {
    var root tree.Node

    root = tree.Node{Value: 3}
    root.Left = &tree.Node{}
    root.Right = &tree.Node{5, nil, nil}
    root.Right.Left = new(tree.Node)
    root.Left.Right = tree.CreateNode(5)

    root.SetValue(200)

    root.Traverse()//0 5 200 0 5 
    fmt.Println()
}