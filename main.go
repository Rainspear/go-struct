package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64 = 0

var mu sync.Mutex

func count(wg *sync.WaitGroup) {
	mu.Lock()
	atomic.AddInt64(&counter, 1)
	fmt.Println("Counter:", atomic.LoadInt64(&counter))
	mu.Unlock()
	wg.Done()
}

func main() {
	// lh := learnHttp{}
	// li := learnInterface{}
	// ls := learnStruct{}
	// lm := learnMap{}
	// lc := learnConcurrency{}
	// lg := learnGroup{}
	// lr := learnRace{}
	// rtc := learnRangeChannel{}
	// lsc := learnSelect{}
	// lqc := learnQuitChannel{}
	// lfic := learnFanIn{}
	lfoc := learnFanOut{}
	// lh.executeMain()
	// li.executeMain()
	// ls.executeMain()
	// lm.executeMain()
	// lc.executeMain()
	// lg.executeMain()
	// lr.executeMain()
	// rtc.executeMain()
	// lqc.executeMain()
	// lfic.executeMain()
	lfoc.executeMain()
}
