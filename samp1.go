package main

import (
	"fmt"
	// "time"
)

// const NUMBER = 1000000

// func test() {
// 	for {
// 	}
// }

// func main() {
// 	fmt.Println(time.Now().UnixNano())
// 	for i := 0; i < NUMBER; i++ {
// 		go test() // 启动新的goroutine（用户态线程，免强占，不切换）， 返回值会被抛弃
// 	}
// 	fmt.Println(time.Now().UnixNano())
// 	for {
// 	}

//  fmt.Scan(&a)  接收一个用户输入
// }

// // 共享内存（设置全局变量，锁变量，见sharemengo）， 消息(封装性好，，是进程内的通信方式，传递对象的过程和调用函数参数传递行为类似，也可以传递指针。类型相关，见channel)

func generateString(strings chan string) {
	strings <- "Monday"
	strings <- "Tuesday"
	strings <- "Wednesday"
	strings <- "Thursday"
	strings <- "Friday"
	strings <- "Saturday"
	strings <- "Sunday"
	close(strings)
}

func main() {
	strings := make(chan string) // 无缓冲channel
	go generateString(strings)

	for s := range strings {
		fmt.Println(s)
	}
	// for {
	//     if s, ok := <-strings; ok {
	//         fmt.Println(s)
	//     } else {
	//         fmt.Println("channel colsed.")
	//         break
	//     }
	// }
}

// √ golang中的select关键字用于处理异步IO，可以与channel配合使用。

// √ golang中的select的用法与switch语言非常类似，不同的是select每个case语句里必须是一个IO操作,必须是一个通信操作，要么是发送要么是接收。

// √ select会一直等待等到某个case语句完成才结束。select随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。一个默认的子句应该总是可运行的。

// select {
// case <-chan1:
//     // 如果chan1成功读到数据，则进行该case处理语句
// case chan2 <- 1:
//     // 如果成功向chan2写入数据，则进行该case处理语句
// default:
//     // 如果上面都没有成功，则进入default处理流程
// }

select 语句的语法：

每个case都必须是一个通信
所有channel表达式都会被求值
所有被发送的表达式都会被求值
如果任意某个通信可以进行，它就执行；其他被忽略。
如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。 
否则：
如果有default子句，则执行该语句。
如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。


func main() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	   case i1 = <-c1:
		  fmt.Printf("received ", i1, " from c1\n")
	   case c2 <- i2:
		  fmt.Printf("sent ", i2, " to c2\n")
	   case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		  if ok {
			 fmt.Printf("received ", i3, " from c3\n")
		  } else {
			 fmt.Printf("c3 is closed\n")
		  }
	   default:
		  fmt.Printf("no communication\n")
	}    
 }
 以上代码执行结果为：
 
 no communication






// func Parse(ch <-chan int) {
//     for value := range ch {
//         fmt.Println("Parsing value", value)
//     }
// }

// func main() {
//     var ch chan int
//     ch = make(chan int)

//     go func() {
//         ch <- 1
//         ch <- 2
//         ch <- 3
//         close(ch)
//     }()

//     Parse(ch)
// }

// 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，
// 如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；
// 标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。


变量声明： 写出类型，类型推断， := 符号（省略var，自动推断，只能在函数内部使用。对于一个变量只能用一次，之后没有‘：’）
//类型相同多个变量, 非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3
var vname1, vname2, vname3 = v1, v2, v3 //和python很像,不需要显示声明类型，自动推断
vname1, vname2, vname3 := v1, v2, v3 //出现在:=左侧的变量不应该是已经被声明过的，否则会导致编译错误



// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)

var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
g, h := 123, "hello"



全局变量可以声明并且不使用。局部变量（函数内的）声明之后必须初始化，同时必须使用


基本类型值存储在栈中，并且是值拷贝，取地址取值声明指针同C语言
一般更复杂的数据会包含多个字节，是引用保存，地址拷贝。即C中指针



_,a = 1,3  // 1 is dropped



常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
const ....  也有类型推断，多个变量同时声明
还可以用作枚举：
const (
    Unknown = 0
    Female = 1
    Male = 2
)
常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：
const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)  // 字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度。  unsafe.Sizeof("hello") 为16， 两部分每部分都是 8 个字节。
)

iota，特殊常量，可以认为是一个可以被编译器修改的常量。
在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。

iota 可以被用作枚举值：
const (
    a = iota
    b = iota
    c = iota
)
第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：
const (
    a = iota
    b
    c
)

