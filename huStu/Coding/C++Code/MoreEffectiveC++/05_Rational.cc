#include<iostream>

class Rational
{
public:
    Rational(int numerator = 0,
             int denominator = 1):numerator_(numerator),denominator_(denominator)//可以将string转为Name
    {

    }
    
    operator double() const
    {
        std::cout << "i am doo" << std::endl;
        return numerator_*1.00 / denominator_;
    }

private:
    int numerator_;
    int denominator_;
};

int main()
{
    Rational r(1, 2);
    std::cout << r << std::endl; //调用operator double() 进行隐式转换
    double d = 0.5 * r; //调用operator double() 进行隐式转换

    return 0;
}