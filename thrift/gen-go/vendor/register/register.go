// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package register

import (
	"bytes"
	"fmt"
	"thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

// Attributes:
//  - Success
//  - Message
//  - Services
type Response struct {
  Success bool `thrift:"success,1" db:"success" json:"success"`
  Message string `thrift:"message,2" db:"message" json:"message"`
  Services map[string]*ServiceInfo `thrift:"services,3" db:"services" json:"services"`
}

func NewResponse() *Response {
  return &Response{}
}


func (p *Response) GetSuccess() bool {
  return p.Success
}

func (p *Response) GetMessage() string {
  return p.Message
}

func (p *Response) GetServices() map[string]*ServiceInfo {
  return p.Services
}
func (p *Response) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Response)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Success = v
}
  return nil
}

func (p *Response)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Message = v
}
  return nil
}

func (p *Response)  ReadField3(iprot thrift.TProtocol) error {
  _, _, size, err := iprot.ReadMapBegin()
  if err != nil {
    return thrift.PrependError("error reading map begin: ", err)
  }
  tMap := make(map[string]*ServiceInfo, size)
  p.Services =  tMap
  for i := 0; i < size; i ++ {
var _key0 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _key0 = v
}
    _val1 := &ServiceInfo{}
    if err := _val1.Read(iprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _val1), err)
    }
    p.Services[_key0] = _val1
  }
  if err := iprot.ReadMapEnd(); err != nil {
    return thrift.PrependError("error reading map end: ", err)
  }
  return nil
}

func (p *Response) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Response"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Response) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("success", thrift.BOOL, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:success: ", p), err) }
  if err := oprot.WriteBool(bool(p.Success)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.success (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:success: ", p), err) }
  return err
}

func (p *Response) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("message", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:message: ", p), err) }
  if err := oprot.WriteString(string(p.Message)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.message (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:message: ", p), err) }
  return err
}

func (p *Response) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("services", thrift.MAP, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:services: ", p), err) }
  if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRUCT, len(p.Services)); err != nil {
    return thrift.PrependError("error writing map begin: ", err)
  }
  for k, v := range p.Services {
    if err := oprot.WriteString(string(k)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    if err := v.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
    }
  }
  if err := oprot.WriteMapEnd(); err != nil {
    return thrift.PrependError("error writing map end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:services: ", p), err) }
  return err
}

func (p *Response) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Response(%+v)", *p)
}

// Attributes:
//  - Name
//  - Host
//  - Port
//  - Nonce
//  - Sign
//  - IsService
type ServiceInfo struct {
  Name string `thrift:"name,1" db:"name" json:"name"`
  Host string `thrift:"host,2" db:"host" json:"host"`
  Port int32 `thrift:"port,3" db:"port" json:"port"`
  Nonce string `thrift:"nonce,4" db:"nonce" json:"nonce"`
  Sign string `thrift:"sign,5" db:"sign" json:"sign"`
  IsService bool `thrift:"isService,6" db:"isService" json:"isService"`
}

func NewServiceInfo() *ServiceInfo {
  return &ServiceInfo{}
}


func (p *ServiceInfo) GetName() string {
  return p.Name
}

func (p *ServiceInfo) GetHost() string {
  return p.Host
}

func (p *ServiceInfo) GetPort() int32 {
  return p.Port
}

func (p *ServiceInfo) GetNonce() string {
  return p.Nonce
}

func (p *ServiceInfo) GetSign() string {
  return p.Sign
}

func (p *ServiceInfo) GetIsService() bool {
  return p.IsService
}
func (p *ServiceInfo) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    case 4:
      if err := p.ReadField4(iprot); err != nil {
        return err
      }
    case 5:
      if err := p.ReadField5(iprot); err != nil {
        return err
      }
    case 6:
      if err := p.ReadField6(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *ServiceInfo)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Name = v
}
  return nil
}

func (p *ServiceInfo)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Host = v
}
  return nil
}

func (p *ServiceInfo)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Port = v
}
  return nil
}

func (p *ServiceInfo)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Nonce = v
}
  return nil
}

