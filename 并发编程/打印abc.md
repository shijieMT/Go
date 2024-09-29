实习面试题
### 手撕过程
经典的打印abc100次，要求用chan和waitgroup实现
>  cat
dog
fish
cat
dog
fish
cat
dog
fish
.....
依次顺序打印这三个，各100次

想了想，我的思路如下：  
向chan中写1、2、3代表a、b、c，然后三个协程中循环读取，如果是自己对应的数字，打印出来，如果不是，再把数字放回chan中

代码如下:
```go
func printCat(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for {
		//fmt.Println("cat count: ", count)
		if count == 100 {
			break
		}
		value, ok := <-ch
		if !ok {
			continue
		} else if value == 1 {
			count++
			fmt.Println("cat")
		} else {
			ch <- value
		}
	}
}
func printDog(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for {
		//fmt.Println("dog count: ", count)
		if count == 100 {
			break
		}
		value, ok := <-ch
		if !ok {
			continue
		} else if value == 2 {
			count++
			fmt.Println("dog")
		} else {
			ch <- value
		}
	}
}
func printFish(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for {
		//fmt.Println("fish count: ", count)
		if count == 100 {
			break
		}
		value, ok := <-ch
		if !ok {
			continue
		} else if value == 3 {
			count++
			fmt.Println("fish")
		} else {
			ch <- value
		}
	}
}
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(3)
	ch <- 1
	go printCat(ch, &wg)
	go printDog(ch, &wg)
	go printFish(ch, &wg)
}
```
死锁，面试官让把chan缓冲区改为1
```go
func main() {
	ch := make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(3)
	ch <- 1
	go printCat(ch, &wg)
	go printDog(ch, &wg)
	go printFish(ch, &wg)
}
```
此时，没有报错，进程直接退出了，面试官提示加上wg.wait()
```go
func main() {
	ch := make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(3)
	ch <- 1
	go printCat(ch, &wg)
	go printDog(ch, &wg)
	go printFish(ch, &wg)
  wg.Wait()
}
```
添加完毕，运行又是死锁，只打印出来一个cat（当时没注意cat成功打印了），看了一会没头绪，手撕结束  
### 复盘过程
后续进行了复盘，结果发现在打印cat之后，没有把2放入chan中，三个打印函数都没处理（当时写的时候手忙脚乱，还没写完就开始运行，然后把“有还没写的逻辑”这事给忘记了）  
修正代码：
```go
func printCat(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for {
		//fmt.Println("cat count: ", count)
		if count == 100 {
			break
		}
		value, ok := <-ch
		if !ok {
			continue
		} else if value == 1 {
			count++
			fmt.Println("cat")
			ch <- 2
		} else {
			ch <- value
		}
	}
}
func printDog(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for {
		//fmt.Println("dog count: ", count)
		if count == 100 {
			break
		}
		value, ok := <-ch
		if !ok {
			continue
		} else if value == 2 {
			count++
			fmt.Println("dog")
			ch <- 3
		} else {
			ch <- value
		}
	}
}
func printFish(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for {
		//fmt.Println("fish count: ", count)
		if count == 100 {
			break
		}
		value, ok := <-ch
		if !ok {
			continue
		} else if value == 3 {
			count++
			fmt.Println("fish")
			ch <- 1
		} else {
			ch <- value
		}
	}
}
func main() {
	ch := make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(3)
	ch <- 1
	go printCat(ch, &wg)
	go printDog(ch, &wg)
	go printFish(ch, &wg)
	wg.Wait()
}
```
这时，已经可以正确打印了，我长舒一口气  
### 继续思考
不过事情还没有结束，在我看来，这个代码能跑起来，很可能是巧合  
按道理讲，依照我的思路，无缓channal完全可以做到正确打印  
结果我尝试将ch改为无缓，直接死锁，没有打印出任何东西  
继续思考
> 1. main函数中的ch <- 1不应该放在协程启动前面
> 2. 当printFish打印出最后一个fish后，不应该继续向ch写数据了

更改代码：
```go
func printFish(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for {
		//fmt.Println("fish count: ", count)
		value, ok := <-ch
		if !ok {
			continue
		} else if value == 3 {
			count++
			fmt.Println("fish")
			if count == 99 {
				return
			}
			ch <- 1
		} else {
			ch <- value
		}
	}
}
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(3)
	go printCat(ch, &wg)
	go printDog(ch, &wg)
	go printFish(ch, &wg)
	ch <- 1
	wg.Wait()
}
```
改完之后发现printCat和printDog死锁了（当然main肯定死锁了）
```go
func printCat(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for {
		//fmt.Println("cat count: ", count)
		value, ok := <-ch
		if !ok {
			continue
		} else if value == 1 {
			count++
			fmt.Println("cat")
			ch <- 2
			if count == 99 {
				return
			}
		} else {
			ch <- value
		}
	}
}
func printDog(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for {
		//fmt.Println("dog count: ", count)
		value, ok := <-ch
		if !ok {
			continue
		} else if value == 2 {
			count++
			fmt.Println("dog")
			ch <- 3
			if count == 99 {
				return
			}
		} else {
			ch <- value
		}
	}
}
func printFish(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for {
		//fmt.Println("fish count: ", count)
		value, ok := <-ch
		if !ok {
			continue
		} else if value == 3 {
			count++
			fmt.Println("fish")
			if count == 99 {
				return
			}
			ch <- 1
		} else {
			ch <- value
		}
	}
}
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(3)
	go printCat(ch, &wg)
	go printDog(ch, &wg)
	go printFish(ch, &wg)
	ch <- 1
	wg.Wait()
}
```
cat和doge在ch插入后终止，fish在ch插入前终止，改完后终于成功打印了，此时发现之前的count == 100的判断条件完全没有用上
### 用多个channal的方式
听面试官说到“用了一个channal”，看来这个题目的通常解法是使用多个channal，简单思考一下

