package heap

//	"log"

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
	IsEmpty() bool
	Size() int
	Keys() []Key
}

type heap struct {
	keys []Key
	N    int
}

// NewHeap returns a new heap with an initialCapacity
func NewHeap(initialCapacity int) Heap {
	heap := &heap{
		keys: make([]Key, 0, initialCapacity),
		N:    0,
	}
	return heap
}

func (heap *heap) Push(v Key) {
	heap.N++
	heap.keys = append(heap.keys, v)
	//log.Printf("heap.keys: %s\n", heap.keys)
	swim(heap.keys, heap.N-1)
}

func (heap *heap) Pop() Key {
	if heap.IsEmpty() {
		return nil
	}
	key := heap.keys[0]              // retrieve key from top, order depends on Key.CompareTo implementation
	exchange(heap.keys, 0, heap.N-1) // exchange with last item
	heap.N--
	heap.keys[heap.N] = nil // avoid loitering
	sink(heap.keys, 0, heap.N-1)
	return key
}

func (heap *heap) IsEmpty() bool {
	return heap.N <= 0
}

func (heap *heap) Size() int {
	return heap.N
}

func (heap *heap) Keys() []Key {
	return nil
}

func less(a []Key, i, j int) bool {
	//log.Printf("a: %s, i: %d, j: %d", a, i, j)
	return a[i].CompareTo(a[j]) < 0
}

func exchange(a []Key, i, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}

// bubble k up the tree unit it is not less than its parent
// or until it is at the root
func swim(a []Key, k int) {
	for k > 0 && less(a, (k-1)/2, k) {
		exchange(a, (k-1)/2, k)
		k = (k - 1) / 2
	}
}

// send down the tree until it is not greater than any of its children
// or unit it is at the end
func sink(a []Key, k, N int) {
	for (2*k + 1) <= N {
		j := (2*k + 1)                // lh child
		if j < N && less(a, j, j+1) { // get the greatest child
			j++
		}
		if !less(a, k, j) {
			break // no child is greater
		}
		exchange(a, k, j) // greatest child becomes parent
		k = j             // keep sinking
	}
}
