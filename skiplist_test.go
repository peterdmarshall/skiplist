package skiplist

import (
	"github.com/stretchr/testify/require"
	"math/rand/v2"
	"testing"
)

type MockRand struct {
	idx      int
	sequence []float64
}

func (m *MockRand) Float64() float64 {
	res := m.sequence[m.idx%len(m.sequence)]
	m.idx += 1
	return res
}

var _ Rand = (*MockRand)(nil)

func NewMockSequenceRand(sequence []float64) Rand {
	return &MockRand{
		idx:      0,
		sequence: sequence,
	}
}

func Test_randomLevel(t *testing.T) {
	tests := []struct {
		description   string
		p             float64
		sequence      []float64
		expectedLevel int
	}{
		{
			description:   "with p=0.5, expect 2",
			p:             0.5,
			sequence:      []float64{0.23, 0.36, 0.6},
			expectedLevel: 2,
		},
		{
			description:   "with p=0.25, expect 0",
			p:             0.25,
			sequence:      []float64{0.8, 0.2, 0.4},
			expectedLevel: 0,
		},
		{
			description:   "with p=0.5, expect 5",
			p:             0.5,
			sequence:      []float64{0.2, 0.4, 0.5, 0.3, 0.4, 0.7},
			expectedLevel: 5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			level := randomLevel(NewMockSequenceRand(tc.sequence), tc.p)
			require.Equal(t, tc.expectedLevel, level)
		})
	}
}

type IntKey int

func (m IntKey) LessThan(k IntKey) bool {
	return m < k
}

func (m IntKey) Equal(k IntKey) bool {
	return m == k
}

func TestList_Insert(t *testing.T) {
	keys := []IntKey{1, 2, 3, 4}
	values := []string{"test1", "test2", "test3", "test4"}

	list := New[IntKey](rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())), 0.5)

	for i := range keys {
		list.Insert(keys[i], values[i])
	}
}

func TestList_Search(t *testing.T) {
	keys := []IntKey{1, 2, 3, 4}
	values := []string{"test1", "test2", "test3", "test4"}

	list := New[IntKey](rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())), 0.5)

	for i := range keys {
		list.Insert(keys[i], values[i])
	}

	res, err := list.Search(IntKey(3))
	require.NoError(t, err)
	require.Equal(t, res.Value, "test3")

	res, err = list.Search(IntKey(9))
	require.Error(t, err)
}
