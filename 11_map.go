package main

import "fmt"

// 十一 map
// map[K]V map的定义，K代表key值的类型，V代表value值的类型
/*m := map[string]string {
    "name": "ccmouse",
    "course": "goland",
    "site": "imooc",
    "quality": "notbad",
}*/

// map也可以做复合定义，map[K1]map[K2]V，K1代表key值的类型，value值也是一个map类型

func main(){

    //map的建立
    m := map[string]string {
        "name": "ccmouse",
        "course": "goland",
        "site": "imooc",
        "quality": "notbad",
    }

    fmt.Println(m)//map[name:ccmouse course:goland site:imooc quality:notbad]

    //map的建立方法二
    m2 := make(map[string]int) // m2 = empty map
    fmt.Println(m2)//map[]

    //map的建立方法三
    var m3 map[string]int // m3 == nil
    fmt.Println(m3)//map[]

    //map的遍历
    fmt.Println("Traversing map")
    for k, v := range m {//只要key, 可以使用k := range m, 只要value, 可以用 _, v := range m
        fmt.Println(k, v)
    }
    // 上述遍历结果为，可能运行几次打印出来的结果顺序会不一样，这是因为map是哈希无序的
    /*name ccmouse
    course goland
    site imooc
    quality notbad*/

    //map 的操作
    fmt.Println("Getting value")
    courseName := m["course"]
    fmt.Println(courseName)//goland

    //map 是允许取不存在的key的value的，当没有这个key的value，直接默认为zero value，value类型是string的话则zero value是空串
    causeName := m["cause"]
    fmt.Println(causeName)// 返回空串

    //如果要判断key值是否存在map里面，可以都加一个值承接
    courseName2, ok := m["course"]
    fmt.Println(courseName2, ok)//goland true

    if causeName2, ok := m["cause"]; ok {
        fmt.Println(causeName2)
    }else{
        fmt.Println("Key does not exists")
    }
    //上述输出 Key does not exists

    //删除map中的值
    fmt.Println("Deleting value")
    name, ok := m["name"]
    fmt.Println(name, ok)//ccmouse true
    delete(m, "name")
    fmt.Println(m)//map[site:imooc quality:notbad course:goland]
    //删除后确认是否还能再取出
    name, ok = m["name"]
    fmt.Println(name, ok)// false
    

    //map 的操作总结
    // 1、创建： make(map[string]int)
    // 2、获取元素： m[key]
    // 3、key不存在时，获得Value类型的初始值
    // 4、用 value, ok := m[key]来判断是否存在key
    // 5、用delete删除一个key
    // 6、使用range来遍历key，或者遍历key,value对
    // 7、不保证遍历顺序，如需顺序，需手动对key排序，需要将其放到slice里面进行排序
    // 8、用len来获得map元素的个数

    //map的key
    // 1、map使用哈希表，必须可以比较相等
    // 2、除了slice、map、function的内建类型都可以作为key
    // 3、Struct类型不包含上述字段，也可作为key


}