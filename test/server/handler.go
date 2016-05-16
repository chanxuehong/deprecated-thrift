package main

import (
	"github.com/chanxuehong/thrift/test/go.thrift/test"
)

var _ test.TestService = (*serviceHandler)(nil)

type serviceHandler struct{}

func (*serviceHandler) Add(x int64, y int64) (r int64, err error) {
	return x + y, nil
}
