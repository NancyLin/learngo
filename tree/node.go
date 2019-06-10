package tree

import "fmt"

//大写代表public，元素变量方法都一样，如果首字母是大写代表是public， 首字母小写代表是private
type Node struct {
    Value int
    Left, Right *Node
}

func CreateNode(value int) *Node {
    return &Node{Value: value}
}

// 定义设置的方法
func (node *Node) SetValue(value int){
    node.Value = value
}

// 定义设置打印的方法
func (node Node) Print() {
    fmt.Print(node.Value, " ")
}

//结构体的不同方法也可以写在不同文件中，遍历节点的方法traverse的方法就写在traverse.go的文件中