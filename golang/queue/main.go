package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Item interface{}

type CustomItem struct {
	Name  string
	Count int
}

func (ci *CustomItem) String() string {
	return fmt.Sprintf("%s: %d", ci.Name, ci.Count)
}

type RateLimitter struct {
	items []Item

	cond *sync.Cond
	wg   *sync.WaitGroup

	process func(Item)
}

func (rl *RateLimitter) Run() {
	rl.cond.L.Lock()
	defer rl.cond.L.Unlock()
	defer rl.wg.Done()

	time.Sleep(100 * time.Millisecond) // 1 処理辺り 100 ms なら最大でも 1 sec 10 になるやろうという雑な ...

	var item Item
	item, rl.items = rl.items[0], rl.items[1:]
	rl.process(item)
}

func (rl *RateLimitter) Add(item Item) {
	rl.cond.L.Lock()
	defer rl.cond.L.Unlock()

	rl.wg.Add(1)
	rl.items = append(rl.items, item)
	go rl.Run()

	rl.cond.Signal()
}

func main() {
	limitter := &RateLimitter{
		items: []Item{},
		cond:  sync.NewCond(&sync.Mutex{}),
		wg:    &sync.WaitGroup{},
		process: func(item Item) {
			cItem := item.(*CustomItem)
			fmt.Println("\t", cItem)
		},
	}

	processing := 0
	before := 0

	go func() {
		for {
			fmt.Println(time.Now().Format(time.RFC3339), "processing", processing-before)
			before = processing
			time.Sleep(time.Second)
		}
	}()

	cond := sync.NewCond(&sync.Mutex{})

	fooCount := 0
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		cond.L.Lock()
		defer cond.L.Unlock()

		fooCount++
		limitter.Add(Item(&CustomItem{
			Name:  "foo",
			Count: fooCount,
		}))
	})

	barCount := 0
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		cond.L.Lock()
		defer cond.L.Unlock()

		barCount++
		limitter.Add(Item(&CustomItem{
			Name:  "bar",
			Count: barCount,
		}))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
