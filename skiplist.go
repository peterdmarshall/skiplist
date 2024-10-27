package skiplist

import (
	"errors"
	"fmt"
)

type Key[T any] interface {
	LessThan(k T) bool
	Equal(k T) bool
}

type Rand interface {
	Float64() float64
}

type Node[T Key[T]] struct {
	successors []*Node[T]
	Key        Key[T]
	Value      any
}

type List[T Key[T]] struct {
	header   *Node[T]
	maxLevel int
	rand     Rand
	p        float64
}

func New[T Key[T]](rand Rand, p float64) *List[T] {
	return &List[T]{
		header: &Node[T]{
			successors: make([]*Node[T], 1),
		},
		maxLevel: 0,
		rand:     rand,
		p:        p,
	}
}

func (s *List[T]) Insert(key T, value any) {
	level := s.randomLevel()

	if level > s.maxLevel {
		for range level - s.maxLevel {
			s.header.successors = append(s.header.successors, (*Node[T])(nil))
		}
		s.maxLevel = level
	}

	n := &Node[T]{
		Key:        key,
		Value:      value,
		successors: make([]*Node[T], s.maxLevel+1),
	}

	for i := s.maxLevel; i >= 0; i-- {
		prev := s.header.successors[i]
		if prev == nil {
			s.header.successors[i] = n
		} else {
			curr := prev.successors[i]
			for curr != nil && curr.Key.LessThan(key) {
				prev = curr
				curr = curr.successors[i]
			}
			prev.successors[i] = n
			n.successors[i] = curr
		}
	}
}

func (s *List[T]) Search(key T) (*Node[T], error) {
	curr := s.header
	for i := s.maxLevel; i >= 0; i-- {
		for curr.successors[i] != nil && curr.successors[i].Key.LessThan(key) {
			curr = curr.successors[i]
		}
	}

	if curr.successors[0] != nil && curr.successors[0].Key.Equal(key) {
		return curr.successors[0], nil
	} else {
		return nil, errors.New("not found")
	}
}

func (s *List[T]) Print() {
	for i := range len(s.header.successors) {
		level := len(s.header.successors) - i - 1
		fmt.Printf("Level %d: ", level)
		curr := s.header.successors[level]
		for curr != nil {
			fmt.Printf("|%v|%v| ", curr.Key, curr.Value)
			curr = curr.successors[level]
		}
		fmt.Println("")
	}
}

func (s *List[T]) randomLevel() int {
	return randomLevel(s.rand, s.p)
}

func randomLevel(rand Rand, p float64) (level int) {
	for rand.Float64() <= p {
		level += 1
	}
	return
}
