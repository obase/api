# obase-api

## 概述介绍

底层集成gin, grpc-go提供Http, WebSocket, Grpc多种访问渠道. 开发者定义protobuf文件, 调用"protoc --go_out=plugins=grpc+aipx:.."
自动生成service的http post, websocket, grpc访问接口, 届时实现XXXServer业务逻辑即可, 大幅提升开发效率!

## 使用步骤

服务接口开发4个步骤:

### 第1步: 使用protov3语法定义服务接口

### 第2步: 调用protoc命令生成服务代码

### 第3步: 实现XXXServer接口编写业务逻辑

### 第4步: 