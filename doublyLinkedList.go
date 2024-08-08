package main

import "fmt"

// Node represents a single element in a doubly linked list
type Node struct {
    data int
    prev *Node
    next *Node
}

// DoublyLinkedList represents the doubly linked list
type DoublyLinkedList struct {
    head *Node
    tail *Node
}

// Insert adds a new node to the end of the list
func (list *DoublyLinkedList) Insert(value int) {
    newNode := &Node{data: value}
    if list.head == nil {
        list.head = newNode
        list.tail = newNode
    } else {
        newNode.prev = list.tail
        list.tail.next = newNode
        list.tail = newNode
    }
}

// Delete removes the first node with the given value
func (list *DoublyLinkedList) Delete(value int) {
    if list.head == nil {
        fmt.Println("List is empty.")
        return
    }

    current := list.head
    for current != nil && current.data != value {
        current = current.next
    }

    if current == nil {
        fmt.Println("Value not found in the list.")
        return
    }

    // If the node to be deleted is the head
    if current == list.head {
        list.head = current.next
        if list.head != nil {
            list.head.prev = nil
        } else {
            list.tail = nil // List became empty
        }
    } else if current == list.tail {
        // If the node to be deleted is the tail
        list.tail = current.prev
        list.tail.next = nil
    } else {
        // Node is in the middle
        current.prev.next = current.next
        current.next.prev = current.prev
    }
}

// Search returns true if a node with the given value exists
func (list *DoublyLinkedList) Search(value int) bool {
    current := list.head
    for current != nil {
        if current.data == value {
            return true
        }
        current = current.next
    }
    return false
}

// DisplayForward prints all the nodes from head to tail
func (list *DoublyLinkedList) DisplayForward() {
    current := list.head
    for current != nil {
        fmt.Print(current.data, " -> ")
        current = current.next
    }
    fmt.Println("nil")
}

// DisplayBackward prints all the nodes from tail to head
func (list *DoublyLinkedList) DisplayBackward() {
    current := list.tail
    for current != nil {
        fmt.Print(current.data, " -> ")
        current = current.prev
    }
    fmt.Println("nil")
}

func main() {
    list := DoublyLinkedList{}

    list.Insert(10)
    list.Insert(20)
    list.Insert(30)
    list.Insert(40)

    fmt.Println("Original List (Forward):")
    list.DisplayForward() // Output: 10 -> 20 -> 30 -> 40 -> nil

    fmt.Println("Original List (Backward):")
    list.DisplayBackward() // Output: 40 -> 30 -> 20 -> 10 -> nil

    fmt.Println("After deleting 20:")
    list.Delete(20)
    list.DisplayForward() // Output: 10 -> 30 -> 40 -> nil

    fmt.Println("Searching for 30:", list.Search(30)) // Output: true
    fmt.Println("Searching for 50:", list.Search(50)) // Output: false
}

