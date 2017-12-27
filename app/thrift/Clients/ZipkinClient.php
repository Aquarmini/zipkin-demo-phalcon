<?php

namespace App\Thrift\Clients;

use App\Thrift\Client;
use Xin\Thrift\ZipkinService\Options;
use Xin\Thrift\ZipkinService\ZipkinClient as ZipkinServiceClient;
use Zipkin\Propagation\RequestHeaders;
use Zipkin\Propagation\TraceContext;
use Zipkin\Tracer;
use Zipkin\Tracing;

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
        $options = end($arguments);
        /** @var Tracing $tracing */
        $tracing = di('tracer');
        $tracer = $tracing->getTracer();

        $d = debug_backtrace()[1];
        $serviceName = env('APP_NAME');
        $spanName = $serviceName . ':' . $d['class'] . '@' . $d['function'];

        if (!$options instanceof Options) {
            $trace = $tracer->newTrace();
            $trace->setName($spanName);
            $trace->start();
            $context = $trace->getContext();
            $options = new Options();
            $options->span = $spanName;
            $options->traceId = $context->getTraceId();
            $options->parentSpanId = $context->getParentId();
            $options->spanId = $context->getSpanId();
            $options->sampled = $context->isSampled();
            $arguments[] = $options;
        } else {
            $context = TraceContext::create(
                $options->traceId,
                $options->spanId,
                $options->parentSpanId,
                $options->sampled
            );
            $trace = $tracer->newChild($context);
            $trace->setName($spanName);
            $trace->start();
        }

        try {
            $result = $this->client->$name(...$arguments);
        } finally {
            $trace->finish();
            $tracer->flush();
        }

        return $result;
    }
}

