<?php
// +----------------------------------------------------------------------
// | AppHandler.php [ WE CAN DO IT JUST THINK IT ]
// +----------------------------------------------------------------------
// | Copyright (c) 2016-2017 limingxinleo All rights reserved.
// +----------------------------------------------------------------------
// | Author: limx <715557344@qq.com> <https://github.com/limingxinleo>
// +----------------------------------------------------------------------
namespace App\Thrift\Services;

use App\Biz\Test;
use App\Utils\Redis;
use Xin\Thrift\MicroService\AppIf;
use Xin\Thrift\MicroService\ThriftException;

class AppHandler extends Handler implements AppIf
{
    /**
     * @desc   返回项目版本号
     * @author limx
     * @return mixed
     * @throws ThriftException
     */
    public function version()
    {
        return $this->config->version;
    }

    /**
     * @desc   测试异常抛出
     * @author limx
     * @throws ThriftException
     */
    public function testException()
    {
        throw new ThriftException([
            'code' => '400',
            'message' => '异常测试'
        ]);
    }

    public function num()
    {
        return Test::getInstance()->num();
    }

    public function num2()
    {
        try {
            $key = uniqid();
            $client = Test::getInstance($key);
            $num = $client->num();
        } finally {
            dump($client->instanceCount());
            $client->flushInstance();
        }

        return $num;
    }

    public function redis()
    {
        Redis::set('1', '1');
        Redis::select(2);
    }
}
