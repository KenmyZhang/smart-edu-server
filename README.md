### 技术选型

框架: gin

日志处理: zap

orm框架: gorm

优雅重启HTTP服务: gracehttp

测试框架: ginkgo

包管理工具 go mod

### 接口风格

 RESTFUL

###  项目结构
- api 接口层，入参解析以及校验
- app 应用的初始化、启动、关闭
- biz 业务逻辑处理
- common 公共类
- conf 配置文件
- log  日志处理
- model 数据模型定义以及数据库操作
- route  路由注册
- main.go - 程序执行入口
- Dockerfile 构建镜像文件
- Makefile 提供编译、打包、测试等功能的脚本文件
- ginkgo 二进制文件，容器内执行测试用例的时候需要使用的命令

### 编译

    make build

### 打包

    make package

### 测试

#### 先创建一个新的Docker网络

    docker network create -d bridge my-net

#### 执行测试用例
创建mysql容器实例会占用3306，确保该端口未被其他应用使用
  
    make test

最后生成的测试报告junit.xml和覆盖率coverprofile.txt文件在对应package的目录下

make test包括下面三个步骤

*  移除容器,排除历史测试数据干扰

    make clean-docker

*  启动docker，创建一个mysql容器实例

    make start-docker

*  开始执行测试用例

   make test-server

### 测试用例覆盖率可视化

  go tool cover -html=coverprofile.txt -o coverprofile.html

  可以很清楚地看到测试用例覆盖的代码和未曾覆盖到的代码

### 管理依赖包工具(go mod)

Add missing and remove unused modules

    go mod tidy -v

Make vendored copy of dependencies

	go mod vendor -v


### 平滑升级

 kill -USR2  PID


### 日志级别控制

debug 级别最低，主要在开发过程中用来打印调试信息；

info  用于打印程序应该出现的正常状态信息， 便于追踪定位；

后三个，警告、错误、严重错误，这三者应该都在系统运行时检测到了一个不正常的状态。

warn  表明系统出现轻微的不合理但不影响运行和使用；

Error 表明出现了系统错误和异常，无法正常完成目标操作。

Fatal 相当严重，可以肯定这种错误已经无法修复，并且如果系统继续运行下去的话后果严重。

