package heap

import (
	"log"
	"math/rand"
	"sync"
	// "strconv"
	"fmt"
	"testing"
	"time"
)

type IntKey struct {
	key   int
	value int
}

func (this *IntKey) CompareTo(other Key) int {
	if other == nil {
		panic(fmt.Sprintf("nil key, this: %t, other: %t", this, other))
	}
	otherK := other.(*IntKey)
	if this.key == otherK.key {
		return 0
	} else if this.key < otherK.key {
		return -1
	} else {
		return 1
	}
}

const (
	trials int = 1000000
)

func logElapsedTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func _TestHeap(t *testing.T) {

	defer logElapsedTime(time.Now(), "TestHeap")

	done := make(chan struct{})
	heap := NewHeap(done, trials)

	for i := 0; i < trials; i++ {
		heap.Push(&IntKey{key: i, value: i})
	}

	curr := heap.Pop().(*IntKey)
	for !heap.IsEmpty() {
		next := heap.Pop().(*IntKey)
		if curr.key < next.key {
			t.Errorf("out of order, curr.key: %d, next.key: %d", curr.key, next.key)
		}
		curr = next
	}

	close(done)

}

func _TestRandomHeap(t *testing.T) {

	defer logElapsedTime(time.Now(), "TestRandomHeap")

	rand.Seed(time.Now().UnixNano())

	done := make(chan struct{})
	heap := NewHeap(done, 0)

	for i := 0; i < trials; i++ {
		j := rand.Intn(trials)
		heap.Push(&IntKey{key: j, value: j})
	}

	curr := heap.Pop().(*IntKey)
	for !heap.IsEmpty() {
		next := heap.Pop().(*IntKey)
		//log.Println(next)
		if curr.key < next.key {
			t.Errorf("out of order, curr.key: %d, next.key: %d", curr.key, next.key)
		}
		curr = next
	}

	close(done)

}

func TestRandomHeapConcurrent(t *testing.T) {

	defer logElapsedTime(time.Now(), "TestRandomHeapConcurrent")

	rand.Seed(time.Now().UnixNano())

	done := make(chan struct{})
	heap := NewHeap(done, 0)

	//go func() {
	for i := 0; i < trials; i++ {
		j := rand.Intn(trials)
		heap.Push(&IntKey{key: j, value: j})
	}
	//}()

	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			curr := heap.Pop().(*IntKey)
			for !heap.IsEmpty() {
				_next := heap.Pop()
				if _next != nil {
					next := _next.(*IntKey)
					//log.Println(j, next)
					if curr.key < next.key {
						t.Errorf("out of order, curr.key: %d, next.key: %d", curr.key, next.key)
					}
					curr = next
				}
			}
		}(i)
	}
	wg.Wait()
	close(done)
}

func TestRandomHeapConcurrentWithPopCh(t *testing.T) {

	defer logElapsedTime(time.Now(), "TestRandomHeapConcurrentWithPopCh")

	rand.Seed(time.Now().UnixNano())

	done := make(chan struct{})
	heap := NewHeap(done, 0)

	//go func() {
	for i := 0; i < trials; i++ {
		j := rand.Intn(trials)
		heap.Push(&IntKey{key: j, value: j})
	}
	//}(

	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(j int, pop chan<- chan Key) {
			defer wg.Done()
			key := make(chan Key)
			pop <- key
			curr := <-key
			for curr != nil {
				pop <- key
				next := <-key
				if next != nil {
					//log.Println(j, next)
					if curr.(*IntKey).key < next.(*IntKey).key {
						t.Errorf("out of order, curr.key: %s, next.key: %s", curr, next)
					}

				}
				curr = next
			}
		}(i, heap.PopCh())
	}
	wg.Wait()
	close(done)
}