示例
func main() {
    const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)
}

实例：
const (
    i=1<<iota
    j=3<<iota
    k
    l
)
func main() {
    fmt.Println("i=",i)
    fmt.Println("j=",j)
    fmt.Println("k=",k)
    fmt.Println("l=",l)
}
以上实例运行结果为：
i= 1
j= 6
k= 12
l= 24



switch 不用写break，支持case后多匹配结果，支持type switch，支持将switch后的表达式移到case后，如下：
switch marks {
	case 90,95,96: grade = "A"     // 多条件匹配
	case 80: grade = "B"
	case 50,60,70 : grade = "C"
	default: grade = "D"  
 }
 switch {
	case grade == "A" :
	   fmt.Printf("优秀!\n" )     
	case grade == "B", grade == "C" :
	   fmt.Printf("良好\n" )      
	case grade == "D" :
	   fmt.Printf("及格\n" )      
	case grade == "F":
	   fmt.Printf("不及格\n" )
	default:
	   fmt.Printf("差\n" );
 }

 type switch 例子如下:
 var x interface{}
 switch i := x.(type) {
	case nil:      
	   fmt.Printf(" x 的类型 :%T",i)                
	case int:      
	   fmt.Printf("x 是 int 型")                       
	case float64:
	   fmt.Printf("x 是 float64 型")           
	case func(int) float64:
	   fmt.Printf("x 是 func(int) 型")                      
	case bool, string:
	   fmt.Printf("x 是 bool 或 string 型" )       
	default:
	   fmt.Printf("未知型")     
 }   

 如果想要执行多个 case，需要使用 fallthrough 关键字，也可用 break 终止：
 switch{
case 1:
...
if(...){
	break
}

fallthrough // 此时switch(1)会执行case1和case2，但是如果满足if条件，则只执行case1

case 2:
...
case 3:
}



Go没有while结构，注意for使用方法。一下代码是找到100以内的全部素数
import "fmt"
func main(){
    // var count,c int   //定义变量不使用也会报错
    var count int
    var flag bool
    count=1
    //while(count<100) {    //go没有while
    for count < 100 {
        count++
        flag = true;
        //注意tmp变量  :=
        for tmp:=2;tmp<count;tmp++ {
            if count%tmp==0{
                flag = false
            }
        }

        // 每一个 if else 都需要加入括号 同时 else 位置不能在新一行
        if flag == true {
            fmt.Println(count,"素数")
        }else{
        continue
        }
    }
}


go的if条件里面可以声明临时变量
if num := 9; num < 0 {
	fmt.Println(num, "is negative")
} else if num < 10 {
	fmt.Println(num, "has 1 digit")
} else {
	fmt.Println(num, "has multiple digits")
}
// 9 has 1 digit
（1） 不需使用括号将条件包含起来

（2） 大括号{}必须存在，即使只有一行语句

（3） 左括号必须在if或else的同一行

（4） 在if之后，条件语句之前，可以添加变量初始化语句，使用；进行分隔

（5） 在有返回值的函数中，最终的return不能在条件语句中





Go的3种for结构：   
for init; condition; post { }
和 C 的 while 一样：
for condition { }
和 C 的 for(;;) 一样：
for { }

for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
for key, value := range oldMap {
    newMap[key] = value
}

示例：
numbers := [6]int{1, 2, 3, 5} 
/* for 循环 */
for a := 0; a < 10; a++ {
	fmt.Printf("a 的值为: %d\n", a)
}
for a < b {
	a++
	fmt.Printf("a 的值为: %d\n", a)
}
for i,x:= range numbers {
	fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
}   


Go中的goto语句，Go 语言的 goto 语句可以无条件地转移到过程中指定的行。， 可以构成循环，条件转移，跳出循环体
goto label;
..
label: statement;

LOOP: for a < 20 {
	if a == 15 {
	   /* 跳过迭代 */
	   a = a + 1
	   goto LOOP
	}
	fmt.Printf("a的值为 : %d\n", a)
	a++     
 }  




Go语言函数
有时不需要返回值，有返回值的时候写出返回类型， 返回多个值的时候需要写出多个类型：
func swap(x, y string) (string, string) {
	return y, x
 }

函数变量：
import (
	"fmt"
	"math"
 )
 func main(){
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
	   return math.Sqrt(x)
	}
	/* 使用函数 */
	fmt.Println(getSquareRoot(9))
 }

 函数闭包：
 匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
 以下实例中，我们创建了函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量
