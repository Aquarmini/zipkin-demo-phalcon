<?php

namespace App\Thrift\Clients;

use App\Thrift\Client;
use Xin\Thrift\ZipkinService\Options;
use Xin\Thrift\ZipkinService\ZipkinClient as ZipkinServiceClient;

class ZipkinClient extends Client
{
    protected $host = '127.0.0.1';

    protected $port = '10086';

    protected $service = 'zipkin';

    protected $clientName = ZipkinServiceClient::class;

    protected $recvTimeoutMilliseconds = 50;

    protected $sendTimeoutMilliseconds;

    /**
     * @desc
     * @author limx
     * @param array $config
     * @return ZipkinServiceClient
     */
    public static function getInstance($config = [])
    {
        return parent::getInstance($config);
    }

    public function __call($name, $arguments)
    {
        $options = new Options();
        $options->span = $name;
        $arguments[] = $options;
        return $this->client->$name(...$arguments);
    }
}

