package main

import (
	"fmt"
	"log"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/chanxuehong/thrift/test/go.thrift/test"
)

func main() {
	var (
		ThriftServerAddr = "127.0.0.1:9999"
	)
	Socket, err := thrift.NewTSocket(ThriftServerAddr)
	if err != nil {
		log.Println(err)
		return
	}
	Transport := thrift.NewTBufferedTransport(Socket, 1024)

	if err = Transport.Open(); err != nil {
		log.Println(err)
		return
	}
	defer Transport.Close()

	ProtocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	clt := test.NewTestServiceClientFactory(Transport, ProtocolFactory)

	for i := int64(0); i < 20; i++ {
		fmt.Println(clt.Add(i, i))
	}
}
