package main

import "fmt"
// 十四 struct

// go语言仅支持封装，不支持继承和多态

//go 没有class，只有struct
//无论地址还是结构，一律使用.来访问成员

type treeNode struct {
    value int
    left, right *treeNode
}

//为结构体定义方法，括号中代表接收者
func (node treeNode) print(){
    fmt.Print(node.value, " ")
}

/*//结构体的方式其实也是一个普通方法，上述的功能也可以写为
func print(node treeNode) {
    fmt.Println(node.value)
}
//只是这种写法的调用为 print(root)*/

//调用结构体的传参都是值传参
func (node treeNode) setValue(value int){
    node.value = value
}

//如果要可以修改结构体，传参时需要使用指针传参，nil指针也可以调用方法
func (node *treeNode) setValue2(value int){
    node.value = value
}

//遍历所有节点
func (node *treeNode) traverse(){
    if node == nil {
        return
    }

    node.left.traverse()
    node.print()
    node.right.traverse()
}

//go语言本身没有构造函数，可以使用自定义工厂函数，注意返回了局部变量的地址
func createNode(value int) *treeNode {
    return &treeNode{value: value}
}

func main() {
    //创建struct
    var root treeNode
    fmt.Println(root)//{0 <nil> <nil>}
    
    //初始化元素，多种创建结构体的方法
    root = treeNode{value: 3}
    root.left = &treeNode{}
    root.right = &treeNode{5, nil, nil}
    root.right.left = new(treeNode)

    //使用工厂函数来创建
    root.left.right = createNode(5)
    fmt.Println(root)

    //节点数组
    nodes := []treeNode {
        {value: 3},
        {},
        {5, nil, &root},
    }

    fmt.Println(nodes)//[{3 <nil> <nil>} {0 <nil> <nil>} {5 <nil> 0xc00000a080}]

    root.print()//3
    fmt.Println()

    //调用方法设置value值，调用struct传参都是值传递
    root.right.left.setValue(3)
    root.right.left.print()// 0
    fmt.Println()

    root.print()
    root.setValue2(100)

    pRoot := &root
    pRoot.print()//100
    pRoot.setValue2(200)
    pRoot.print()//200
    fmt.Println()
    
    //遍历root节点，从左节点开始打印
    root.traverse()// 0 5 200 0 5 
    fmt.Println()

}

// 值接收者 vs 指针接收者
// 1. 要改变内容必须使用指针接收者
// 2. 结构过大也考虑使用指针接收者
// 3. 如果有指针接收者，最好都是指针接收者

// 值接收者是go语言特有
// 值/指针接收者均可接收值/指针
