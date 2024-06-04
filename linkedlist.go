// Linked List is a list with node that contains the next node information if it has
// ex:
// node1 as Head           node2
// {                       {
//  data: 1,                data: 2,
//  next: node2,            next: nil,
// }                       {

// First node is called head which is also a type of node
// Linked List type contains the head & the length
// When we add a next node in head & every next add iteration will increase the leangth
// l.length++

package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

// Linked List module
func LL() {
	newList := LinkedList{}
	newList.prepend(&node{
		data: 10,
	})
	newList.prepend(&node{
		data: 30,
	})
	newList.prepend(&node{
		data: 45,
	})

	newList.prepend(&node{
		data: 90,
	})

	err := newList.deleteWithValue(46)
	if err != nil {
		fmt.Println(err)
	}
	newList.printNodes()

}

func (l *LinkedList) prepend(n *node) {
	next := l.head
	l.head = n
	l.head.next = next
	l.length++
}

func (l *LinkedList) deleteWithValue(v int) error {

	// If l.head.data == v
	// We will search if the head value is v
	// We replace the head with l.head.next

	//If l.head.data != v
	// search until the next data is not equal to v
	// when found then replace the next with the next to next
	// ex: suppose nodex is the node that has a next node that conatains the value
	// means: nodex.next.data == v
	// we will replace the next with nodex.next.next cuz every next is a node
	// and every node has a next that represents another node

	if l.length == 0 {
		return fmt.Errorf("[X] Linked List is empty")
	}
	prev := l.head
	if prev.data == v {
		l.head = prev.next
		l.length--
		return nil
	}

	for prev.next.data != v {
		fmt.Printf("Current data: %d\n", prev.data)
		fmt.Printf("Next to Current data: %d\n", prev.next.data)
		fmt.Println()
		if prev.next.next == nil {
			return fmt.Errorf("Value doesn't exists in linked list")
		}
		prev = prev.next
	}

	prev.next = prev.next.next
	l.length--

	return nil
}

func (l LinkedList) printNodes() {
	var currentNode *node
	currentNode = l.head

	fmt.Println("Lenght of linked list: ", l.length)

	if l.length != 0 {
		for l.length != 0 {
			fmt.Printf("Node %d, Value: %d\n", l.length, currentNode.data)
			fmt.Println("-----------------------------------------")
			currentNode = currentNode.next
			l.length--
		}
	}
}
