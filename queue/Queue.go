package queue

/**
* 使用别名的方式扩展
*/
type Queue []int

func (q *Queue) Push(value int) {
    *q = append(*q, value)
}

func (q *Queue) Pop() int{
    head := (*q)[0]
    *q = (*q)[1:]
    return head
}

func (q *Queue) IsEmpty() bool{
    return len(*q) == 0
}