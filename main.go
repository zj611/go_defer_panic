package main

import (
	"fmt"
	"runtime"
)

func main() {
	err := panicAndReturnErr()
	if err != nil{
		fmt.Printf("err is %+v\n", err)
	}
	fmt.Println("returned normally from panicAndReturnErr")
}
func panicAndReturnErr() (err error){
	defer func() {//如果defer内有recover，则从 panic 中恢复，不会终止程序
		if e := recover(); e != nil {//error occur!
			// 打印栈信息
			buf := make([]byte, 1024)
			buf = buf[:runtime.Stack(buf, false)]
			err = fmt.Errorf("[PANIC]%v\n%s\n", e, buf)
		}
	}()
	fmt.Println("panic begin")
	panic("error occur!")//开始执行defer函数
	fmt.Println("panic over")
	return nil
}
