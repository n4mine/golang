# golang

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [golang toolkits(一些常用的代码片段)](#golang-toolkits%E4%B8%80%E4%BA%9B%E5%B8%B8%E7%94%A8%E7%9A%84%E4%BB%A3%E7%A0%81%E7%89%87%E6%AE%B5)
- [golang toolkits(开源的一些库/工具)](#golang-toolkits%E5%BC%80%E6%BA%90%E7%9A%84%E4%B8%80%E4%BA%9B%E5%BA%93%E5%B7%A5%E5%85%B7)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->


## golang toolkits(一些常用的代码片段)
* [get_outbound_ipaddr](./get_outbound_ipaddr) - 获取本机对外IP
* [remove_dup](./remove_dup) - 返回指定slice的非重复版本
* [reverse_string](./reverse_string) - 反转一个字符串
* [split_by_n](./split_by_n) - 分割target, 每段size为n
    * 可用于分割请求等场景
    * 返回每个slice的start和end pos
* [union_slice](./union_slice) - 求2个slice的交集
* [split_int](./split_int) - 拆分一个int, 例如将12345，拆成`[]int{5,4,3,2,1}`的形式
* [get_jiffies](./get_jiffies) - 获取jiffies, 只适用于Linux
* [is_power_of_two](./is_power_of_two) - 判断给定数字是否是2的N次方
* [inverted_index](./inverted_index) - 一个简单的倒排索引DEMO
* [queue](./queue) - 线程安全的queue

## golang toolkits(开源的一些库/工具)
* [go-metrics](https://github.com/rcrowley/go-metrics) - Go port of Coda Hale's Metrics library
* [Guru](http://golang.org/s/using-guru) - golang 代码导航工具
* [go-tools](https://github.com/dominikh/go-tools#tools) - golang 代码检查工具
* [mapstructure](https://github.com/mitchellh/mapstructure) - Go library for decoding generic map values into native Go structures.
* [sarama](https://github.com/Shopify/sarama) - Sarama is a Go library for Apache Kafka 0.8, 0.9, and 0.10.
* [sarama-cluster](https://github.com/bsm/sarama-cluster) - Cluster extensions for Sarama, the Go client library for Apache Kafka 0.9 (and later).
* [glob](https://github.com/gobwas/glob) - Go glob
* [sqlx](https://github.com/jmoiron/sqlx) - general purpose extensions to golang's database/sql
* [consistent](https://github.com/stathat/consistent) - Consistent hash package for Go.
* [testify](https://github.com/stretchr/testify/) - A toolkit with common assertions and mocks that plays nicely with the standard library
* [xorm](https://github.com/go-xorm/xorm) - Simple and Powerful ORM for Go, support mysql,postgres,tidb,sqlite3,mssql,oracle
* [Leaktest](https://github.com/fortytw2/leaktest) - goroutine leak detector
* [go-torch](https://github.com/uber/go-torch) - Stochastic flame graph profiler for Go programs
