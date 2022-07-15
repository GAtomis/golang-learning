<!--
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-06-28 15:00:32
 * @LastEditTime: 2022-07-15 17:35:42
 * @LastEditors: Gavin
-->
# golang-learning
golang学习笔记

## 学习GO 

 * build:编译包和依赖
* clean:移除对象文件
* doc:显示包或者符号文档
* env: 打印go的环境信息
* bug:启动错误报告
* get:下载包
* install:安装源码包

### 使用go get
初始化
```
//初始化
go mod init started
```
安装remote
```
//初始化
go get github.com/go-sql-driver/mysql
```
### 创建一个项目
1. 新建项目文件夹
2. 初始化项目
3. 建立/user/user.go 



### 字符串
#### 字符串切片
```
n:=3//索引
m:=5
var s string ="hellow goto"
fmt.println(s[n])//打印索引3的原始字符
fmt.println(s[n:])//后续n以后的片段
fmt.println(s[n:m])//n-m的片段
fmt.println(s[:m])//从头到m的片段
fmt.Printf("s(len): %v\n", len(s)) //获取长度
```
#### 转义字符串
* /n  换行
* /*  转义
#### 输出占位
* %v 输出值
* %#v 包全路径
* %T 输出类型
* %t 布尔类型输出
* %b 二进制输出
* %c asckII 码

#### 运用strings包的方法
具体请查看包api

### golang格式化输出

####

### 流程控制
具体看包内实现
* array
* slice
* map
* for
* if eles

### 函数实现
* 高阶函数
* 定义函数类型和返回值
* 函数闭包
* 函数递归
* defer 
* init函数

### golang 指针
每个类型声明变量都有个内存地址,这个地址就叫一个指针

| 名称  |  描述    | 字符|
| ----- | ------ |-----|
| 指针地址  | 代表变量所在内存的地址 |&|
|指针类型  | 代表所取指针地址的指针类型 |*int,*string....|
|指针取值  | 代表通过地址去取原值 |*|

### 结构体
go中没有类这个概念所以取而代之的是结构体
如何定义一个结构体
```
type Person struct{
  id,age int
  name string
 
  sex bool
}
//具体实现请看demo

```
* 结构体声明
* 结构体指针
* 通过new关键字创造结构体指针
* 结构体嵌套

### golang方法
go中的方法作为函数是特殊的存在,于struct绑定 相当于类的方法

### interface
接口的实现
### 协程
通过go关键字去创建一个协程
### 并发编程-channel
通道可以跨越协程去传输数据,可实现执行同步
使用方法
```
//声明一个通道
var [Channel_name] chan [Type]
//函数创建
[Channel_name]:=make(chan [Type])
//send 压入数据
Mychannel <- element
//receice 弹出数据
element <- Mychannel

```

### 并发编程-waitGroup
实行执行栈去队列同步执行
>场景=>主协程结束后,子协程也会结束

### 并发编程-Mutex互斥锁实现同步
通过携程上锁从而保持无其他携程对执行上下文产生副作用
