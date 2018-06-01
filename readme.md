# Go语言多用户在线聊天系统

## 功能
* 注册
* 登录
* 发消息
* 获取用户列表

## 工具
### Protobuf数据传输协议
* 优点：编码速度快、数据体积小、规范统一
* 缺点：改动协议字段需要重新生成文件、数据没有可读性
* 安装protoc：

1. https://github.com/google/protobuf/releases/tag/v3.4.1 
2. win下找到win32包 解压并配置环境变量
3. linux下需要编译
4. 将可执行文件拷贝到$GOPATH bin目录（前提将$GOPATH/bin目录加入到系统环境变量）
* 安装protobuf库文件
```
go get github.com/golang/protobuf/proto
```
* 安装插件
```
go get github.com/golang/protobuf/protoc-gen-go
```
* 生成go文件
```
protoc --go_out=*.proto;
```
* test.proto
```
syntax = "proto3";  //指定版本，必须要写（proto3、proto2）  
package pb;
enum FOO 
{ 
    X = 0; 
};
//message是固定的。UserInfo是类名，可以随意指定，符合规范即可
message UserInfo{
    string message = 1;   //消息
    int32 length = 2;    //消息大小
    int32 cnt = 3;      //消息计数
}
```