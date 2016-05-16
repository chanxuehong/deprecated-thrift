package thrift

import (
	"log"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type TProcessor interface {
	thrift.TProcessor
	GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool)
}

var _ thrift.TProcessor = (*ServiceProcessor)(nil)

type ServiceProcessor struct {
	processor TProcessor
}

func NewServiceProcessor(processor TProcessor) *ServiceProcessor {
	return &ServiceProcessor{
		processor: processor,
	}
}

func (p *ServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.processor.GetProcessorFunction(name); ok {
		start := time.Now()
		success, err = processor.Process(seqId, iprot, oprot)
		end := time.Now()
		log.Printf("%32s %16v %9t %v\n", name, end.Sub(start), success, err)
		return
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x11 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x11.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x11
}
