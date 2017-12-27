namespace php Xin.Thrift.ZipkinService
namespace go vendor.service

exception ZipkinException {
    1: i32 code,
    2: string message
}

struct Options {
    1: string span,
    2: string traceId,
    3: string parentSpanId,
    4: string spanId,
    5: string sampled,
}

service Zipkin {
    // 返回项目版本号
    string version(1: Options options) throws (1:ZipkinException ex)

    // 测试链路监控
    string test(1: Options options) throws (1:ZipkinException ex)
}