func getSequence() func() int {
   i:=0
   return func() int {
      i+=1
     return i  
   }
}
func main(){
   /* nextNumber 为一个函数，函数 i 为 0 */
   nextNumber := getSequence()
   /* 调用 nextNumber 函数，i 变量自增 1 并返回 */
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   /* 创建新的函数 nextNumber1，并查看结果 */
   nextNumber1 := getSequence()  
   fmt.Println(nextNumber1())
   fmt.Println(nextNumber1())
}
// 1
// 2
// 3
// 1
// 2

带参数的闭包函数调用:
func main() {
    add_func := add(1,2)
    fmt.Println(add_func())
    fmt.Println(add_func())
    fmt.Println(add_func())
}
// 闭包使用方法
func add(x1, x2 int) func()(int,int)  {
    i := 0
    return func() (int,int){
        i++
        return i,x1+x2
    }
}
闭包带参数补充:
func main() {
    add_func := add(1,2)
    fmt.Println(add_func(1,1))
    fmt.Println(add_func(0,0))
    fmt.Println(add_func(2,2))
} 
// 闭包使用方法
func add(x1, x2 int) func(x3 int,x4 int)(int,int,int)  {
    i := 0
    return func(x3 int,x4 int) (int,int,int){ 
       i++
       return i,x1+x2,x3+x4
    }
}


GO语言中的方法就是包含接受者（命名类型、结构体的数据结构或是一个指针）的函数，所有给定类型的方法属于该类型的方法集合。：
/* 定义结构体 */
type Circle struct {
	radius float64
}
func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) = ", c1.getArea())
}
//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}


Go语言同C语言的变量作用域： 局部变量（包括形参），全局变量。 同名的话局部覆盖全局
全局变量可以再整个包（package），或者被导出的外部包中使用。


Go 语言数组：
函数形参若是数组， 可以写出维数，也可以不写
var balance [10] float32
var balance = [  ]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
a = [3][4]int{  
	{0, 1, 2, 3} ,   /*  第一行索引为 0 */
	{4, 5, 6, 7} ,   /*  第二行索引为 1 */
	{8, 9, 10, 11}   /*  第三行索引为 2 */
   }

Go指针：
声明，赋值，取值 同 C。
定义没赋值之前为nil。
指针数组： var ptr [3]*int;
指针的指针： var ptr **int;



Go结构体：
可以方便的传给函数，定义指针（结构体指针取成员不用->，用点.）。
type Books struct {
	title string
	author string
	subject string
	book_id int
}
func main() {
	var Book1 Books
	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407 
	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	// 忽略的字段为 0 或 空
fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}



Go的接口：
必须要和结构体一起用
type Phone interface {
    call()
}
type NokiaPhone struct {
}
func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}
type IPhone struct {
}
func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}
func main() {
    var phone Phone
    phone = new(NokiaPhone)
    phone.call()
    phone = new(IPhone)
    phone.call()
}
另一个示例:
type Man interface {
    name() string;
    age() int;
}
type Woman struct {
}
func (woman Woman) name() string {
   return "Jin Yawei"
}
func (woman Woman) age() int {
   return 23;
}
type Men struct {
}
func ( men Men) name() string {
   return "liweibin";
}
func ( men Men) age() int {
    return 27;
}
func main(){
    var man Man;
    man = new(Woman);
    fmt.Println( man.name());
    fmt.Println( man.age());
    man = new(Men);
    fmt.Println( man.name());
    fmt.Println( man.age());
}


Go的错误处理:
type error interface {
    Error() string
}
我们可以在编码中通过实现 error 接口类型来生成错误信息。

函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息：
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // 实现
}
在调用Sqrt的时候传递的一个负数，然后就得到了non-nil的error对象
result, err:= Sqrt(-1)
if err != nil {
   fmt.Println(err)
}

完整示例:
// 定义一个 DivideError 结构
type DivideError struct {
    dividee int
    divider int
}
// 实现 `error` 接口
func (de *DivideError) Error() string {
    strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
    return fmt.Sprintf(strFormat, de.dividee)
}
// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
    if varDivider == 0 {
        dData := DivideError{
            dividee: varDividee,
            divider: varDivider,
        }
        errorMsg = dData.Error()
        return
    } else {
        return varDividee / varDivider, ""
    }
}
func main() {
    // 正常情况
    if result, errorMsg := Divide(100, 10); errorMsg == "" {
        fmt.Println("100/10 = ", result)
    }
    // 当被除数为零的时候会返回错误信息
    if _, errorMsg := Divide(100, 0); errorMsg != "" {
        fmt.Println("errorMsg is: ", errorMsg)
    }
}


