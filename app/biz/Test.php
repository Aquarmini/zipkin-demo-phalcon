<?php
// +----------------------------------------------------------------------
// | Test.php [ WE CAN DO IT JUST THINK IT ]
// +----------------------------------------------------------------------
// | Copyright (c) 2016-2017 limingxinleo All rights reserved.
// +----------------------------------------------------------------------
// | Author: limx <715557344@qq.com> <https://github.com/limingxinleo>
// +----------------------------------------------------------------------
namespace App\Biz;

use Xin\Traits\Common\InstanceTrait;

class Test
{
    use InstanceTrait;

    public $num = 0;

    public function num()
    {
        return $this->num++;
    }

    public function flush()
    {

    }
}