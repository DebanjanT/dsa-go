//  Stack linear data structure
// Last In  First Out
// [0,1,2,3] -> POP -> [0,1,2]
package main

import "fmt"

type stack struct {
	data []int
}

func (s *stack) Push(v int) {
	// ex: [0] [0,1]
	s.data = append(s.data, v)
}

func (s *stack) Pop() int {
	// ex: [0] [0,1]
    PopElement := s.data[len(s.data)-1]
	PopLength := len(s.data)-1
	split := s.data[:PopLength]
    s.data = split

	fmt.Println(split)
    fmt.Println("Popped Element: ", PopElement)

    return PopElement
}

func GoStack() {
	s := stack{}
    s.Push(10)
    s.Push(20)
    s.Push(40)
    s.Push(30)
    s.Push(60)
    s.Push(90)
    fmt.Println(s)
    s.Pop()
    s.Pop()
    fmt.Println(s)
}

