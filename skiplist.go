package skiplist

type Rand interface {
	Float64() float64
}

type Node[T comparable] struct {
	successors []Node[T]
	key        T
	val        any
}

type List[T comparable] struct {
	heads    []Node[T]
	maxLevel int
}

func New[T comparable]() *List[T] {
	return &List[T]{
		heads:    make([]Node[T], 0),
		maxLevel: 0,
	}
}

func (s *List[T]) Insert(key T, value any) {

}

func (s *List[T]) Search(key T) (*Node[T], error) {
	return nil, nil
}

func randomLevel(rand Rand, p float64) (level int) {
	for rand.Float64() <= p {
		level += 1
	}
	return
}
