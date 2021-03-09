<p align="center"><img src="https://github.com/maogou/ginapi/blob/master/docs/ginapi.gif" width="256px"/></p> 


### GinApi项目目录说明

- app 主要存放控制器-model-模板-service等
    - handler 相当于控制器
    - middleware 中间件
        - jwt中间件验证token
        - recovery中间恢复异常
        - access_log 请求上下文日志记录
    - model 模型 
        - 更多用法 https://gorm.io/zh_CN/ 
        - Model 指定运行DB操作模型实例,默认解析该结构体的名字为表名,格式大写驼峰转小写下划线驼峰,也可以使用TableName方法手动指定表名
        - Where 设置搜索条件,可以接受map,struct,string作为条件
        - Offset 偏移量
        - Limit 搜索的条数
        - Find 查找符合条件的所有记录[多条]
        - Updates 更新所选字段
        - Delete 删除数据
        - Count 统计总条数
    - service 共用服务
    - dao 模型的抽象层
- config 配置文件
- routes 路由
- docs api的接口文档地址
    - swagger 工具安装
        - go get -u github.com/swaggo/swag/cmd/swag
        - go get -u github.com/swaggo/gin-swagger
        - go get -u github.com/swaggo/files
        - go get -u github.com/alecthomas/template
    - 验证swag 是否安装成功
        - swag -v 
    - Swag 注解常用标识说明
          
          | 注解       |描述   | 
          | --------  | -----:  | 
          | @Summary  |摘要   |  
          | @Produce  |API可以产生的MIME类型列表,可以理解为响应的类型  | 
          | @Param    |参数格式,从左到右分别:参数名,入参类型,数据类型,是否必填和注释 |  
          | @Success  |成功响应,从左到右分别:参数名,入参类型,数据类型和注释 | 
          | @Failure  |失败响应,从左到右分别:参数名,入参类型,数据类型和注释 | 
          | @Router   |路由,从左到右分别:路由地址和HTTP方法 | 
     
    - 注册swag路由
        - 导入对应的包
            - _ "github.com/maogou/ginapi/docs"
            - 	swaggerFiles "github.com/swaggo/files"
            - 	ginSwagger "github.com/swaggo/gin-swagger" 
        - 配置路由
            - url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
            - router.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))    
     
    - 生成文档 swag init   
    
    - 访问文档 
        - http://xxxx:8000/swagger/index.html   
          
             
- global 全局的变量
- pkg 项目用的包
    - 错误码标准化 errcode
    - 配置管理 setting
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
