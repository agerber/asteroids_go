package prime

import (
	"sync"
)

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	head  *Node
	tail  *Node
	count int        // Keeps track of the number of elements in the list
	lock  sync.Mutex // Using Mutex for thread safety
}

func NewLinkedList() *LinkedList {
	return &LinkedList{lock: sync.Mutex{}}
}

func (l *LinkedList) Clear() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.head = nil
	l.tail = nil
	l.count = 0
}

// Effectively enqueuing to the tail
func (l *LinkedList) Add(data interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	newNode := &Node{Data: data}
	if l.head == nil {
		l.tail = newNode
		l.head = newNode
	} else {
		l.tail.Next = newNode
		l.tail = newNode
	}
	l.count++
}

// Iterator for the linked list
func (l *LinkedList) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		l.lock.Lock()
		defer l.lock.Unlock()
		current := l.head
		for current != nil {
			ch <- current.Data
			current = current.Next
		}
	}()
	return ch
}

// Enqueue (same as Add)
func (l *LinkedList) Enqueue(data interface{}) {
	l.Add(data)
}

func (l *LinkedList) Dequeue() interface{} {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.count == 0 {
		return nil
	} else if l.count == 1 {
		temp := l.head
		l.head = nil
		l.tail = nil
		l.count--
		return temp.Data
	} else {
		temp := l.head
		l.head = l.head.Next
		l.count--
		return temp.Data
	}
}

func (l *LinkedList) Remove(data interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	current := l.head
	previous := (*Node)(nil)
	for current != nil {
		if current.Data == data {
			if previous != nil {
				previous.Next = current.Next
				if current == l.tail {
					l.tail = previous
				}
			} else {
				l.head = current.Next
				if current == l.tail {
					l.tail = nil
				}
			}
			l.count--
			return
		}
		previous = current
		current = current.Next
	}
}

func (l *LinkedList) PrintList() {
	l.lock.Lock()
	defer l.lock.Unlock()
	current := l.head
	for current != nil {
		print(current.Data, " -> ")
		current = current.Next
	}
	println("None")
}

// Len returns the number of elements in the linked list.
func (l *LinkedList) Len() int {
	// Leverage the existing count field for efficiency
	return l.count
}
