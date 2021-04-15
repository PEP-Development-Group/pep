这是某校物理实验平台选课系统（单机应用），中途因为某个原因不用了，特此开源出来供大家学习。

技术栈：Gin + Vue + ElementUI。（特此鸣谢 gin-vue-admin 框架，这个项目是基于这个魔改的结果）

这里介绍一下项目大致结构和部署方式，如果你觉得对你有帮助，给我们一个 :star: **star** 就是对我们以及开源精神最大的支持 :)，感谢！

如果你有任何问题，欢迎提 issue 告诉我们，或者直接发起 pr。

也欢迎通过 issue 与我们交流 :smile:

## Backend

### 目录结构：

```
.
|-- db_dump.sql			// 数据库dump
|-- Dockerfile
|-- Makefile
|-- api		
|-- config				// 配置项的结构体定义
|-- config.yaml			// 配置文件
|-- constant 			// 常量定义
|-- core				// 启动服务器
|-- docs				// swagger文档
|-- global				// 全局变量
|-- initialize			// gorm，redis，router，token-bucket等的初始化
|-- log					// 日志文件
|-- main.go
|-- middleware			// casbin，cors，jwt等中间件
|-- model				// 表模型定义
|-- resource
|   |-- rbac_model.conf	// rbac模型
|-- router				// 路由
|-- service				// 业务层
|-- test.xlsx			// 批量导入学生的示例文件
`-- utils				// 一些工具
```

### 运行

- 修改 server/config.yaml 文件中的相关配置
- `go run main.go`

### 注意事项

- server/config.yaml 文件中的数据库改为你本地的数据库信息，然后导入server 目录下的 pep.sql 生成测试用的用户和课程数据。

## Frontend

### 运行

- npm run serve

## TODO

- [ ] 单机到分布式系统，主从，冷热备份机制
- [ ] 代理服务器，负载均衡
- [ ] 分布式节点间通信，rpc
- [ ] 采用一种消息队列解耦消息业务
- [ ] 采用HTTP2.0推送？