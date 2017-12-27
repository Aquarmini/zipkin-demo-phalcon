<?php
// +----------------------------------------------------------------------
// | Console.php [ WE CAN DO IT JUST THINK IT ]
// +----------------------------------------------------------------------
// | Copyright (c) 2016-2017 limingxinleo All rights reserved.
// +----------------------------------------------------------------------
// | Author: limx <715557344@qq.com> <https://github.com/limingxinleo>
// +----------------------------------------------------------------------
namespace App\Core\Services;

use Phalcon\Config;
use Phalcon\DI\FactoryDefault;
use Zipkin\Endpoint;
use GuzzleHttp\Client;
use Zipkin\Annotation;
use Zipkin\Samplers\BinarySampler;
use Zipkin\TracingBuilder;
use Zipkin\Reporters\HttpLogging;

class Tracer implements ServiceProviderInterface
{
    public function register(FactoryDefault $di, Config $config)
    {
        $di->setShared('tracer', function () use ($di) {
            $endpoint = Endpoint::createFromGlobals();
            $client = new Client();
            // Logger to stdout
            $logger = new \App\Common\Logger();

            $reporter = new HttpLogging($client, $logger);
            $sampler = BinarySampler::createAsAlwaysSample();
            $tracing = TracingBuilder::create()
                ->havingLocalEndpoint($endpoint)
                ->havingSampler($sampler)
                ->havingReporter($reporter)
                ->build();

            return $tracing;
        });
    }

}