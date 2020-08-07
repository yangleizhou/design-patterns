创建模式
----
- ## `本质`
    - 创建对象的模式:对类的实例化进行抽象
- ## `特点`
    - 1.封装了具体类的信息
    - 2.隐藏了类的实例化过程

- [单例模式](#1)
- [简单工厂模式](#2)
- [工厂方法模式](#3)
- [抽象工厂模式](#4)
- [原型模式](#5)
- [建造者模式](#6)
- [对象池模式](#7)
- [功能选项模式](#8)
- [工人队列模式](#9)

- ## <i id="1"></i>`Singleton 单例模式`  
- [Singleton](https://github.com/yangleizhou/design-patterns/tree/master/03-behavioral/01-singleton.go)        
    - 单例模式是最简单的设计模式之一, 保证一个类仅有一个实例, 并提供一个全局的访问接口
    - 缺点：1.单例类的职责过重，里面的代码可能会过于复杂，在一定程度上违背了“单一职责原则”。
           2.如果实例化的对象长时间不被利用，会被系统认为是垃圾而被回收，这将导致对象状态的丢失。
    - 实现方式: 
            a.初始化单例类时即创建单例
                a.1 饿汗式
                a.2 枚举类型
            b.按需、延迟创建单例
                b.1 懒汉式
                    b.1.1 基础实现
                    b.1.2 同步锁
                    b.1.3 双重检验锁
                b.2 静态内部类实现
- ## <i id="2"></i>`Simple Factory 简单工厂模式`  
- [Simple Factory](https://github.com/yangleizhou/design-patterns/tree/master/03-behavioral/02-simple_factory.go)    
    -  又称为静态工厂方法(Static Factory Method)模式，它属于类创建型模式。在简单工厂模式中，可以根据参数的不同返回不同类的实例。  
       简单工厂模式专门定义一个类来负责创建其他类的实例，被创建的实例通常都具有共同的父类。
    - 简单工厂模式包含如下角色：
        - Factory：工厂角色
        - Product：抽象产品角色
        - ConcreteProduct：具体产品角色
    - 说明:1.外界通过调用工厂类的静态方法，传入不同参数从而创建不同具体产品类的实例
    - 缺点：1.违背"开放封闭原则"，一旦添加新产品就不得不修改工厂类的逻辑，这样就会造成工厂逻辑过于复杂。
            2.工厂类集中了所有实例（产品）的创建逻辑，一旦这个工厂不能正常工作，整个系统都会受到影响。
            3.简单工厂模式由于使用了静态工厂方法，静态方法不能被继承和重写，会造成工厂角色无法形成基于继承的等级结构。
- ## <i id="3"></i>`Factory Method 工厂方法模式`  
- [Factory Method](https://github.com/yangleizhou/design-patterns/tree/master/03-behavioral/03-factory_method.go)     
    - 使一个类的实例化(具体产品创建)延迟到其子类(具体工厂)中完成, 定义一个用于创建对象的接口, 让子类决定将哪一个类实例化。
    - 工厂方法模式包含如下角色：
        - Factory: 抽象工厂
        - ConcreteFactory:具体工厂
        - Product: 抽象产品
        - ConcreteProduct:具体产品
    - 说明：1.创建抽象工厂类，定义具体工厂的公共接口
            2.创建抽象产品类 ，定义具体产品的公共接口
            3.创建具体产品类（继承抽象产品类） & 定义生产的具体产品
            4.创建具体工厂类（继承抽象工厂类），定义创建对应具体产品实例的方法
            5.外界通过调用具体工厂类的方法，从而创建不同具体产品类的实例
    - 缺点：1.一个具体工厂只能创建一种具体产品
- ## <i id="4"></i>`Abstract Factory 抽象工厂模式`  
- [Abstract Factory](https://github.com/yangleizhou/design-patterns/tree/master/03-behavioral/04-abstract_factory.go)       
    - 提供一个创建一系列相关或相互依赖对象的接口, 而无需指定它们具体的类
    - 备注概念：
        1.产品等级结构 ：产品等级结构即产品的继承结构，如一个抽象类是电视机，其子类有海尔电视机、海信电视机、TCL电视机，则抽象电视机与具体品牌的电视机之间构成了一个产品等级结构，抽象电视机是父类，而具体品牌的电视机是其子类。  
        2.产品族 ：在抽象工厂模式中，产品族是指由同一个工厂生产的，位于不同产品等级结构中的一组产品，如海尔电器工厂生产的海尔电视机、海尔电冰箱，海尔电视机位于电视机产品等级结构中，海尔电冰箱位于电冰箱产品等级结构中。
    - 抽象工厂模式包含如下角色：
        - AbstractProduct:抽象产品族,描述抽象产品的公共接口,抽象产品父类
        - Product:抽象产品,描述具体产品的公共接口,具体产品的父类
        - ConcreteProduct:具体产品,描述生产的具体产品,工厂类创建的目标类
        - Factory:抽象工厂,描述具体工厂的公共接口
        - ConcreteFactory:具体工厂,实现FactoryMethod工厂方法创建产品的实例
    - 说明：1.创建抽象工厂类，定义具体工厂的公共接口
            2.创建抽象产品族类 ，定义抽象产品的公共接口
            3.创建抽象产品类 （继承抽象产品族类），定义具体产品的公共接口
            4.创建具体产品类（继承抽象产品类） & 定义生产的具体产品
            5.创建具体工厂类（继承抽象工厂类），定义创建对应具体产品实例的方法
            6.客户端通过实例化具体的工厂类，并调用其创建不同目标产品的方法创建不同具体产品类的实例
    - 缺点：抽象工厂模式很难支持新种类产品的变化。
            这是因为抽象工厂接口中已经确定了可以被创建的产品集合，如果需要添加新产品，此时就必须去修改抽象工厂的接口，这样就涉及到抽象工厂类的以及所有子类的改变，这样也就违背了“开发——封闭”原则。
- ## <i id="5"></i>`Prototype 原型模式`  
- [Prototype](https://github.com/yangleizhou/design-patterns/tree/master/03-behavioral/05-prototype.go)  
    - 用一个已经创建的实例作为原型，通过复制该原型对象来创建一个和原型相同或相似的新对象，同时又能保证性能
    - 原型模式包含如下角色：
        - 抽象原型类,规定了具体原型对象必须实现的接口
        - 具体原型类,实现抽象原型类的 clone() 方法，它是可被复制的对象
        - 访问类：使用具体原型类中的 clone() 方法来复制新的对象
    缺点：1.配备克隆方法需要对类的功能进行通盘考虑，这对于全新的类不是很难，但对于已有的类不一定很容易，特别当一个类引用不支持串行化的间接对象，或者引用含有循环  
           结构的时候。对已有的类进行改造时，不一定是件容易的事，必须修改其源代码，违背了“开闭原则”。
         2.必须实现 Cloneable 接口。
         3.在实现深克隆时需要编写较为复杂的代码。
- ## <i id="6"></i>`Builder 建造者模式`  
- [Builder](https://github.com/yangleizhou/design-patterns/tree/master/03-behavioral/06-builder.go)      
    - 将一个复杂对象的构建与它的表示分离, 使得同样的构建过程可以创建不同的表示。
    - 建造者模式包含如下角色：
        - Builder：抽象建造者
        - ConcreteBuilder：具体建造者
        - Director：指挥者
        - Product：产品角色
    - 说明：隐藏创建对象的建造过程 & 细节，使得用户在不知对象的建造过程 & 细节的情况下，就可直接创建复杂的对象。
    - 缺点：1.建造者模式所创建的产品一般具有较多的共同点，其组成部分相似；如果产品之间的差异性很大，则不适合使用建造者模式，因此其使用范围受到一定的  
               限制。
            2.如果产品的内部变化复杂，可能会导致需要定义很多具体建造者类来实现这种变化，导致系统变得很庞大。
- ## <i id="7"></i>`Object Pool 对象池模式`  
- [Object Pool](https://github.com/yangleizhou/design-patterns/tree/master/03-behavioral/07-object_pool.go)      
    - 根据需求将预测的对象保存到channel中， 用于对象的生成成本大于维持成本
    - 说明：1.对象池模式(The Object Pool Pattern)是单例模式的一个变种，它提供了获取一系列相同对象实例的入口。当你需要对象来代表一组可替代资源的时候  
            就变的很有用，每个对象每次可以被一个组件使用
            2.对象池模式经常用在频繁创建、销毁对象(并且对象创建、销毁开销很大)的场景，比如数据库连接池、线程池、任务队列池等
    - 缺点：1.并发环境中， 多个线程可能(同时)需要获取池中对象， 进而需要在堆数据结构上进行同步或者因为锁竞争而产生阻塞， 这种开销要比创建销毁对象的  
            开销高数百倍
            2.由于池中对象的数量有限， 势必成为一个可伸缩性瓶颈
    - channel注意事项:
            1.不要在读取端关闭 channel ，因为写入端无法知道 channel 是否已经关闭，往已关闭的 channel 写数据会 panic 
            2.有多个写入端时，不要在写入端关闭 channle ，因为其他写入端无法知道 channel 是否已经关闭，关闭已经关闭的 channel 会发生 panic 
            3.如果只有一个写入端，可以在这个写入端放心关闭 channel 。
              关闭 channel 粗暴一点的做法是随意关闭，如果产生了 panic 就用 recover 避免进程挂掉。稍好一点的方案是使用标准库的 sync 包来做关闭 channel 时的协程同步，不过使用起来也稍微复杂些
    - channel报painc:
            1.panic: close of nil channel 关闭nil的channel
            2.panic: close of closed channel 关闭已关闭的channel
            3.panic: send on closed channel 写已关闭的channel 
- ## <i id="8"></i>`Options 功能选项模式`  
- [Options](https://github.com/yangleizhou/design-patterns/tree/master/03-behavioral/08-options.go)     
    - 扩展的构造函数和其他公共API中的可选参数
    - 关键代码
        - 定义一个Option接口，拥有一个设置参数的函数
        - 定义一个optFunc实现Option接口
        - 构造结构体时，接收可变类型的Option
        - 遍历options，调用option中的设置参数方法
- ## <i id="9"></i>`Worker 工人队列模式`  
- [Worker](https://github.com/yangleizhou/design-patterns/tree/master/03-behavioral/09-worker.go)   
    - 高性能处理任务

    
