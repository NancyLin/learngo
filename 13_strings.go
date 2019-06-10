package main

import (
    "fmt"
    "unicode/utf8"
)

// 十三
func main(){

    s := "Yes我爱慕课网!" 
    fmt.Println(s)//Yes我爱慕课网!

    //将他们的字符输出来
    for _, ch := range []byte(s) {
        fmt.Printf("%X ", ch)//59 65 73 E6 88 91 E7 88 B1 E6 85 95 E8 AF BE E7 BD 91 21
    }
    fmt.Println()
    
    //ch实际是rune，是int(32)
    for i, ch := range s {
        fmt.Printf("(%d %X) ", i, ch)//(0 59) (1 65) (2 73) (3 6211) (6 7231) (9 6155) (12 8BFE) (15 7F51) (18 21)
    }
    fmt.Println()
    
    //可以用RuneCountInString得到字符串有多少个字符
    fmt.Println("Rune count:", utf8.RuneCountInString(s)) //Rune count: 9

    //可以用utf8.DecodeRune()来得到每个字符和字符大小
    bytes := []byte(s)
    for len(bytes) > 0 {
        ch, size := utf8.DecodeRune(bytes)
        bytes = bytes[size:]
        fmt.Printf("%c ", ch)
    }
    //上述最终打出
    //Y e s 我 爱 慕 课 网 !

    fmt.Println()

    //如果要按字符顺序键值打印出来
    for i, ch := range []rune(s) {
        fmt.Printf("(%d %c) ", i, ch)
    }
    //上述最终打出
    //(0 Y) (1 e) (2 s) (3 我) (4 爱) (5 慕) (6 课) (7 网) (8 !) 
    fmt.Println()
}

// rune相当于go的char
// 使用range 遍历pos, rune对
// 使用utf8.RuneCountInString获得字符数量
// 使用len获得字节长度
// 使用[]byte获得字节

// 其他字符串操作
// Fields, Split, Join
// Contains, Index
// ToLower, ToUpper
// Trim, TrimRight, TrimLeft
