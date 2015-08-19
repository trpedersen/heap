package heap

//	"log"

import "fmt"

// from Algorithms 4th Ed., by Sedgewick

// Key interface - you'll need to implement the CompareTo method.
// CompareTo returns -1, 0 +1 if <, =, > to other
type Key interface {
	CompareTo(other Key) int
}

// Heap interface
type Heap interface {
	Push(v Key)
	Pop() Key
	PopCh() chan<- chan Key
	IsEmpty() bool
	Size() int
	Keys() []Key
}

type heap struct {
	keys []Key
	done <-chan struct{}
	push chan Key
	pop  chan chan Key
}

func dump(heap *heap) {
	for i, key := range heap.keys {
		fmt.Printf("i: %d, key: %s\n", i, key)
	}
	fmt.Printf("N: %d\n", len(heap.keys))
}

// NewHeap returns a new heap with an initialCapacity
func NewHeap(done <-chan struct{}, initialCapacity int) Heap {
	heap := &heap{
		keys: make([]Key, 0, initialCapacity),
		//	N:    0,
		push: make(chan Key),
		pop:  make(chan chan Key),
		done: done,
	}
	go heap.run()
	return heap
}

// func (heap *heap) Halt() error {
// 	select {
// 	case <-heap.done: // already closed
// 		return nil
// 	default:
// 		close(heap.done)
// 	}
// 	return nil
// }

func (heap *heap) run() {

run:
	for {
		select {
		case <-heap.done:
			// log.Println("done")
			break run
		case key := <-heap.push:
			// log.Println("push")
			heap._push(key)
		case ch := <-heap.pop:
			// log.Println("pop")
			heap._pop(ch)
		}
	}
	close(heap.pop)
	close(heap.push)
	return
}

func (heap *heap) Push(key Key) {
	if key == nil {
		panic("Push: nil key")
	}
	heap.push <- key
}

func (heap *heap) _push(key Key) {
	if key == nil {
		panic("nil key _push")
	}
	heap.keys = append(heap.keys, key)
	//log.Printf("heap.keys: %s\n", heap.keys)
	heap.swim(len(heap.keys) - 1)
}

func (heap *heap) Pop() Key {
	ch := make(chan Key)
	heap.pop <- ch
	return <-ch
}

func (heap *heap) PopCh() chan<- chan Key {
	return heap.pop
}

func (heap *heap) _pop(ch chan Key) {

	if heap.IsEmpty() {
		ch <- nil
	} else {

		key := heap.keys[0]                         // retrieve key from top, order depends on Key.CompareTo implementation
		heap.exchange(0, len(heap.keys)-1)          // exchange with last item
		heap.keys = heap.keys[0 : len(heap.keys)-1] // avoid loitering
		heap.sink(0)
		ch <- key
	}

}

func (heap *heap) IsEmpty() bool {
	return len(heap.keys) == 0
}

func (heap *heap) Size() int {
	return len(heap.keys)
}

func (heap *heap) Keys() []Key {
	return heap.keys
}

func (heap *heap) less(i, j int) bool {
	// log.Printf("a: %s, i: %d, j: %d", a, i, j)
	if heap.keys[i] == nil {
		dump(heap)
		panic(fmt.Sprintf("i: %d", i))
	}
	if heap.keys[j] == nil {
		dump(heap)
		panic(fmt.Sprintf("j: %d", j))
	}
	return heap.keys[i].CompareTo(heap.keys[j]) < 0
}

func (heap *heap) exchange(i, j int) {
	// t := [i]
	// a[i] = a[j]
	// a[j] = t
	heap.keys[i], heap.keys[j] = heap.keys[j], heap.keys[i]
}

// bubble k up the tree unit it is not less than its parent
// or until it is at the root
func (heap *heap) swim(k int) {
	for k > 0 && heap.less((k-1)/2, k) {
		heap.exchange((k-1)/2, k)
		k = (k - 1) / 2
	}
}

// send down the tree until it is not greater than any of its children
// or unit it is at the end
func (heap *heap) sink(k int) {
	N := len(heap.keys) - 1
	for (2*k + 1) <= N {
		j := (2*k + 1)                  // lh child
		if j < N && heap.less(j, j+1) { // get the greatest child
			j++
		}
		if !heap.less(k, j) {
			break // no child is greater
		}
		heap.exchange(k, j) // greatest child becomes parent
		k = j               // keep sinking
	}
}
