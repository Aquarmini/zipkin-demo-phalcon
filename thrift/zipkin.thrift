namespace php Xin.Thrift.ZipkinService
namespace go vendor.service

exception ZipkinException {
    1: i32 code,
    2: string message
}

struct Options {
    1: string span,
}

service Zipkin {
    // 返回项目版本号
    string version(1: Options options) throws (1:ZipkinException ex)
}