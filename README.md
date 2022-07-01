<!--
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-06-28 15:00:32
 * @LastEditTime: 2022-07-01 15:40:08
 * @LastEditors: Gavin
-->
# golang-learning
为了给我自己的Artemis后台搭一个 适合后台

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