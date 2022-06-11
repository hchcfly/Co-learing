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
> 多态和指针算术不能混用,多态和数组不能混用,数组几乎总是会涉及指针的算术运算.

> 由于数组的访问array[i]相当于*(array+i),也就是需要进行指针运算,而指针运算需要根据当前对象的sizeof()来移动,由于派生类和基类的sizeof不同,所以指针的移动不会到理想的位置.

```c++
class BST {...};
class BalancedBST : public BST {...};

void printBSTArray(ostream& s,const BST array[],int numElements)
{
    for(int i=0;i<numElements;i++)
    {
        s << array[i];
    }
}
BST BSTArray[10];
printBSTArray(cout ,BSTArray,10); //没问题
BalancedBST BalancedBSTArray[10]
//无法预期的结果. 指针运算的是sizeof(BST),而不是sizeof(BalancedBST)
printBSTArray(cout ,BalancedBSTArray,10);
```

+ 如何避免?(条款34)

```
具体类不要继承自另一个具体类? 
```

## 条款4:非必要不提供default constructor

+ 什么是default construct?

> 不给任何自变量就可以调用的构造函数.

+ 没有default ctor的类如何生成对象数组?

```c++
class EquipmentPiece {
public:
    EquipmentPiece(int IDNumber)
    {   
    }
};

//EquipmentPiece bestPieces[10]; //没有default ctor该语句出错.
    
    //EquipmentPiece *bestPieces = new EquipmentPiece[10];//没有default ctor该语句出错.
    
    //可以产生非heap数组
    // EquipmentPiece bestPiece[] = {
    //     EquipmentPiece(1),
    //     EquipmentPiece(2)
    // };

    //如何产生heap数组呢?
    //使用指针数组,然后对每一个指针 new EquipmentPiece(id)

```

# 操作符

> C++允许编译器在不同类型之间执行隐式转换.其中有两种函数允许编译器执行这样的转换,单自变量ctor和隐式类型转换操作符.

```c++
单自变量ctor:是指能够以单一自变量成功调用的ctor,该ctor可能声明拥有单一参数,也可能声明拥有多个参数,并且除了第一参数之外都有默认值.

class Name
{
public:
    Name(const string& s); //可以将string转为Name
};

class Rational
{
public:
    Rational(int numerator = 0,
             int denominator = 1); //可以将string转为Name
};

void fun(int a);
void fun2(inta ,int b=0);
```

```c++
隐式类型转换操作符:
class Rational
{
public:
    operator double() const; //将Rational转换为double
};

Rational r(1,2);
double d = 0.5 * 4;  //r在此会转换为doule类型然后进行运算.
```
+ 如何消除隐式类型转换操作符的副作用?

> 应该尽量避免使用类型转换操作符.例如标准库中,string并不含有隐式转换到C风格的str的,而是提供了一个显式的成员函数用来转化string.c_str().

+ 如何消除单自变量ctor完成的隐式转换?

> 关键词explicit,防止隐式类型转换

> 使用代理类来阻止隐式转换
```c++
template<class T>
class Array
{
public:

class ArraySize
{
public:
    ArraySize(int numElements):theSize(numElements) { }
    int size() const { return theSize; }

    private: 
        int theSize;
}

Array(ArraySize size); //传入代理类
};
```

## 条款6:区别incr/decr操作符的前置和后置形式

> 如何区分incr/decr操作符的前置和后置形式?

```c++
由于incr/decr操作符的前置和后置形式都不需要参数,这样不好区分,所以让后置式有一个int自变量,并且在调用时,编译器默默给int = 0

class UPInt
{
public:
    UPInt& operator++();  //前置
    const UPInt operator++(int); //后置
};
```
> 前置返回一个reference,后置返回一个const对象.(++++i合法,i++++不合法)

