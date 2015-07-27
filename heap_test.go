package heap

import (
	"log"
	// "math"
	// "strconv"
	"testing"
	// "time"
)

type IntKey struct {
	key   int
	value int
}

func (this *IntKey) CompareTo(other Key) int {
	otherK := other.(*IntKey)
	if this.key == otherK.key {
		return 0
	} else if this.key < otherK.key {
		return -1
	} else {
		return 1
	}
}

func TestHeap(t *testing.T) {

	heap := NewHeap(1000)

	for i := 0; i < 1000; i++ {
		heap.Push(&IntKey{key: i, value: i})
	}

	for !heap.IsEmpty() {
		key := heap.Pop()
		log.Println(key)
	}

}
