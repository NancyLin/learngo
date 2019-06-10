// 4-3、 扩展已有类型

// 如何扩充系统类型或者别人的类型
// 定义别名
// 使用组合

// 如果想要在原有的 tree.Node 基础上使用其他遍历方法
// 使用小写，private不给其他人调用
package main

import (
    "./tree"//在GOROOT和GOPATH目录中找
    "./queue"
    "fmt"
)

// 用组合的方式扩展
type myTreeNode struct {
    node *tree.Node
}

func (myNode *myTreeNode) postOrder(){
    if myNode == nil || myNode.node == nil {
        return
    }

    left := myTreeNode{myNode.node.Left}
    right := myTreeNode{myNode.node.Right}

    left.postOrder()
    right.postOrder()
    myNode.node.Print()

}

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

    myNode := myTreeNode{&root}
    myNode.postOrder()//5 0 0 5 200
    fmt.Println()

    queue := queue.Queue{};
    queue.Push(3)
    queue.Push(2)
    fmt.Println(queue)//[3 2]
    fmt.Println(queue.Pop())//3
    fmt.Println(queue.IsEmpty())//false
    fmt.Println(queue.Pop())//2
    fmt.Println(queue.IsEmpty())//false

}