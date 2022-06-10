# 基础议题

## 条款1:仔细区别pointers和references

> 何时考虑用ponit

1. 当你需要考虑不指向任何对象的可能性时
或者在不同时间指向不同的对象时,可以采用ponit.

> 何时采用reference

1. 确定总要代表某个对象,并且一旦代表了该对象就不能改变了.
2. 某些操作符必须使用reference,例如operator[].

## 条款2:最好使用c++转型操作符
> 旧式的类型转换不安全,且难以识别,grep工具也不好区分.

+ C++的四个转型操作符
static_cast|const_cast|dynamic_cast|reinterpret_cast|

> static_cast

```
static支持与C旧式转型相同的意义,以及相同的限制.
```
> const_cast

```
const_cast用来改变表达式中的常量性或易变性,这项意愿由编译器贯彻执行.
```

```c++
class Widget {...};
class SpecialWidget : public Widget {...};
void update(SpecialWidget *psw);
SpecialWidget sw;
const SpecialWidget& csw = sw;

update(&csw);     //error,不能将const对象传给需要非const对象的函数
update(const_cast<SpecialWidget*>(&csw)); //可以.将&csw的常量性质去除了.

```

> dynamic_cast

```
用来执行集成体系中"安全的向下转型或者跨系转型动作.(什么是跨系转型？)
```

+ 当转型失败时，point--->nullptr

+ 当转型失败时,reference--->抛出一个exception

> dynamic_cast适用于有虚函数的类中,无法在缺乏虚函数的类型身上,也不能改变类型的常量性.

条款24
```


```



>reinterpret_cast

```
此类型转换与编译平台相关,不具有移植性.
```

1. 可以对函数指针类型进行转换,但是由于不具有移植性,所以尽量不要使用.


## 条款3:绝对不要以polymorphically方式处理数组








## 条款4:非必要不提供default constructor