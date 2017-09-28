# golang

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
