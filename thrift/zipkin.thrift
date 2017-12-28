namespace php Xin.Thrift.ZipkinService
namespace go vendor.service

exception ZipkinException {
    1: i32 code,
    2: string message
}

struct Options {
    1: string traceId,
    2: string parentSpanId,
    3: string spanId,
    4: string sampled,
}

service Zipkin {
    // 返回项目版本号
    string version(1: Options options) throws (1:ZipkinException ex)

    // 测试链路监控
    string test(1: Options options) throws (1:ZipkinException ex)

    // 测试调用三次
    string test3(1: string name, 2: Options options) throws (1:ZipkinException ex)
}