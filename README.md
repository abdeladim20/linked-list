# linked-list

# linkedList1 : 
Insert Function: Adds a new node to the end of the list.
Delete Function: Removes the first node with the specified value from the list.
Search Function: Returns true if a node with the given value is found, otherwise false.
Display Function: Prints the entire list in a readable format.
###
This example demonstrates how to insert, delete, and search within a linked list in Go. You can further extend this implementation to add more features, like inserting at the beginning, inserting after a specific node, or deleting the entire list.

# Doubly Linked List : 
This example demonstrates how to implement a doubly linked list in Go, with functionalities for inserting, deleting, and searching nodes, as well as displaying the list in both directions.
### 
A doubly linked list is similar to a singly linked list, but each node has two pointers: one pointing to the next node and another pointing to the previous node. This allows traversal in both directions.
###
Insert Function: Adds a new node to the end of the list. It updates both the prev and next pointers.
Delete Function: Removes the first node with the specified value. It adjusts both the next and prev pointers based on the node's position.
Search Function: Returns true if a node with the given value is found, otherwise false.
DisplayForward & DisplayBackward Functions: Print the list from head to tail and tail to head, respectively.