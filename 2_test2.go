
//二、内建变量类型

//(1) bool, string
//(2) 整数类型 (u)int, (u)int8, (u)int6, (u)int32, (u)int64, uintptr指针
//不规定长度的，就与操作系统相关，32位或者64位
//byte, rune(字符型，所说的char类型，用int32位来代表，整数的别名)
//浮点数和复数 float32, float64, complex64, complex128
//复数 i = 根号-1， i的平方=-i
//复数 3+4i，3为实数，4为复数
//3+4i取模，|3+4i| = 根号(3的平方+4的平方)
//go语言只有强制类型装换，没有隐示类型转换