package retriver

import (
    "time"
    "net/http"
    "net/http/httputil"
    )

/**
* 实现Retriver接口
*/
type RetriverReal struct {
    UserAgent string
    TimeOut time.Duration
}

func (r RetriverReal) Get(url string) string {

    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }

    result, err := httputil.DumpResponse(resp, true)

    resp.Body.Close()

    if err != nil {
        panic(err)
    }

    return string(result)
}
