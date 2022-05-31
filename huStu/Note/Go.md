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
+ bool类型

```go
判断时只能用bool类型不能用非bool类型.与c/c++不同
例如:
a := 3
if a {
    //error
} else {

}
```

+ 整数类型

```go
没有doulbe只有float32 float64

```


## Go string
> Go字符串是一个任意字节的常量序列.[]byte字节数组

> golang里面的字符串都是不可变的,每次运算都会产生一个新的字符串,所以会产生很多临时的无用的字符串,不仅没有用,还会给gc带来额外的负担,所以性能比较差.

+ strings.Join()

```go
name := "Tom"
	age := "20"
	msg := strings.Join([]string{name,age},",")
	fmt.Printf("msg: %v\n",msg)
```

> join会先根据字符串数组的内容,计算出一个拼接之后的长度,然后申请对应长度大小的内存,一个一个字符串填入,在已有一个数组的情况下,这种效率很高。本来没有需要构造的话,代价也不小.

+ buffer.WriteString()

```go

```

> 可以当成可变字符使用,对内存的增长有优化,如果能预估字符串的长度,可以用buffer.Grow()接口来设置capacity.



## Go函数
> Go语言没有类的概念,所有的功能单元都定义在函数中,可以重复使用,函数的名称,参数列表和返回值类型构成了函数的签名.

+ Go函数特性

```go
1.包含的函数类型:普通函数,匿名函数,方法(定义在struct上的函数,struct为接收者).[与c++区别]
2.go语言不允许函数重载.[与c++区别]
3.go语言中的函数不能嵌套函数,但是可以嵌套匿名函数[与c++区别]
4.函数是一个值,可以讲函数赋值给变量,使得变量称为函数
5.函数可以作为参数传递给另一个函数
6.函数的返回值可以是一个函数
```

```go
go中经常会使用其中一个返回值作为函数是否执行成功,是否有错误信息的判断条件

当有多个返回值时,可以考虑放进slice(相同类型的返回值)或者map(不同类型的返回值)中

使用_可以丢弃返回值
```

+ 函数传参

```go
map,slice,interface,channek这些数据类型本来就是指针类型的,拷贝传的值是指针值,修改他们可能会修改外部数据结构的值

```


## Golang闭包

> 定义在一个函数内部的函数. 函数+引用环境 = 闭包



## Golang defer关键字

> go语言会将defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时,将延迟处理的语句按照defer的逆序进行执行(类似于栈).
+ defer的特征
```go
defer用于注册延迟调用

这些调用直到return前才会被执行,可以做资源管理,类似于c++中的RAII?

多个defer语句,按先进后出的方式执行.
```

+ defer的用途

```go
1.关闭文件句柄

2.锁资源的释放

3.数据库连接释放

```


## Golang init函数

> Golang init函数理解为verliog中的initial or c++中的ctor.  先于main函数之前执行,实现包级别的一些初始化动作.


+ init函数的主要特点

```go

1.init函数先于main函数自动执行,不能被其他函数调用

2.init函数没有输入参数,返回值

3.每个包可以有多个init函数

4.包的每个源文件也可以有多个init函数,这点比较特殊

5.同一个包的init执行顺序,golang没有明确定义,自己注意顺序.

6.不同包的init函数按照包导入的依赖关系决定执行顺序.
```

+ 初始化顺序

>初始化顺序:变量初始化->init()->main()



## Golang 指针

+ golang指针特性

```go
golang指针,不能做偏移和运算.

只有两个操作:&和*

```

+ golang指向数组的指针

```go

并不是像c++一样指向数组的首地址,而是指向数组的每一个元素

```

## Golang类型别名

```go
//类型别名(没有定义一个新的类型,类型还是原来的类型)
type NewType  = type

写项目时,有利于批量修改变量类型

//类型定义(相当于定义了一个全新的类型,与之前的类型不同)
type NewType type
```

## Golang结构体

```
可以使用new关键字创建结构体指针

```

## Golang方法

> Golang语言没有面向对象的特性,也没有类和对象的概念,但是可以使用结构体来模拟这些特性,声明一些方法,属于某个结构体.

> struct为方法的接受者.方法是一种特殊的函数.


## Golang接口

1. 一个类型可以实现多个接口.
2.多个类型可以实现同一个接口(多态)


## Go OCP原则

> 开闭原则open=close Principle.对扩展是开放的,对修改是关闭的.

```go
package main

import "fmt"

type Pet interface {
	eat()
	sleep()
}

type Dog struct {
	name string
}

type Cat struct {
	name string	
}

//Dog实现Pet接口
func (dog Dog)eat() {
	fmt.Println("Dog eat...")
}

func (dog Dog)sleep() {
	fmt.Println("Dog sleep...")
}

//Cat实现Pet接口
func (cat Cat)eat() {
	fmt.Println("Cat eat...")
}

func (cat Cat)sleep() {
	fmt.Println("Cat sleep...")
}

type Person struct {
	
}


func (person Person)care(pet Pet) {
	pet.eat()
	pet.sleep()
}

func main() {
	dog := Dog{
	}
	cat := Cat{
	}
	person := Person{
    }
    //同一个接口,不同的实现
	person.care(dog)
	person.care(cat)
}

```

## Go 继承写法

```go

type Animal struct {
    name string
    age int
}

type Dog struct {
    Aniaml
    color string
}

func main() {
    dog := Dog{
		Animal{name:"hc",age:2},
		"黑色",
	}
	dog.name()
	dog.age()
}
```


## Golang 构造函数ctor

> Golang中没有构造函数的概念,但是可以使用函数来模拟构造函数的功能.


## Golang包

> 理解为文件夹

## Golang包管理工具go module

> golang1.11新加的特性,用来管理模块中包的依赖关系,之前需要GOPATH.

+ go mod使用方法

```go
//初始化模块
go mod init <ProjectName>

//根据go.mod文件,进行依赖关系处理
go mod tidy

```

