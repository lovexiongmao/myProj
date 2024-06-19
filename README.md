## 编程风格

https://gocn.github.io/styleguide/docs/03-decisions/#%E5%91%BD%E5%90%8Dnaming

## 项目骨架说明

apiservice api server

clients 用于存放接受用户请求的地方，是用户请求的入口

conf 用于存放配置相关的代码

dao 用于存放和数据库交互的代码

dig_container 用于存放 dig 依赖注入的地方

middleware 用于存放第三方插件

models 用于存放一些数据模型文件的地方

sdk 用于存放一些 sdk 工具

services 用于存储具体业务逻辑实现的地方

static 用于存储静态文件的地方

template 用于存储模板文件的地方

utils 用于存储一些常用的工具的地方

## 功能补全

1. 日志 done
2. dig 依赖注入及管理 done
3. 实现对一个用户数据的增删改查 done
4. 鉴权中间件
5. 读取配置，并且连接数据库 done
6. 通过参数来灵活的确定是否要迁移模型数据=>mysql done
7. 编译脚本本地测试
