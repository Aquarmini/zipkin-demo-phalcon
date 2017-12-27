namespace php Xin.Thrift.MicroService
namespace go vendor.service

exception ThriftException {
  1: i32 code,
  2: string message
}

service App {
    // 返回项目版本号
    string version() throws (1:ThriftException ex)

    // 测试异常抛出
    string testException() throws(1:ThriftException ex)

    // 测试内存自增
    string num() throws (1:ThriftException ex)

    // 判断redis db切换操作
    void redis() throws (1:ThriftException ex)
}