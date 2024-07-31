# Linked List

# Doubly Linked List

# Queue

# Boot Queue

- queue.push(item)
  - Adds an item to the tail of the queue (index 0 of list)
- queue.pop()
  - Removes and returns an item from the head of the queue (last index of list)
- queue.peek()
  - Returns an item from the head of the queue
- queue.size()
  - Returns the number of items in the queue

# trees

## binary tree: each node has at most 2 children
## Binary Search Tree (BST)
  - left child is smaller than parent
  - right child is bigger than parent

## tree traversal

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

## AVL tree

> self-balancing binary search tree

```
if ( | right subtree height - left subtree height | > 1 ):
  balance the tree
```

the tree is balanced via 2 operations: left rotate and right rotate

```
T1, T2 and T3 are subtrees of the tree, rooted with y (on the left side) or x (on the right side)

     y                               x
    / \     Right Rotation          /  \
   x   T3   - - - - - - - >        T1   y
  / \       < - - - - - - -            / \
 T1  T2     Left Rotation            T2  T3
 
```

Keys in both of the above trees follow the following order
keys(T1) < key(x) < keys(T2) < key(y) < keys(T3)
So BST property is not violated anywhere



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


### AVL vs red-black trees

- insertion, deletion and loop-up are guaranteed O(n * log(n))
- because AVL is balanced, lookups in AVL are faster than red-black
- red-black trees are better for insert intensive tasks
