package main

import "fmt"

// Node represents a single element in a linked list
type Node struct {
    data int
    next *Node
}

// LinkedList represents the linked list
type LinkedList struct {
    head *Node
}

// Insert adds a new node to the end of the list
func (list *LinkedList) Insert(value int) {
    newNode := &Node{data: value}
    if list.head == nil {
        list.head = newNode
    } else {
        current := list.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
}

// Delete removes the first node with the given value
func (list *LinkedList) Delete(value int) {
    if list.head == nil {
        fmt.Println("List is empty.")
        return
    }

    // If the node to be deleted is the head
    if list.head.data == value {
        list.head = list.head.next
        return
    }

    current := list.head
    for current.next != nil && current.next.data != value {
        current = current.next
    }

    // If the node was found, unlink it
    if current.next != nil {
        current.next = current.next.next
    } else {
        fmt.Println("Value not found in the list.")
    }
}

// Search returns true if a node with the given value exists
func (list *LinkedList) Search(value int) bool {
    current := list.head
    for current != nil {
        if current.data == value {
            return true
        }
        current = current.next
    }
    return false
}

// Display prints all the nodes in the list
func (list *LinkedList) Display() {
    current := list.head
    for current != nil {
        fmt.Print(current.data, " -> ")
        current = current.next
    }
    fmt.Println("nil")
}

func main() {
    list := LinkedList{}

    list.Insert(10)
    list.Insert(20)
    list.Insert(30)
    list.Insert(40)

    fmt.Println("Original List:")
    list.Display() // Output: 10 -> 20 -> 30 -> 40 -> nil

    fmt.Println("After deleting 20:")
    list.Delete(20)
    list.Display() // Output: 10 -> 30 -> 40 -> nil

    fmt.Println("Searching for 30:", list.Search(30)) // Output: true
    fmt.Println("Searching for 50:", list.Search(50)) // Output: false
}

