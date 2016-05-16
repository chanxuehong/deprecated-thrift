## thrift rpc 服务时间统计
```
2016/05/16 14:44:19                              add          24.21µs      true <nil>
2016/05/16 14:44:19                              add         16.186µs      true <nil>
2016/05/16 14:44:19                              add         15.003µs      true <nil>
2016/05/16 14:44:19                              add         23.475µs      true <nil>
2016/05/16 14:44:19                              add          19.16µs      true <nil>
2016/05/16 14:44:19                              add         30.473µs      true <nil>
2016/05/16 14:44:19                              add         18.909µs      true <nil>
2016/05/16 14:44:19                              add         15.044µs      true <nil>
2016/05/16 14:44:19                              add         15.306µs      true <nil>
2016/05/16 14:44:19                              add         16.332µs      true <nil>
2016/05/16 14:44:19                              add         29.563µs      true <nil>
2016/05/16 14:44:19                              add         16.499µs      true <nil>
2016/05/16 14:44:19                              add         26.008µs      true <nil>
2016/05/16 14:44:19                              add         25.315µs      true <nil>
2016/05/16 14:44:19                              add         17.123µs      true <nil>
2016/05/16 14:44:19                              add         29.919µs      true <nil>
2016/05/16 14:44:19                              add         21.628µs      true <nil>
2016/05/16 14:44:19                              add         19.308µs      true <nil>
2016/05/16 14:44:19                              add         35.516µs      true <nil>
2016/05/16 14:44:19                              add          15.68µs      true <nil>
```

### 使用方法
```Go
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
```