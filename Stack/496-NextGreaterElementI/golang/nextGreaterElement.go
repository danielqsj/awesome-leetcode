package nextGreaterElement

import (
	"sync"
)

type stack struct {
	lock sync.Mutex // you don't have to do this if you don't want thread safety
	s    []int
}

func NewStack() *stack {
	return &stack{sync.Mutex{}, make([]int, 0)}
}

func (s *stack) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *stack) Pop() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return 0
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res
}

func (s *stack) Peek() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return 0
	}

	res := s.s[l-1]
	return res
}

func (s *stack) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	return l == 0
}

func nextGreaterElement(findNums []int, nums []int) []int {
	stack := NewStack()
	result := make([]int, len(findNums))
	numsHash := make(map[int]int) // map element to nextGreaterElement
	for _, num := range nums {
		for !stack.IsEmpty() && stack.Peek() < num {
			numsHash[stack.Pop()] = num
		}
		stack.Push(num)
	}
	for i, num := range findNums {
		if v, ok := numsHash[num]; ok {
			result[i] = v
		} else {
			result[i] = -1
		}
	}
	return result
}