Go切片：
是对数组的抽象。Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，
Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
var slice1 []type              未初始化之前slice1默认为 nil，长度为 0。 可用于判断
切片不需要说明长度。
或使用make()函数来创建切片:
var slice1 []type = make([]type, len)
也可以简写为
slice1 := make([]type, len)
也可以指定容量，其中capacity为可选参数。
make([]T, length, capacity)
这里 len 是数组的长度并且也是切片的初始长度。
s :=[] int {1,2,3 } 
直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3

s := arr[:] 
初始化切片s,是数组arr的引用

s := arr[startIndex:endIndex] 
将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片

s := arr[startIndex:] 
缺省endIndex时将表示一直到arr的最后一个元素

s := arr[:endIndex] 
缺省startIndex时将表示从arr的第一个元素开始

s1 := s[startIndex:endIndex] 
通过切片s初始化切片s1

s :=make([]int,len,cap) 
通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片

切片是可索引的，并且可以由 len() 方法获取长度。切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。 索引方法与切片类似python。
从拷贝切片的 copy 方法和向切片追加新元素的 append 方法： 
var numbers []int
 /* 允许追加空切片 */
 numbers = append(numbers, 0)
/* 同时添加多个元素 */
numbers = append(numbers, 2,3,4)
/* 创建切片 numbers1 是之前切片的两倍容量*/
numbers1 := make([]int, len(numbers), (cap(numbers))*2)
/* 拷贝 numbers 的内容到 numbers1 */
copy(numbers1,numbers)

从切片生成的切片，实际上底层是指向原切片的数据结构，其cap = 原切片cap-新切片在原切片的起点。
len切片<=cap切片<=len数组，切片由三部分组成：指向底层数组的指针、len、cap



Go中的range:
用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对的 key 值
for _, num := range nums {
	sum += num
}
fmt.Println("sum:", sum)
//range也可以用在map的键值对上。
kvs := map[string]string{"a": "apple", "b": "banana"}
for k, v := range kvs {
	fmt.Printf("%s -> %s\n", k, v)
}
//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
for i, c := range "go" {
	fmt.Println(i, c)
}


Go中的map:
无序，使用for range遍历的时候无顺序
/* 声明变量，默认 map 是 nil */   不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
var 变量名 map[key类型]值类型
/* 使用 make 函数 */
map_variable := make(map[key类型]值类型)
示例：
var countryCapitalMap map[string]string /*创建集合 */
countryCapitalMap = make(map[string]string)
for country := range countryCapitalMap {
	fmt.Println(country, "首都是", countryCapitalMap [country])
}
/*查看元素在集合中是否存在 */ ok是bool， captial是value
captial, ok := countryCapitalMap [ "美国" ] /*如果确定是真实的,则存在,否则不存在 */

countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
delete(countryCapitalMap,  France)




下面列举了 Go 代码中会使用到的 25 个关键字或保留字：

break	default	func	interface	select
case	defer	go	map	struct
chan	else	goto	package	switch
const	fallthrough	if	range	type
continue	for	import	return	var
除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：

append	bool	byte	cap	close	complex	complex64	complex128	uint16
copy	false	float32	float64	imag	int	int8	int16	uint32
int32	int64	iota	len	make	new	nil	panic	uint64
print	println	real	recover	string	true	uint	uint8	uintptr



Go语言类型转换：
直接用类型名加括号去转换



布尔型
布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。
2	数字类型
整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且原生支持复数，其中位的运算采用补码。
也有基于架构的类型，例如：int、uint 和 uintptr。
uint8 uint16 uint32 uint64 int8 ... ...
float32 float64 complex64(32 位实数和虚数) complex128(64 位实数和虚数)
byte(uint8) rune(int32) uint(32 or 64) int(uint) 
3	字符串类型:
字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。Go语言的字符串的字节使用UTF-8编码标识Unicode文本。
4	派生类型:
包括：
(a) 指针类型（Pointer）
(b) 数组类型
(c) 结构化类型(struct)
(d) Channel 类型
(e) 函数类型
(f) 切片类型
(g) 接口类型（interface）
(h) Map 类型




