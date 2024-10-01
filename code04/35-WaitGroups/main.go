package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	//このWaitGroupは、ここで起動したすべてのゴルーチンが完了するのを待つために使われる。
	// WaitGroup を関数に引数として渡す場合は、ポインタで渡す必要がある
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
}

// Worker 5 starting
// Worker 3 starting
// Worker 2 starting
// Worker 1 starting
// Worker 4 starting
// Worker 4 done
// Worker 5 done
// Worker 1 done
// Worker 3 done
// Worker 2 done