func (p *ServiceInfo)  ReadField5(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 5: ", err)
} else {
  p.Sign = v
}
  return nil
}

func (p *ServiceInfo)  ReadField6(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(); err != nil {
  return thrift.PrependError("error reading field 6: ", err)
} else {
  p.IsService = v
}
  return nil
}

func (p *ServiceInfo) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("ServiceInfo"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
    if err := p.writeField5(oprot); err != nil { return err }
    if err := p.writeField6(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ServiceInfo) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:name: ", p), err) }
  if err := oprot.WriteString(string(p.Name)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.name (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:name: ", p), err) }
  return err
}

func (p *ServiceInfo) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("host", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:host: ", p), err) }
  if err := oprot.WriteString(string(p.Host)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.host (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:host: ", p), err) }
  return err
}

func (p *ServiceInfo) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("port", thrift.I32, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:port: ", p), err) }
  if err := oprot.WriteI32(int32(p.Port)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.port (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:port: ", p), err) }
  return err
}

func (p *ServiceInfo) writeField4(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("nonce", thrift.STRING, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:nonce: ", p), err) }
  if err := oprot.WriteString(string(p.Nonce)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.nonce (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:nonce: ", p), err) }
  return err
}

func (p *ServiceInfo) writeField5(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("sign", thrift.STRING, 5); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:sign: ", p), err) }
  if err := oprot.WriteString(string(p.Sign)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.sign (5) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 5:sign: ", p), err) }
  return err
}

func (p *ServiceInfo) writeField6(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("isService", thrift.BOOL, 6); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:isService: ", p), err) }
  if err := oprot.WriteBool(bool(p.IsService)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.isService (6) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 6:isService: ", p), err) }
  return err
}

func (p *ServiceInfo) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ServiceInfo(%+v)", *p)
}

type Register interface {
  Version() (r string, err error)
  // Parameters:
  //  - ServiceInfo
  Heartbeat(serviceInfo *ServiceInfo) (r *Response, err error)
}

type RegisterClient struct {
  Transport thrift.TTransport
  ProtocolFactory thrift.TProtocolFactory
  InputProtocol thrift.TProtocol
  OutputProtocol thrift.TProtocol
  SeqId int32
}

func NewRegisterClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *RegisterClient {
  return &RegisterClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewRegisterClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *RegisterClient {
  return &RegisterClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

func (p *RegisterClient) Version() (r string, err error) {
  if err = p.sendVersion(); err != nil { return }
  return p.recvVersion()
}

func (p *RegisterClient) sendVersion()(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("version", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := RegisterVersionArgs{
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *RegisterClient) recvVersion() (value string, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "version" {
    err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "version failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "version failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error3 error
    error3, err = error2.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error3
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "version failed: invalid message type")
    return
  }
  result := RegisterVersionResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}

// Parameters:
//  - ServiceInfo
func (p *RegisterClient) Heartbeat(serviceInfo *ServiceInfo) (r *Response, err error) {
  if err = p.sendHeartbeat(serviceInfo); err != nil { return }
  return p.recvHeartbeat()
}

func (p *RegisterClient) sendHeartbeat(serviceInfo *ServiceInfo)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("heartbeat", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := RegisterHeartbeatArgs{
  ServiceInfo : serviceInfo,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *RegisterClient) recvHeartbeat() (value *Response, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "heartbeat" {
    err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "heartbeat failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "heartbeat failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error5 error
    error5, err = error4.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error5
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "heartbeat failed: invalid message type")
    return
  }
  result := RegisterHeartbeatResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}


type RegisterProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler Register
}

func (p *RegisterProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *RegisterProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *RegisterProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewRegisterProcessor(handler Register) *RegisterProcessor {

  self6 := &RegisterProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self6.processorMap["version"] = &registerProcessorVersion{handler:handler}
  self6.processorMap["heartbeat"] = &registerProcessorHeartbeat{handler:handler}
return self6
}

func (p *RegisterProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x7 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x7.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush()
  return false, x7

}

type registerProcessorVersion struct {
  handler Register
}

func (p *registerProcessorVersion) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := RegisterVersionArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("version", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := RegisterVersionResult{}
var retval string
  var err2 error
  if retval, err2 = p.handler.Version(); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing version: " + err2.Error())
    oprot.WriteMessageBegin("version", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("version", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type registerProcessorHeartbeat struct {
  handler Register
}

func (p *registerProcessorHeartbeat) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := RegisterHeartbeatArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("heartbeat", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := RegisterHeartbeatResult{}
var retval *Response
  var err2 error
  if retval, err2 = p.handler.Heartbeat(args.ServiceInfo); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing heartbeat: " + err2.Error())
    oprot.WriteMessageBegin("heartbeat", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("heartbeat", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

type RegisterVersionArgs struct {
}

func NewRegisterVersionArgs() *RegisterVersionArgs {
  return &RegisterVersionArgs{}
}

func (p *RegisterVersionArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *RegisterVersionArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("version_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *RegisterVersionArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("RegisterVersionArgs(%+v)", *p)
}

// Attributes:
//  - Success
type RegisterVersionResult struct {
  Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewRegisterVersionResult() *RegisterVersionResult {
  return &RegisterVersionResult{}
}

var RegisterVersionResult_Success_DEFAULT string
func (p *RegisterVersionResult) GetSuccess() string {
  if !p.IsSetSuccess() {
    return RegisterVersionResult_Success_DEFAULT
  }
return *p.Success
}
func (p *RegisterVersionResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *RegisterVersionResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *RegisterVersionResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *RegisterVersionResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("version_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *RegisterVersionResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteString(string(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *RegisterVersionResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("RegisterVersionResult(%+v)", *p)
}

// Attributes:
//  - ServiceInfo
type RegisterHeartbeatArgs struct {
  ServiceInfo *ServiceInfo `thrift:"serviceInfo,1" db:"serviceInfo" json:"serviceInfo"`
}

func NewRegisterHeartbeatArgs() *RegisterHeartbeatArgs {
  return &RegisterHeartbeatArgs{}
}

var RegisterHeartbeatArgs_ServiceInfo_DEFAULT *ServiceInfo
func (p *RegisterHeartbeatArgs) GetServiceInfo() *ServiceInfo {
  if !p.IsSetServiceInfo() {
    return RegisterHeartbeatArgs_ServiceInfo_DEFAULT
  }
return p.ServiceInfo
}
func (p *RegisterHeartbeatArgs) IsSetServiceInfo() bool {
  return p.ServiceInfo != nil
}

func (p *RegisterHeartbeatArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *RegisterHeartbeatArgs)  ReadField1(iprot thrift.TProtocol) error {
  p.ServiceInfo = &ServiceInfo{}
  if err := p.ServiceInfo.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.ServiceInfo), err)
  }
  return nil
}

func (p *RegisterHeartbeatArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("heartbeat_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *RegisterHeartbeatArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("serviceInfo", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:serviceInfo: ", p), err) }
  if err := p.ServiceInfo.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.ServiceInfo), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:serviceInfo: ", p), err) }
  return err
}

func (p *RegisterHeartbeatArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("RegisterHeartbeatArgs(%+v)", *p)
}

// Attributes:
//  - Success
type RegisterHeartbeatResult struct {
  Success *Response `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewRegisterHeartbeatResult() *RegisterHeartbeatResult {
  return &RegisterHeartbeatResult{}
}

var RegisterHeartbeatResult_Success_DEFAULT *Response
func (p *RegisterHeartbeatResult) GetSuccess() *Response {
  if !p.IsSetSuccess() {
    return RegisterHeartbeatResult_Success_DEFAULT
  }
return p.Success
}
func (p *RegisterHeartbeatResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *RegisterHeartbeatResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *RegisterHeartbeatResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &Response{}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *RegisterHeartbeatResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("heartbeat_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *RegisterHeartbeatResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *RegisterHeartbeatResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("RegisterHeartbeatResult(%+v)", *p)
}

