//5-1、5-2 接口和接口的定义

// duck typing
// 描述事物的外部行为而非内部结构
// 严格说go属于结构化类型系统，类似duck typing

// 接口由使用者定义

package main

import (
    "fmt"
    "./retriver"
    )

//定义一个接口
type Retriver interface {
    //接口中的方法，不需要用func关键字，interface默认里面的都是函数
    Get(url string) string
}

func download(r Retriver) string {
    return r.Get("http://www.imooc.com")
}

func main() {
    var r Retriver
    r = retriver.Retriver{"this is a fake imooc.com"}
    fmt.Println(download(r))//this is a fake imooc.com

    r = retriver.RetriverReal{}
    fmt.Println(download(r))//返回http://www.imooc.com网页内容
}

// 接口的实现
// 接口的实现是隐式的
// 只要实现接口里的方法
