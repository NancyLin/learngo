package main

import "fmt"

//十二、map练习题
// 例：寻找最长不含有重复字符的子串
// 对于每一个字母x
// lastOccurred[x]不存在，或者<start，则无需操作
// lastOccurred[x] >= start, 则需要更新start为x出现的位置+1
// 更新lastOccurred[x], 更新maxLength

func lengthOfNonRepeatingSubStr(s string) int {

    lastOccurred := make(map[byte]int)
    start := 0
    maxLength := 0
    //由于字符是int32的类型，可以将其转为[]byte
    for i, ch := range []byte(s){
        //如果最后一次出现的位置大于start的位置，更新start
        if lasI, ok := lastOccurred[ch]; ok && lastOccurred[ch] >= start {
            start = lasI + 1
        }

        //更新目前判断不重复的长度
        if i - start + 1 > maxLength {
            maxLength = i - start + 1
        }

        lastOccurred[ch] = i
    }
    
    return maxLength
}

//适合各种字符的版本
func  lengthOfNonRepeatingRuneSubStr(s string) int {
    
    lastOccurred := make(map[rune]int)
    start := 0
    maxLength := 0

    for i, ch := range []rune(s) {
        if lasI, ok := lastOccurred[ch]; ok && lastOccurred[ch] >= start {
            start = lasI + 1
        }

        if i - start + 1 > maxLength {
            maxLength = i - start + 1
        }

        lastOccurred[ch] = i
    }

    return maxLength   
}

func main() {
    fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))//3
    fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))//1
    fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))//3
    fmt.Println(lengthOfNonRepeatingSubStr(""))//0
    fmt.Println(lengthOfNonRepeatingSubStr("b"))//1
    fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))//6
    fmt.Println(lengthOfNonRepeatingSubStr(""))//0
    fmt.Println(lengthOfNonRepeatingSubStr("这是慕课网"))//9 由于将其转为字节，有些中文占了3个字节，会有unicode的问题

    fmt.Println(lengthOfNonRepeatingRuneSubStr("这是慕课网"))//5
    fmt.Println(lengthOfNonRepeatingRuneSubStr("13我们93!一二三二一"))//8
}