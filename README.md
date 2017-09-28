# golang

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [golang toolkits(内部)](#golang-toolkits%E5%86%85%E9%83%A8)
- [golang toolkits(外部)](#golang-toolkits%E5%A4%96%E9%83%A8)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->


## golang toolkits(内部)
* [get_outbound_ipaddr](./get_outbound_ipaddr) - 获取本机对外IP
* [remove_dup](./remove_dup) - 返回指定slice的非重复版本
* [reverse_string](./reverse_string) - 反转一个字符串
* [split_by_n](./split_by_n) - 分割target, 每段size为n
    * 可用于分割请求等场景
    * 返回每个slice的start和end pos
* [union_slice](./union_slice) - 求2个slice的交集
* [split_int](./split_int) - 拆分一个int, 例如将12345，拆成`[]int{5,4,3,2,1}`的形式
* [get_jiffies](./get_jiffies) - 获取jiffies, 只适用于Linux

## golang toolkits(外部)
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
* [Leaktest](https://github.com/fortytw2/leaktest) | goroutine leak detector
