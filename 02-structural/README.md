结构模式
----
- ## `本质`
    - 如何将类或对象按某种布局组成更大的结构,就像搭积木，可以通过简单积木的组合形成复杂的、功能更为强大的结构
- ## `特点`
    - 1.它分为类结构型模式和对象结构型模式，前者采用继承机制来组织接口和类，后者釆用组合或聚合来组合对象
    - 2.除了适配器模式分为类结构型模式和对象结构型模式两种，其他的全部属于对象结构型模式
- ## `Proxy 代理模式`
    - 代理模式用于延迟处理操作或者在进行实际操作前后对真实对象进行其它处理
    - 代理模式包含如下角色：
        - Subject: 抽象主题类,通过接口或抽象类声明真实主题和代理对象实现的业务方法
        - RealSubject:真实主题类,实现了抽象主题中的具体业务，是代理对象所代表的真实对象，是最终要引用的对象
        - Proxy: 代理类,提供了与真实主题相同的接口，其内部含有对真实主题的引用，它可以访问、控制或扩展真实主题的功能
    - 说明：给目标对象提供一个代理对象，并由代理对象控制对目标对象的引用
    - 缺点：1.在客户端和目标对象之间增加一个代理对象，会造成请求处理速度变慢
            2.增加了系统的复杂度
- ## `Adapter 适配器`
    - 将一个类的接口转换成客户希望的另外一个接口，使得原本由于接口不兼容而不能一起工作的那些类能一起工作
    - 适配器模式包含如下角色：
        - Target: 目标接口，当前系统业务所期待的接口，它可以是抽象类或接口
        - Adaptee: 适配者类,当前系统业务所期待的接口，它可以是抽象类或接口
        - Adapter: 适配器类,它是一个转换器，通过继承或引用适配者的对象，把适配者接口转换成目标接口，让客户按目标接口的格式访问适配者
    - 配器模式的形式分为：1.类的适配器模式 & 对象的适配器模式
                          2.前者类之间的耦合度比后者高，且要求程序员了解现有组件库中的相关组件的内部结构，所以应用相对较少些
    - 缺点：过多的使用适配器，会让系统非常零乱，不易整体进行把握
    - 类适配器模式：
        - 说明：类的适配器模式是把适配的类的API转换成为目标类的API
        - 缺点：高耦合，灵活性低(使用对象继承的方式，是静态的定义方式)
    - 对象适配器模式：
        - 说明：1.与类的适配器模式相同，对象的适配器模式也是把适配的类的API转换成为目标类的API
                2.与类的适配器模式不同的是，对象的适配器模式不是使用继承关系连接到Adaptee类，而是使用委派关系连接到Adaptee类
        - 缺点：使用复杂,需要引入对象实例,特别是需要重新定义Adaptee行为时需要重新定义Adaptee的子类，并将适配器组合适配
- ## `Bridge 桥接\柄体\接口`
    - 将抽象与实现分离，使它们可以独立变化。它是用组合关系代替继承关系来实现的，从而降低了抽象和实现这两个可变维度的耦合度
    - 桥接模式包含如下角色：
        - Abstraction:抽象化角色,抽象化给出的定义，并保存一个对实现化对象的引用
        - RefinedAbstraction:扩展抽象化角色，改变和修正父类对抽象化的定义
        - Implementor:实现化角色,这个角色给出实现化角色的接口，但不给出具体的实现。必须指出的是，这个接口不一定和抽象化角色的接口定义相同，  
                      实际上，这两个接口可以非常不一样。实现化角色应当只给出底层操作，而抽象化角色应当只给出基于底层操作的更高一层的操作
        - ConcreteImplementor：具体实现化角色,这个角色给出实现化角色接口的具体实现
    - 使用场景：1. 如果一个系统需要在抽象化和具体化之间增加更多的灵活性，避免在两个层次之间建立静态的继承关系，通过桥接模式可以使它们在抽象层  
                建立一个关联关系
                2.抽象部分”和“实现部分”可以以继承的方式独立扩展而互不影响，在程序运行时可以动态将一个抽象化子类的对象和一个实现化子类的对象进行组合，即系统需要对抽象化角色和实现化角色进行动态耦合
                3.一个类存在两个（或多个）独立变化的维度，且这两个（或多个）维度都需要独立进行扩展
                4.对于那些不希望使用继承或因为多层继承导致系统类的个数急剧增加的系统，桥接模式尤为适用
    - 说明：抽象化角色就像是一个水杯的手柄，而实现化角色和具体实现化角色就像是水杯的杯身。手柄控制杯身，这就是此模式别名“柄体”的来源
    - 缺点：1.桥接模式的使用会增加系统的理解与设计难度，由于关联关系建立在抽象层，要求开发者一开始就针对抽象层进行设计与编程
            2.桥接模式要求正确识别出系统中两个独立变化的维度，因此其使用范围具有一定的局限性，如何正确识别两个独立维度也需要一定的经验积累
