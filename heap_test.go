package heap

import (
	"log"
	"math/rand"
	// "strconv"
	"fmt"
	"testing"
	"time"
)

type IntKey struct {
	key   int
	value int
}

var h Heap

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

func TestHeap(t *testing.T) {

	defer logElapsedTime(time.Now(), "TestHeap")

	done := make(chan struct{})
	heap := NewHeap(done, trials)
	h = heap

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

}

func TestRandomHeap(t *testing.T) {

	defer logElapsedTime(time.Now(), "TestRandomHeap")

	rand.Seed(time.Now().UnixNano())

	done := make(chan struct{})
	heap := NewHeap(done, 0)
	h = heap

	for i := 0; i < trials; i++ {
		j := rand.Intn(trials)
		heap.Push(&IntKey{key: j, value: j})
	}

	curr := heap.Pop().(*IntKey)
	for !heap.IsEmpty() {
		next := heap.Pop().(*IntKey)
		if curr.key < next.key {
			t.Errorf("out of order, curr.key: %d, next.key: %d", curr.key, next.key)
		}
		curr = next
	}

}
