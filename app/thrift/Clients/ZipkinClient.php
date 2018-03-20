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

    protected $recvTimeoutMilliseconds = 1000;

    protected $sendTimeoutMilliseconds = 1000;

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

        if (!$options instanceof Options) {
            // 首次调用
            $d = debug_backtrace()[1];
            $spanName = $d['class'] . '@' . $d['function'];

            $trace1 = $tracer->newTrace();
            $trace1->setName($spanName);
            $trace1->start();
            $context = $trace1->getContext();
            $options = new Options();
            $options->traceId = $context->getTraceId();
            $options->parentSpanId = $context->getParentId();
            $options->spanId = $context->getSpanId();
            $options->sampled = $context->isSampled();
            $arguments[] = $options;
        }

        $spanName = get_called_class() . '@' . $name;
        $context = TraceContext::create(
            $options->traceId,
            $options->spanId,
            $options->parentSpanId,
            $options->sampled
        );
        $trace2 = $tracer->newChild($context);
        $trace2->setName($spanName);
        $trace2->start();
        $context = $trace2->getContext();
        $options = array_pop($arguments);
        $options->traceId = $context->getTraceId();
        $options->parentSpanId = $context->getParentId();
        $options->spanId = $context->getSpanId();
        $options->sampled = $context->isSampled();
        $arguments[] = $options;

        try {
            $result = $this->client->$name(...$arguments);
        } finally {
            if (isset($trace2)) {
                $trace2->finish();
            }
            if (isset($trace1)) {
                $trace1->finish();
            }
            $tracer->flush();
        }

        return $result;
    }
}
