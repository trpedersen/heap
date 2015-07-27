package heap

import (
	"log"
	"math/rand"
	// "strconv"
	"testing"
	"time"
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

const (
	trials int = 1000000
)

func logElapsedTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func TestHeap(t *testing.T) {

	defer logElapsedTime(time.Now(), "TestHeap")

	heap := NewHeap(trials)

	for i := 0; i < trials; i++ {
		heap.Push(&IntKey{key: i, value: i})
	}

	for !heap.IsEmpty() {
		//key :=
		heap.Pop()
		//log.Println(key)
	}

}

func TestRandomHeap(t *testing.T) {

	defer logElapsedTime(time.Now(), "TestRandomHeap")

	rand.Seed(time.Now().UnixNano())

	heap := NewHeap(trials)

	for i := 0; i < trials; i++ {
		j := rand.Intn(trials)
		heap.Push(&IntKey{key: j, value: j})
	}

	for !heap.IsEmpty() {
		//key :=
		heap.Pop()
		//log.Println(key)
	}

}
