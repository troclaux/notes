
# data structures

- when adding or removing elements, update the data structure's size accordingly
- when adding elements, update the references on the previous element
- when removing elements
  - update the references on the next element
  - check if data structure is empty and raise error if necessary

## linked list

> linear data structure where each element stores a reference of the address of the next element

### Singly linked list

| element1 | => | element2 | => | element3 | => | element4 |

### doubly linked list

| element1 | <=> | element2 | <=> | element3 | <=> | element5 |

### circular linked list

| element1 | => | element2 | => | element3 | => | element4 | => | element1 |

## queue

- FIFO: First In First Out

basic methods:
- queue.push(item): adds an item to the tail of the queue (index 0 of list)
- queue.pop(): removes and returns an item from the head of the queue (last index of list)
- queue.peek(): returns an item from the head of the queue
- queue.size(): returns the number of items in the queue

### priority queue

- retrieves elements based on their priority
- elements are ordered based on their priority
- highest priority element is removed first

## stack

- LIFO: Last In First Out

- `push()`/`append()`: add element to top of stack
- `pop()`: remove element from top of stack

## hash table

- hash function: maps keys to values
- hash collision: when two keys map to the same value
  - can be resolved by chaining (linked list) or open addressing (find next available slot)

## graph

- node or vertex
- edge: connection between two nodes in a graph can be directed or undirected
- adjacent nodes: two nodes with an edge directly connecting them
- incident edge
- degree: number of edges that enter or leave a vertex
- path: a sequence of edges that leads from one node to another
- cycle: path where the starting node is the same as the ending node, and no other nodes are repeated
- order: the total number of vertices in the graph

<!--adjacency matrix-->

## trees

- height: the number of elements from the root to the deepest branch
  - if a tree has only one element, its height is 1
- level: the 'floor' of a node
  - the root has level 0
  - direct children of the root have level 1
  - grandchildren of the root have level 2, and so on
- degree: the number of children each node can have
  - binary trees have a maximum degree of 2


TODO
vertex removal:
- leaf node: remove the node
- node with one child: remove the node and connect the child to the parent of the removed node
- node with two children: remove the node and replace it with the smallest node in the right subtree

#### tree traversal order

> [!TIP]
> the name of each traversal order describes when the root is visited

- preorder: root is visited first
- inorder: root is visited in the middle-ish
- postorder: root is visited last

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


### binary tree

> tree where each node has at most 2 children

- children don't follow any particular order based on the value of the node

#### complete binary tree

binary tree where:
- all levels are completely filled except possibly for the last level
- all nodes are as far left as possible

#### heap

> binary tree-based data structure that satisfies the heap property and also is a complete binary tree

- heap property: parent node is either greater than or less than its children
  - max heap: parent's value is greater than children
  - min heap: parent's value is less than children

key methods:
- `insert()`: add a new element to the heap
- `delete()`: remove the root element from the heap
- `peek()`: return the root element of the heap
- `size()`: return the number of elements in the heap
- `heapify()`: convert an array into a heap

#### Binary Search Tree (BST)

binary tree where:
- left child is smaller than parent
- right child is bigger than parent


#### AVL tree

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

#### red-black tree

- summary
  - tree that self-balances to reduce time complexity of basic operations (insert, remove, etc)
- rotation
  - x and y nodes are defined by ordering (x < y)
  - x and y have parent-child relationship and switch relationship when rotating
  - the input of the operation is the parent node (not the pivot node)
    - the pivot node is the initial child node
    - after rotating, the pivot node becomes the parent
  - keeps all the nodes and subtrees in the correct order


#### AVL vs red-black trees

- insertion, deletion and loop-up are guaranteed O(n * log(n))
- because AVL is balanced, lookups in AVL are faster than red-black
- red-black trees are better for insert intensive tasks
