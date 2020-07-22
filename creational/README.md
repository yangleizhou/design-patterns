创建模式
---- 
- `Simple Factory 简单工厂模式`
    - 又称为静态工厂方法(Static Factory Method)模式，它属于类创建型模式。在简单工厂模式中，可以根据参数的不同返回不同类的实例。简单工厂模式专门定义一个类来负责创建其他类的实例，被创建的实例通常都具有共同的父类。
    - 简单工厂模式包含如下角色：
        - Factory：工厂角色
        - Product：抽象产品角色
        - ConcreteProduct：具体产品角色
- `Builder 建造者模式`
    - 将一个复杂对象的构建与它的表示分离, 使得同样的构建过程可以创建不同的表示
    - 建造者模式包含如下角色：
        - Builder：抽象建造者
        - ConcreteBuilder：具体建造者
        - Director：指挥者
        - Product：产品角色

- 