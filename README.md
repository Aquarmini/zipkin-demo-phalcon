# PHALCON基础开发框架

> 本项目以[limingxinleo/phalcon-project](https://github.com/limingxinleo/phalcon)为基础，进行简易封装。

[![Build Status](https://travis-ci.org/limingxinleo/phalcon-unit-test.svg?branch=master)](https://travis-ci.org/limingxinleo/phalcon-unit-test)
[![Total Downloads](https://poser.pugx.org/limingxinleo/phalcon-project/downloads)](https://packagist.org/packages/limingxinleo/phalcon-project)
[![Latest Stable Version](https://poser.pugx.org/limingxinleo/phalcon-project/v/stable)](https://packagist.org/packages/limingxinleo/phalcon-project)
[![Latest Unstable Version](https://poser.pugx.org/limingxinleo/phalcon-project/v/unstable)](https://packagist.org/packages/limingxinleo/phalcon-project)
[![License](https://poser.pugx.org/limingxinleo/phalcon-project/license)](https://packagist.org/packages/limingxinleo/phalcon-project)


[Phalcon 官网](https://docs.phalconphp.com/zh/latest/index.html)

[wiki](https://github.com/limingxinleo/simple-subcontrollers.phalcon/wiki)

### Zipkin测试
- 安装Zipkin服务
~~~
docker run --name zipkin -d -p 9411:9411 openzipkin/zipkin 
~~~

- 启动RPC服务
~~~
php run thrift:service
~~~

- 调用测试命令
~~~
php run zipkin:test@zipkin
~~~

### 封装版本
- [Thrift GO服务版本](https://github.com/limingxinleo/thrift-go-phalcon-project)
- [Phalcon快速开发框架](https://github.com/limingxinleo/biz-phalcon)
- [Phalcon基础开发框架](https://github.com/limingxinleo/basic-phalcon)


