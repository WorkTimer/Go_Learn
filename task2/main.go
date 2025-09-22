package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	fmt.Println("=== 增加10 ===")
	num := 5
	increaseByTen(&num)
	fmt.Printf("输入: %d -> 输出: %d\n", num, num)

	fmt.Println("=== 乘以2 ===")
	nums := []int{1, 2, 3, 4, 5}
	doubleSliceElements(&nums)
	fmt.Printf("输入: %v -> 输出: %v\n", nums, nums)

	fmt.Println("\n=== 协程并发打印 ===")
	printNumbersWithGoroutines()

	fmt.Println("\n=== 任务调度器 ===")
	taskScheduler(func() {
		fmt.Println("任务1")
	}, func() {
		fmt.Println("任务2")
	})

	fmt.Println("\n=== Shape 接口 ===")
	shapes := []Shape{
		Rectangle{3, 4},
		Circle{5},
	}
	for i, s := range shapes {
		fmt.Printf("形状%d: 面积=%.2f, 周长=%.2f\n", i+1, s.Area(), s.Perimeter())
	}

	fmt.Println("\n=== Employee 结构体 ===")
	employee := Employee{
		Person:     Person{Name: "张三", Age: 30},
		EmployeeID: "123456",
	}
	employee.PrintInfo()

	fmt.Println("\n=== 通道通信 ===")
	channelCommunication()

	fmt.Println("\n=== 带有缓冲的通道 ===")
	bufferedChannel()

	fmt.Println("\n=== Mutex 计数器 ===")
	mutexCounter()

	fmt.Println("\n=== 原子操作计数器 ===")
	atomicCounter()
}

// 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值
// 考察点 ：指针的使用、值传递与引用传递的区别。
func increaseByTen(ptr *int) {
	*ptr += 10
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
func doubleSliceElements(slicePtr *[]int) {
	for i := 0; i < len(*slicePtr); i++ {
		(*slicePtr)[i] *= 2
	}
}

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
func printNumbersWithGoroutines() {
	var wg sync.WaitGroup

	wg.Go(func() {
		for i := 1; i <= 10; i += 2 {
			fmt.Printf("奇数: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	})

	wg.Go(func() {
		for i := 2; i <= 10; i += 2 {
			fmt.Printf("偶数: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	})

	wg.Wait()
	fmt.Println("所有协程执行完毕")
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
func taskScheduler(tasks ...func()) {
	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Go(func() {
			start := time.Now()
			task()
			duration := time.Since(start)
			fmt.Printf("任务执行时间: %v\n", duration)
		})
	}

	wg.Wait()
}

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct{ w, h float64 }

func (r Rectangle) Area() float64      { return r.w * r.h }
func (r Rectangle) Perimeter() float64 { return 2 * (r.w + r.h) }

type Circle struct{ r float64 }

func (c Circle) Area() float64      { return 3.14 * c.r * c.r }
func (c Circle) Perimeter() float64 { return 2 * 3.14 * c.r }

// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
//考察点 ：组合的使用、方法接收者。
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工ID: %s, 姓名: %s, 年龄: %d\n", e.EmployeeID, e.Name, e.Age)
}

// 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
func channelCommunication() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Go(func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	})

	wg.Go(func() {
		for n := range ch {
			fmt.Printf("接收: %d\n", n)
		}
	})

	wg.Wait()
}

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。
func bufferedChannel() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Go(func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	})

	wg.Go(func() {
		for n := range ch {
			fmt.Printf("%d ", n)
		}
		fmt.Println()
	})

	wg.Wait()
}

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
func mutexCounter() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Go(func() {
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		})
	}
	wg.Wait()
	fmt.Printf("最终计数: %d\n", counter)
}

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。
func atomicCounter() {
	var counter atomic.Int64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Go(func() {
			for j := 0; j < 1000; j++ {
				counter.Add(1)
			}
		})
	}

	wg.Wait()
	fmt.Printf("最终计数: %d\n", counter.Load())
}
