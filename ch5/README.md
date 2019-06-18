[TOC]

# 函数

##  5.1 函数声明

1. 函数声明

   ```go
   func name(parameter-list)(result-list){
   	body
   }
   ```

   * parameter-list 为形参列表, 参数作为此函数的局部变量.`函数声明时的参数为形参,调用者提供的对应形参的值为实参`.
   * 参数列表中若要强调某个参数未被使用.可以使用`_`符号在参数位置代替.
   * 函数的形参知识实参的值复制,函数对形参操作时不改变实参的值,除非传入引用类型,这也是通过函数改变实参的值的一个技巧.
   * result-list 为返回值.单返回值时若无参数字面量可以不用括号.Go语言中函数可以有多返回值, 按照顺序对应`return`.

2. `函数类型` 也称为函数的标识符.若函数的形参列表与返回值列表的变量类型一一对应,那么它们被看成相同的函数类型和标识符.

   * 若函数的参数列表和返回值列表中有相同的类型,不必为每一个都写出类型,可以写成`a, b, ... Type`的形式.

3. 对于没有哦函数体的函数声明,则表示该函数不是以Go实现的.这样的声明定义了函数标识符.如标准库math包里的:

   ```go
   package math
   
   func Sin(x float64) float	//汇编实现
   ```



## 5.2 递归

1. 函数递归是指函数可以直接或间接的调动自己.
2. 编程语言使用固定长度的函数调用栈;大小在64kb到2MB之间. 递归的深度受限于栈的大小, 当进行深度递归调用的时候必须谨防`栈溢出`.
3. Go的栈实现了可变的长度,栈的大小可随着使用的而增长,最大可以申请到1GB.



## 5.3 多返回值

1. Go中函数支持多返回值.
   * 返回值可以不设置字面量, 但良好的名称使得返回值更有意义,特别是返回多个结果且类型相同的时候.
   * 习惯上最后的一个布尔返回值表示成功与否,但`error`通常不需要额外说明.
2. 一个多值调用可以作为单独的实参传递给拥有多个形参的函数中.
3. 函数若有命名的返回值, 可以省略`return`语句的操作数,称为`裸返回`.应该谨慎使用.



## 5.4 错误

1. Go中函数运行失败会返回错误信息,错误信息被任务是预期的值而不是异常`exception`.错误处理是Go程序的一个重要组成部分,程序运行失败是几个预期结果之一.
2. 异常不是错误, 异常处理机制用于处理未被预料到的错误, 即bug.
3. 内置的`error`是接口类型.`nill`表示函数运行成功,`non-nill`表示运行失败.
4. Go使用控制流机制来处理异常(如`if`和`return`).



### 5.4.1 错误处理策略

* Go中大部分的函数代码结构几乎相同:首先是一些列的初始检查,防止错误发生,之后是函数的实际逻辑.通常将处理失败的逻辑代码放在处理成功的代码之前.若错误导致函数返回,那么成功的逻辑代码不应该放在`else`语句块中,应直接放在函数体里.

* 根据情况不同,应该选则不同的错误处理策略,以下是五中常用的方式:

1. 传递错误.某个子程序的失败,会变成该函数的失败.

   * 当向上层传递错误时可以使用`fmt.Errorf`格式化错误信息并返回.应该使用函数前缀添加上额外的上下文信息到原始错误信息.到`main`函数处理时,错误信息应该提供清晰的从原因到后过的因果链.
   * 由于错误信息经常是以链式组合在一起的,所以错误信息中应该避免大写和换行符.
   * 编写错误信息的时候, 确保错误信息对问题的细节描述是详细的.
   * 注意错误信息表达的一致性,相同的函数或同包内的同一组函数返回的错误在构成和处理方式上是相似的.
   * 通常而言,北调函数会将`调用信息`和`参数信息`作为发生错误时的上下文放在错误信息中并返回给调用者,函数调用时时调用者需要添加一些错误信息不包含的信息.

2. 若错误的发生是偶然性的,或不可预知的问题导致的,可以重新尝试失败的操作,重试时,需要限制重试的时间间隔或重视的次数防止无限制的重试.

   ```go
   //!+
   // WaitForServer attempts to contact the server of a URL.
   // It tries for one minute using exponential back-off.
   // It reports an error if all attempts fail.
   func WaitForServer(url string) error {
   	const timeout = 1 * time.Minute
   	deadline := time.Now().Add(timeout)
   	for tries := 0; time.Now().Before(deadline); tries++ {
   		_, err := http.Head(url)
   		if err == nil {
   			return nil // success
   		}
   		log.Printf("server not responding (%s); retrying...", err)
   		time.Sleep(time.Second << uint(tries)) // exponential back-off
   	}
   	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
   }
   
   //!-
   ```

3. 若程序发生错误,无法继续运行,应该输出错误信息并结束程序.

   * 这种策略只应该在`main`函数中使用,对于库函数而言,应仅向上传递错误,除非程序内部bug,才能在库函数中结束程序.

   * `log`包中的所有函数都会默认在错误信息之前输出信息,长期运行的服务器采用默认的时间格式,可以设置`log`的前置信息屏蔽时间信息,一般而言,前缀信息会被设置成命令名字.

     ```go
     log.SetPrefix("wait: ")
     log.SetFlags(0)
     ```

   * `log`包中所有的函数会为没有换行符的字符增加换行符.

   * 交互式工具很少采用包含如此多信息的格式可以使用`fmt`包里的函数实现类似功能.

4. 只需要输出信息,不需要中断程序的运行.

5. 直接忽略掉错误.

   * 采用该策略时应该谨慎, 虽然程序无错误处理,但是逻辑不会受到该错误的影响.
   * 应该在每次函数调用之后,都养成错误处理的习惯,当采用此策略时,应该清晰记录下意图.



### 5.4.2 文件结尾错误`EOF`

* `io`包保证任何由文件结束引起的读取失败都返回同一个错误`io.EOF`.
* 因为文件结束这种错误不需更多的描述，所以有固定的错误`EOF`.

##  5.5 函数值

1. 在Go中函数为第一类型值(变量), 函数也可以作为变量类型.并且可以传递给变量或者从其他函数返回,函数变量可以像其他函数一样调用:

   ```go
   func square(n int) int	{return n*n }
   func product(m, n int) {return m*n }
   s := square 
   fmt.Println(s(3))
   s = product		//编译错误	两个函数类型不匹配
   ```

2. 函数类型的零值为`nil`,调用一个空的函数变量会宕机.  函数变量可以和空值比较但是 不可以和函数变量比较.由此不可以作为键值出现于map中.

3.  

## 5.6 匿名函数

1. `func`关键字后面没有函数的名称.为一个表达式,值称为 匿名函数.
2. 函数变量类似于使用闭包方法实现的变量,通常把函数变量称为闭包.

### 5.6 警告：捕获迭代变量



## 5.7 可变参数



## 5.8 `Deferred`函数



## 5.9  `Panic`异常



## 5.10 `Recover`捕获异常




