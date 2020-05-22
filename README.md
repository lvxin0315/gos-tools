# gos-tools
golang写的日常工具


### 1.tail/listen_tp5_log.go 监听tp5日志中的err 并将 err 钉钉发消息


------------


### 2.tail/listen_nginx_log.go 监听nginx error日志中的内容并钉钉发消息


------------


### 3.mysql/db_info.go 监听指定mysql指定表的变化
使用mysql的information_schema.TABLES实现，为什么不用binlog？
- 已经用[canal](https://github.com/alibaba/canal "canal")做了一套，不过对资源消耗比较大
- 小团队大部分业务都是通过定时，开发这个就是配合这样的被动的方式
- 资源占用小，通过修改间隔时间，基本能满足日常业务使用

------------

