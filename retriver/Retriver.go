package retriver

/**
* 实现Retriver接口
*/
type Retriver struct {
    Contents string
}

func (r Retriver) Get(url string) string {
    return r.Contents
}