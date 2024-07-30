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

## trees

### tree traversal

- preorder
  1. visit parent node
    - root gets visited first
  2. visit node on left subtree then visit the rest of the left subtree
  3. visit node on right subtree then visit the rest of the right subtree
    - rightmost node gets visited last
- inorder
  1. visit left subtree in inorder
    - leftmost leaf gets visited first
  2. visit parent node
    - root gets visited in the middle
  3. visit right subtree in inorder
    - rightmost node gets visited last
- postorder
  1. visit left subtree in postorder
    - leftmost leaf gets visited first
  2. visit right subtree in postorder (on the same level as previous step)
  3. visit parent node
    - root gets visited last


### Red-black tree

- Summary
  - Tree that self-balances to reduce time complexity of basic operations (insert, remove, etc)
- Rotation
  - x and y nodes are defined by ordering (x < y)
  - x and y have parent-child relationship and switch relationship when rotating
  - the input of the operation is the parent node (not the pivot node)
    - the pivot node is the initial child node
    - after rotating, the pivot node becomes the parent
  - keeps all the nodes and subtrees in the correct order
