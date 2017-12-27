<?php
// +----------------------------------------------------------------------
// | ZipkinHandler.php [ WE CAN DO IT JUST THINK IT ]
// +----------------------------------------------------------------------
// | Copyright (c) 2016-2017 limingxinleo All rights reserved.
// +----------------------------------------------------------------------
// | Author: limx <715557344@qq.com> <https://github.com/limingxinleo>
// +----------------------------------------------------------------------
namespace App\Thrift\Services;

use App\Thrift\Clients\ZipkinClient;
use Xin\Thrift\ZipkinService\ZipkinIf;
use Xin\Thrift\ZipkinService\ZipkinException;
use Xin\Thrift\ZipkinService\Options;
use Zipkin\Tracer;

class ZipkinHandler extends Handler implements ZipkinIf
{
    public function version(Options $options)
    {
        return $this->config->version;
    }

    public function test(Options $options)
    {
        $client = ZipkinClient::getInstance();
        $version = $client->version($options);
        return $version;
    }
}