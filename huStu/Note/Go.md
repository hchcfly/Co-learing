# Go

## Go配置

配置go项目GoPATH(Go1.1之前)

+ go mod配置

```shell
go mod init <projectName>
```

## Go特殊语法

+ 区分大小写

```
任何需要对外暴露的名字必须以大写字母开头,不需要对外暴露的则应该以小写字母开头(相当于private).
```

+ 包名称

```
保持packet的名字和目录保持一致,且包名应该为小写字母,不要使用下划线或者混合大小写.
packet dao
packet service
```

+ 文件命名

```
小写单词,使用下划线分隔各个单词.
customer_dao.go
```

+ 结构体命名

```Go
驼峰命名法,首字母根据访问控制大写(public)或小写(private).
struct:
type CustomerOrder struct {
    Name string
    Address string
}
```

+ 接口命名

```go
单个函数的结构名以'er'作为后缀.
type Reader interface {
	Read(p []byte) (n int,err error)
}
```

+ 变量命名

```
首字母根据访问控制大小或小写.
特有名词:bool类型,名称应该以Has,Is,Can,Allow开头
```

+ 常量命名

```go
常量均需全部大写字母组成,并使用下划线分词.
const APP_URL  = "https://www.baidu.com"

枚举变量需要先创建相应类型
```

+ 错误处理

```go
if err != nil {
    return
}
```



