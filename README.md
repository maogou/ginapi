### GinApi项目目录说明

- app 主要存放控制器-model-模板-service等
    - handler 相当于控制器
    - middleware 中间件
    - model 模型
    - service 共用服务
    - dao 模型的抽象层
- config 配置文件
- routes 路由
- docs api的接口文档地址
- global 全局的变量
- pkg 项目用的包
    - 错误码标准化 errcode
    - 配置管理
    - 数据库连接
- storage 日志文件存放的目录
- public 静态资源文件

#### GinApi要达到的预期目标

- 错误码标准化
    - 返回正确响应的结果集
    - 返回错误响应的错误码和消息体,提示客户具体是什么错误
- 配置管理
- 数据库连接管理
- 访问日志的统一化
    - 用于记录请求和响应的上下文和请求过程中的debug
- 响应的标准化