- ## `Decorator 装饰`
    - 在不改变现有对象结构情况下(在不影响其他对象的情况下),动态地给一个对象增加一些职责，即增加其额外的功能
    - 装饰模式包含如下角色：
        - Component:抽象构件,定义一个对象接口或抽象类，以规范需要装饰的主体类，比如：房子
        - ConcreteComponent:具体构件,具体的需要装饰的主体(实际被动态地添加职责的对象)，比如，具体的谁家的房子
        - Decorator：抽象装饰者类,持有抽象主体的实例，并实现或继承抽象主体(实现了Component接口，用来扩展Component类的功能，但对于Component来说，  
                     是无需知道Decorator的存在的)。比如：抽象的房屋装饰品
        - ConcreteDecorator：具体装饰者类,具体的装饰品，需继承或实现抽象装饰(动态地添加职责的对象)。表示具体的装饰品，如沙发等
    - 说明：使用装饰模式来实现扩展比继承更加灵活，它以对客户透明的方式动态地给一个对象附加更多的责任。装饰模式可以在不需要创造更多子类的情况下，  
            将对象的功能加以扩展
    - 缺点：1.使用装饰模式进行系统设计时将产生很多小对象，这些对象的区别在于它们之间相互连接的方式有所不同，而不是它们的类或者属性值有所不同，  
            同时还将产生很多具体装饰类。这些装饰类和小对象的产生将增加系统的复杂度，加大学习与理解的难度
            2.这种比继承更加灵活机动的特性，也同时意味着装饰模式比继承更加易于出错，排错也很困难，对于多次装饰的对象，调试时寻找错误可能需要逐级排查，较为烦琐
- ## `Facade 外观`
    - 为多个复杂的子系统提供一个一致的接口，使这些子系统更加容易被访问
    - 外观模式包含如下角色：
        - Facade:外观角色,为多个子系统对外提供一个共同的接口
        - Sub System:子系统角色，实现系统的部分功能，客户可以通过外观角色访问它
        - Client:客户角色，通过一个外观角色访问各个子系统的功能
    - 说明：在层次化结构中，可以使用外观模式定义系统中每一层的入口
    - 优点：1.简化调用流程，客户端不需要知道子系统的实现，提高了安全性，符合"迪特米原则"
            2.提高灵活性，降低用户类和子系统类的耦合度，实现松耦合
            3.更好的划分访问层次
    - 缺点：如果新增子系统，需要修改外观类，违背了"开闭原则"
- ## `Flyweight 享元`
    - 运用共享技术来有效地支持大量细粒度对象的复用
    - 享元模式包含如下角色：
        - Flyweight:抽象享元类,抽象享元类声明一个接口，通过它可以接受并作用于外部状态
        - ConcreteFlyweight:具体享元类,具体享元类实现了抽象享元接口，其实例称为享元对象
        - UnsharedConcreteFlyweight：非共享具体享元类,非共享具体享元是不能被共享的抽象享元类的子类
        - FlyweightFactory：享元工厂类,享元工厂类用于创建并管理享元对象，它针对抽象享元类编程，将各种类型的具体享元对象存储在一个享元池中
    - 说明：1.在享元模式中可以共享的相同内容称为内部状态(IntrinsicState)，而那些需要外部环境来设置的不能共享的内容称为外部状态(Extrinsic State)
            2.享元模式中通常会出现工厂模式，需要创建一个享元工厂来负责维护一个享元池(Flyweight Pool)用于存储具有相同内部状态的享元对象
            3.在实际使用中，能够共享的内部状态是有限的，因此享元对象一般都设计为较小的对象，它所包含的内部状态较少，这种对象也称为细粒度对象
    - 优点：1.享元模式的优点在于它可以极大减少内存中对象的数量，使得相同对象或相似对象在内存中只保存一份
            2.享元模式的外部状态相对独立，而且不会影响其内部状态，从而使得享元对象可以在不同的环境中被共享
    - 缺点：1.元模式使得系统更加复杂，需要分离出内部状态和外部状态，这使得程序的逻辑复杂化
            2.为了使对象可以共享，享元模式需要将享元对象的状态外部化，而读取外部状态使得运行时间变长
- ## `Composite 组合`
    - 将对象组合成树状层次结构，使用户对单个对象和组合对象具有一致的访问性
    - 组合模式包含如下角色：
        - Component:抽象构件角色，组合中的对象声明接口，在适当的情况下，实现所有类共有接口的默认行为。声明一个接口用于访问和管理Component子部件
        - Leaf:树叶构件角色,在组合中表示叶子结点对象，叶子结点没有子结点
        - Composite:树枝构件角色,定义有子部件的组合部件行为，存储子部件，在Component接口中实现与子部件有关的操作。
    - 说明：1.部分整体模式，是用于把一组相似的对象当作一个单一的对象。组合模式依据树形结构来组合对象，用来表示部分以及整体层次。这种类型的设计模式  
            属于结构型模式，它创建了对象组的树形结构。
           2.这种模式创建了一个包含自己对象组的类。该类提供了修改相同对象组的方式，客户程序可以像处理简单元素一样来处理复杂元素，从而使得客户程序与复杂元素的内部结构解耦。
    - 优点：1.高层模块(客户端)调用简单。组合模式使得客户端代码可以一致地处理单个对象和组合对象，无须关心自己处理的是单个对象，还是组合对象，  
            这简化了客户端代码；
            2.节点自由增加,更容易在组合体内加入新的对象，客户端不会因为加入了新的对象而更改源代码，满足“开闭原则”
    - 缺点：1.在使用组合模式时，其叶子和树枝的声明都是实现类，而不是接口，违反了依赖倒置原则
            2.设计较复杂，客户端需要花更多时间理清类之间的层次关系；
            3.不容易限制容器中的构件；
            4.不容易用继承的方法来增加构件的新功能；