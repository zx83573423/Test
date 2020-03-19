package OutTime

import (
	"fmt"
	"time"
)

var g_chTest chan int
var timeOut time.Duration

func init() {
	g_chTest = make(chan int, 1)
	timeOut = time.Millisecond * 100
}

func Add(nTest int) {
	select {
	case g_chTest <- nTest:
		fmt.Println("添加成功")
	case <-time.After(timeOut):
		fmt.Println("添加超时")
	}
}
