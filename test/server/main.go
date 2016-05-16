package main

import (
	"log"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	thriftx "github.com/chanxuehong/thrift"
	"github.com/chanxuehong/thrift/test/go.thrift/test"
)

func main() {
	var (
		ThriftListenAddr          = ":9999"
		ThriftTransportBufferSize = 1024
	)
	processor := thriftx.NewServiceProcessor(test.NewTestServiceProcessor(&serviceHandler{})) // NOTE: 就是在这里
	serverTransport, err := thrift.NewTServerSocketTimeout(ThriftListenAddr, time.Second*5)
	if err != nil {
		log.Println(err)
		return
	}
	transportFactory := thrift.NewTBufferedTransportFactory(ThriftTransportBufferSize)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	if err = server.Serve(); err != nil {
		log.Println(err)
		return
	}
}
