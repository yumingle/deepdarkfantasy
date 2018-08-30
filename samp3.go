package main

import "fmt"

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func aaaamain() {
	// var channel chan int
	// channel := make(chan int)
	// channel <- 123
	// var ch chan int            // int类型channel
	// var m map[string]chan bool // bool类型channel的map
	// t := <-channel

	// fmt.Println(t)

	// close(cnannel)	应该在生产者处关闭

	//下面是使用channel实现goroutine全部完成的例子
	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}

// channel 超时设置2

// timeout := make(chan bool)

// go func() {
//     time.Sleep(3 * time.Second) // sleep 3 seconds
//     timeout <- true
// }()

// // 实现了对ch读取操作的超时设置。
// ch := make(chan int)
// select {
// case <-ch:
// case <-timeout:
//     fmt.Println("timeout!")
// }

// 判断channel是否关闭

// value, ok := <-chanName

// if ok {
//     // channel未关闭
// } else {
//     // channel已关闭
// }

// channel 是进程内通信，进程间要用socket和http

// channel空读取，满写入会造成阻塞，类似queue。无缓冲的channel读写通常都会发生阻塞，带缓冲的channel在channel满时写数据阻塞，在channel空时读数据阻塞。

// c := make(chan int, 1024)
// // 从带缓冲的channel中读数据
// for i:=range c {
// 　　...
// }

// 单向channel的作用有点类似于c++中的const关键字，用于遵循代码“最小权限原则”。
// √ golang中假如一个channel只允许读，那么channel肯定只会是空的，因为没机会往里面写数据。
// √ golang中假如一个channel只允许写，那么channel最后只会是满的，因为没机会从里面读数据。
// √ 单向channel概念，其实只是对channel的一种使用限制，即在将一个channel变量传递到一个函数时，可以通过将其指定为单向channel变量，从而限制该函数中可以对此channel的操作，达到权限控制作用。
// 单向channel变量的声明：

// var ch1 chan int  　　　　// 普通channel
// var ch2 chan <- int 　　 // 只用于写int数据
// var ch3 <-chan int 　　 // 只用于读int数据

// 可以通过类型转换，将一个channel转换为单向的：

// ch4 := make(chan int)
// ch5 := <-chan int(ch4)   // 单向读
// ch6 := chan<- int(ch4)  //单向写

// 在一个函数中使用单向读channel：

// func Parse(ch <-chan int) {
//     for value := range ch {
//         fmt.Println("Parsing value", value)
//     }
// }

//channel作为一种原生类型，本身也可以通过channel进行传递，例如下面这个流式处理结构：

// type PipeData struct {
//     value int
//     handler func(int) int
//     next chan int
// }

// func handle(queue chan *PipeData) {
//     for data := range queue {
//         data.next <- data.handler(data.value)
//     }
// }

// 在UNIX中，select()函数用来监控一组描述符，该机制常被用于实现高并发的socket服务器程序。Go语言直接在语言级别支持select关键字，用于处理异步IO问题，大致结构如下：

// select {
//     case <- chan1:
//     // 如果chan1成功读到数据

//     case chan2 <- 1:
//     // 如果成功向chan2写入数据

//     default:
//     // 默认分支
// }

// Go语言没有对channel提供直接的超时处理机制，但我们可以利用select来间接实现，例如：

// timeout := make(chan bool, 1)

// go func() {
//     time.Sleep(1e9)
//     timeout <- true
// }()

// switch {
//     case <- ch:
//     // 从ch中读取到数据

//     case <- timeout:
//     // 没有从ch中读取到数据，但从timeout中读取到了数据
// }
// 这样使用select就可以避免永久等待的问题，因为程序会在timeout中获取到一个数据后继续执行，而无论对ch的读取是否还处于等待状态。

// sync包提供了两种锁类型：sync.Mutex和sync.RWMutex

// var lck sync.Mutex
// func foo() {
//     lck.Lock()
//     defer lck.Unlock()
//     // ...
// }

// lck.Lock()会阻塞直到获取锁，然后利用defer语句在函数返回时自动释放锁

// 对于从全局角度只需要运行一次的代码，比如全局初始化操作，Go语言提供了一个once类型来保证全局的唯一性操作，如下：

// var flag int32
// var once sync.Once

// func initialize() {
//     flag = 3
//     fmt.Println(flag)
// }

// func setup() {
//     once.Do(initialize)
// }

// func main() {

//     setup()
//     setup()
// }
// flag只别打印 了一次。

// sync包还提供了一个atomic子包，支持对于一些基础数据类型的原子操作函数，比如经典的CAS函数：

// func CompareAndSwapUnit64(val *uint64, old, new uint64) (swapped bool)
