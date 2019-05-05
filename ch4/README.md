# 复合数据类型

* 复合数据类型由基本数据类型以各种方式组成.
* 四种复合数据类型:
   1. 数组
   2. slice
   3. map
   4. struct
* 数组和slice都是聚合类型,值在内存中以一组变量构成

## 4.1 数组

1. 具有固定长度,元素类型相同.
2. 在数组字面量里,若`...`出现在申明数组长度的位置,那么数组的长度由初始化数组的元素个数决定.
3. 数组的长度为数组数据类型的一部分[1]int与[2]int不一样.长度在编译时就固定.
4. 相同的数组类型可以互相比较(指的是长度相同的数组).
5. 当数组作为参数传递给函数时,Go会把数组以及其他类型都作为值传递.若需要使用应用传递,可以使用指针处理.

## 4.2 slice

1. 相同类型元素的可变长度序列. 通常写成[]T

2. slice可以访问数组的部分或者全部元素, 这个数组为该slice的底层数组.

3. slice有三个属性,指针,长度,容量:

   1. 指针指向数组中第一个可以从slice访问的元素.
   2. 长度指slice中的元素个数, 不能超过容量.
   3. 容量的大小指从slice的起始元素到底层数组的最后一个元素间元素的个数.

4. 使用内置函数len() cap() 可以获取slice的长度和容量.

5. slice操作符 s[i:j] o<= i,j<= cap(s), 此处的s可以表示数组, 指向数组的指针,slice.

6. slice引用超过容量会宕机,但是超过被引用长度只是会比原来的长度长.

   ```go
   months := [...]string{1:January, /* ... */, 12:Decemeber}	//声明数组
   summer := months[6:9]
   fmt.Println(summer[:20])	//宕机: 越界
   fmt.Println(summer[:5])		//元素长度比summer长一个, 多了"October"
   ```

7. slice中包含指向数组的指针 ,创建一个数组slice就是为数组创建了一个别名.

8. 内置函数make()可以创建一个指定元素类型,长度,容量的slice.容量参数可以省略. 

   ```go
   make([]T, len)
   make([]T, len, cap)	//与make([]T,cap)[:len]作用相同
   ```

   * 原理为make创建了一个只能使用slice访问的数组并返回其slice.
   * 第一行代码返回了整个数组slice.
   * 第二行创建了一个容量为cap 的数组,返回长度为len的slice. 



####  4.2.1 append函数

1. ```go
   var x []int
   x = append(x,1)
   x = append(x,2,3)
   x = append(x,x...)	//追加x中所有的元素
   fmt.Println(x)	//'[1 2 3 1 2 3]'
   ```

   * append函数用法如上,可以给slice添加单个或者多个元素.甚至另一个slice里所有的元素.
   * 函数第一个参数位置放的是被追加的元素.

2.  对于任何函数, 只要有可能改变slice的长度或者容量或者是改变指向的底层数组,都学要跟新slice变量.

3. ```go
   type IntSlice struct {
       ptr		*int
       len,cap  int
   }
   ```

   slice从第二点的角度看,不像是纯引用类型,更像是聚合类型,如上面代码.

4. append函数的容量分配策略较为复杂, 可以先可以简化理解为线性增长.

   ```go
   cap = sqrt(2,n)	//0, 1, 2, 4, 8, ...
   ```

5. append函数参数申明用到了"...",表示参数可接受可变长度的参数列表

   ```go
   func append(x []T, y ...T){...}	
   ```

   参数y可以为任意个(不为零).

6. slice的容量改变意味着底层数组重新分配和复制.复制slice可以使用内置copy函数:

   ```go
   var x,y []int
   z := copy(x,y)
   ```

   copy函数返回参数中两者最小的长度, 这样就不会发生slice越界.





#### 4.2.2 slice就地修改

1. 对slice进行操作的时,输入的slice和输出的slice如果拥有相同的底层数组, 可以避免函数内部重新分配数组.这种情况下,底层数组被修改.

2. slice可以用来实现栈:

   ```go
   //stack 为空slice
   //push v
   stack = append(stack,v)	
   //栈顶
   top := stack[len(stack)-1]
   //pop
   stack = stack[:len(stack)-1]
   //移除中间的元素
   copy(stack[i:],stack[i+1:])
   ```

   



## 4.3 map

1. 散列表是一个拥有键值对元素的无序集合. 

   * 键值唯一,键对应的值通过键来操作.
   * 无论散列表多大, 对键的比较操作都在常量时间.

2. map是go语言里散列表的引用.map类型为map[k]v.

   * k为键,v为值.
   * map的键和值都是对应拥有相同的数据类型, 但是键和值的数据类型不一定相同.

3. 内置函数make可以创建map

   ```go
   //sting到int的map
   ages := make(map[string]int)
   //添加nick为键值并赋值为31
   ages["nick"] = 31
   ```

4. 使用map字面量创建

   ```go
   ages := map[string]int{
       "nick" = 31,
   }
   ```

   空map的表达式为map[string]int{}.可以用这个表达式对map进行初始化.

5. 内置函数delete根据键移除元素.若元素不存在返回值类型的零值,操作安全.

   ```go
   delete(ages,"nick")
   ```

6. 不可以获取map的地址:

   * map不是变量，不能获取地址
   * map的增长会导致地址的变化.

7. 迭代map元素顺序是不固定的, 一般情况下认为是随机的.

8. map类型零值为nil, 大部分对于map的操作可以安全的在nil的map上执行. 但是不能对零值的map进行赋值,会导致宕机错误.设置元素之前要先进行初始化,初始化操作见2和3.

9.  如果需要判断一个键是否存在于map之中:

   ```go
   ages := map[string]int{}
   	if age, ok := ages["nick"]; !ok { /* 0 false */
       println(age, ok)
   }
   ```

   age为键指向的值, false表示键不存在.

10. 

## 4.4 结构体

1. 结构体是将零个或者多个任意变量类型组合在一起的聚合数据类型.

   * 每个变量叫做结构体的成员.
   * 结构体可以实例化,叫做结构体变量

2. 下面语句声明一个叫User的结构体和一个结构体变量nick:

   ```go
   //声明结构体
   type User struct {
       Name string
       Age  int
   }
   //实例化结构体 或者说声明结构体变量
   var nick User
   //为该结构体赋值
   nick.Name ="causingbrick"
   nick.Age = 999
   //获取nick Name的地址
   position := &nick.Name
   
   ```

   * 每个结构体变量的成员都由`.`来访问.

   * 结构体变量的成员变量地址可以获取,并可以通过指针访问, 并且`.`可以用在结构体指针上

     ```go
     var student *User = &nick
     //将nick的名字改为nick
     student.Name = "nick"
     //等价于
     (*student).Name = "nick"
     ```

3. 成员变量的顺序对于结构体同一性很重要,将顺序交换之后结构体类型会发生改变.

4. 成员变量可以通过首字母是否大小写控制是否导出.

5. 结构体的成员变量不能包含自己, 但是可以包含一个自身的类型的指针.

### 4.4.1 结构体字面量

1. 结构体类型的值通过结构体类型自变量来设置,通常有两种方法:

   * ```go
     type point struct {x, y int}
     p := Point{1,2}
     ```

     适用于简单声明, 可读性差.

   * 

## 4.5 JSON

