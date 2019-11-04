

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

3. 对于没有函数体的函数声明,则表示该函数不是以Go实现的.这样的声明定义了函数标识符.如标准库math包里的:

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
   
     ```go
     func size(rect image.Rectangle)(width, height int)
     ```
   
   * 习惯上最后的一个布尔返回值表示成功与否,但`error`通常不需要额外说明.
   
2. 一个多值调用可以作为单独的实参传递给拥有多个形参的函数中.

   ```go
   log.Println(findlinks(url))
   ```

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

## 5.6 匿名函数

1. `func`关键字后面没有函数的名称.为一个表达式,值称为 匿名函数.
2. 函数变量类似于使用闭包方法实现的变量,通常把函数变量称为闭包.
   ```go
   func main() {
      f := squares()
      fmt.Println(f())  //1
      fmt.Println(f())  //4
      fmt.Println(f())  //9
   }

   func squares() func() int {
      var x int
      return func() int {
         x++
         return x * x
      }
   }
   ```
   squares返回一个匿名函数, 其中x为闭包,这个例子还说明, 函数值f会保存状态.

#### 5.6.1警告：捕获迭代变量

给出下面删除文件夹的例子

```go
var rmdirs []func()
//temp
for _, d := range tempDirs{
    dir := d
    os.MkdirAll(dir,0777)	//creats parent directories too
    rmdirs = append(rmdirs, func(){
        os.RemoveAll(dir)
    })
}
```

在range内部里的`dir := d`这行不能缺少, 若缺少掉直接对d进行操作, 则删除的时候只会删除最后一个文件路径.如下分析: dir为range这个词法块创建的局部变量, 内部共享该变量, 而对于函数值而言, 里面存的是变量的指针而不是变量的值, 故循环结束之后, 函数值里面的dir都是同一个指针, 指向最后一次遍历的值.当后续执行rmdir里面的函数时, 都是同一个值.

```go
var rmdirs []func()

for _, dir := range tempDirs{
    os.MkdirAll(dir,0777)	//creats parent directories too
    rmdirs = append(rmdirs, func(){
        os.RemoveAll(dir)
    })
}
```

再比如:

```go
var any []func()

for i := 0; i < 3; i++ {
    d := i
    any = append(any, func() {
        println(d)
    })
}
for _, v := range any {
    v()//0, 1, 2
}
```

若不使用d暂存则

```go
var any []func()

for i := 0; i < 3; i++ {
    any = append(any, func() {
        println(i)
    })
}
for _, v := range any {
    v()//2, 2, 2
}
```



## 5.7 可变参数

对于函数传递参数时,可以提供可变的参数个数, 在参数列表最后的类型名称之前使用省略号`...`表示可变参数, 表示可以传递该类型任意数目的参数.

```go
func sum(vals ...int) int {
    total := 0
    for _,val := range vals {
        total += val
    }
    return tatal
}
```

调用时显示申请一个数, 将参数复制到数组并将这个数组slice传递到函数.

若已经是一个slice在最后一个参数后面放一个省略号即可

```go
values := []int{1, 2, 3, 4}
fmt.Println(sum(values...))
```

**虽然`...int`像函数体内slice, 但是并不是同一种类型**

## 5.8 `deffer`语句

- deffer语句是go里用来处理延时函数的语句, 放在deffer语句后面的语句, 会在最后函数结束返回时被执行(无论函数是否正常返回),类似于栈机制, 每个次deffer后面的语句都被压入栈, 最后返回时从栈顶取出内容并执行.

- 语法

  > deffer 函数(或者方法调用)

- deffer语句常用于成对的操作, 比如打开关闭, 连接和断开, 加锁解锁.又因为延迟函数的执行在return语句之后, 故可以从deffer后面的语句修改函数的返回值(利用匿名函数可以访问其外层函数作用域的特点).



## 5.9  `Panic`宕机

- go 类型系统会在编译器很多错误, 但有些错误需要在运行时检查, 当检测到这种错误时就是panic异常, 即宕机.如slice越界, 空指针引用.

- 运行过程: 一般而言, 宕机会使程序执行终止, 并且`goroutine`中的延迟函数被执行(deffer机制), 然后程序异常退出并且留下日志信息.其中日志信息包括宕机的值, 其中包含了函数的栈跟踪信息, 可以借助这条信息来诊断信问题的原因,.

- panic可以在调用内置的panic函数是发生.当逻辑碰到不能发生的状况时, 宕机是最好的处理方式, 但是会导致程序异常退出, 故只有在严重错误时才会使用panic

- [Golang panic用法](https://www.cnblogs.com/liuzhongchao/p/10112739.html)

  Go语言追求简洁优雅，所以，Go语言不支持传统的 try…catch…finally 这种异常，因为Go语言的设计者们认为，将异常与控制结构混在一起会很容易使得代码变得混乱。因为开发者很容易滥用异常，甚至一个小小的错误都抛出一个异常。在Go语言中，使用多值返回来返回错误。不要用异常代替错误，更不要用来控制流程。在极个别的情况下，也就是说，遇到真正的异常的情况下（比如除数为 0了）。才使用Go中引入的Exception处理：defer, panic, recover。

  这几个异常的使用场景可以这么简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。

  ```go
package main
  
  import "fmt"
  
  func main(){
  
      defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
          fmt.Println("c")
          if err:=recover();err!=nil{
              fmt.Println(err) // 这里的err其实就是panic传入的内容，55
          }
          fmt.Println("d")
      }()
  
      f()
  }
  
  func f(){
      fmt.Println("a")
      panic(55)
      fmt.Println("b")
      fmt.Println("f")
  }
  
  输出结果：
  a
  c
  55
  d
  exit code 0, process exited normally.
  ```

## 5.10 `Recover`捕获异常

异常panic是可以恢复的,阻止程序的奔溃.一般而言,不应该对`panic`做任何处理, 比如在web服务器遇到不可预料的严重问题时, 在奔溃时应该先关闭所有连接.

- [Golang panic用法](https://www.cnblogs.com/liuzhongchao/p/10112739.html)

  Go语言追求简洁优雅，所以，Go语言不支持传统的 try…catch…finally 这种异常，因为Go语言的设计者们认为，将异常与控制结构混在一起会很容易使得代码变得混乱。因为开发者很容易滥用异常，甚至一个小小的错误都抛出一个异常。在Go语言中，使用多值返回来返回错误。不要用异常代替错误，更不要用来控制流程。在极个别的情况下，也就是说，遇到真正的异常的情况下（比如除数为 0了）。才使用Go中引入的Exception处理：defer, panic, recover。

  这几个异常的使用场景可以这么简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。

  ```go
  package main
  
  import "fmt"
  
  func main(){
  
      defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
          fmt.Println("c")
          if err:=recover();err!=nil{
              fmt.Println(err) // 这里的err其实就是panic传入的内容，55
          }
          fmt.Println("d")
      }()
  
      f()
  }
  
  func f(){
      fmt.Println("a")
      panic(55)
      fmt.Println("b")
      fmt.Println("f")
  }
  
  输出结果：
  a
  c
  55
  d
  exit code 0, process exited normally.
  ```

 