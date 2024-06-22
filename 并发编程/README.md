2024.6.22
美图二面后，我才发现我对go的并发编程一无所知

面试时，面试官给了这样一道题目（问我这应该打印什么）：

```go
package dir

import (
	"fmt"
	"sync"
	"testing"
)

func TestName(t *testing.T) {
	s := sync.WaitGroup{}
	list := []string{"a", "b", "c", "d", "e", "f"}
	ch := make(chan bool)
	for _, v := range list {
		s.Add(1)
		go func() {
			fmt.Println(v)
			ch <- true
			s.Done()
		}()
	}
	for _ = range list {
		<-ch
	}
	s.Wait()
}
```

我先问了问面试官，sync.WaitGroup{}是做什么的，然后又问了go func() 是什么  
思考了一会后，我给出了我的答案：打印 a  
给面试官解释为什么打印a的时候，我又说打印 a，b  
然后，面试结束了
再然后，就有了这个文件夹hhh~  

