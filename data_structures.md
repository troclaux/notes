# Data structures

## Linked List

## Doubly Linked List

## Queue

## Boot Queue

- queue.push(item)
  - Adds an item to the tail of the queue (index 0 of list)
- queue.pop()
  - Removes and returns an item from the head of the queue (last index of list)
- queue.peek()
  - Returns an item from the head of the queue
- queue.size()
  - Returns the number of items in the queue

## Red-black tree

- Summary
  - Tree that self-balances to reduce time complexity of basic operations (insert, remove, etc)
- Rotation
  - x and y nodes are defined by ordering (x < y)
  - x and y have parent-child relationship and switch relationship when rotating
  - the input of the operation is the parent node (not the pivot node)
    - the pivot node is the initial child node
    - after rotating, the pivot node becomes the parent
  - keeps all the nodes and subtrees in the correct order
