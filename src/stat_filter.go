package main

import (
	"fmt"
	"time"

	"github.com/hprose/hprose-golang/rpc"
)

type statFilter struct {
	Message string
}

func (sf statFilter) handler(
	request []byte,
	context rpc.Context,
	next rpc.NextFilterHandler) (response []byte, err error) {
	start := time.Now()
	response, err = next(request, context)
	end := time.Now()
	fmt.Printf("%v takes %d ms.\r\n", sf.Message, end.Sub(start)/time.Millisecond)
	return
}
