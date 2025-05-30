
# data structures

> organized way to store, manage and retrieve data efficiently in a computer so that it can be used effectively

- remember when implementing data structures:
  - when adding or removing elements, update the data structure's size/length accordingly
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

- basic operations:
  - `queue.push(item)`: adds an item to the tail of the queue (index 0 of list)
  - `queue.pop()`: removes and returns an item from the head of the queue (last index of list)
  - `queue.peek()`: returns an item from the head of the queue
  - `queue.size()`: returns the number of items in the queue

### priority queue

> data structure where each element has a priority value that determines its processing order

- elements with higher priority are processed before elements with lower priority
- typically implemented using a heap data structure

- basic operations:
  - `enqueue(element, priority)`: add element with given priority
  - `dequeue()`: remove and return highest-priority element
  - `peek()`: view highest-priority element without removing it
  - `is_empty()`: check if queue is empty

## stack

- LIFO: Last In First Out

- basic operations:
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
- adjacency matrix: 2D array representation of a graph
  - size: n x n where n is number of vertices
  - `matrix[i][j]` = 1 if there is an edge from vertex i to j
  - `matrix[i][j]` = 0 if there is no edge
  - for undirected graphs, matrix is symmetric

## trees

> hierarchical data structure composed by nodes, where each node has a value and may have child nodes

- node: stores data and link to children if they exist
  - root: topmost node of a tree
    - has no parent
  - leaf: a node without children
  - parent: node that has a child node
  - children: node that has a parent node
- edge: connection between two nodes (parent to child)
- subtree: part of a tree, starting at a node and including its descendants
- height of a tree: the number of edges between the root to the deepest leaf
  - if a tree has only one element, its height is 1
- depth of a node: number of edges from the root to the node
- level: the 'floor' of a node
  - the root has level 0
  - direct children of the root have level 1
  - grandchildren of the root have level 2, and so on
- degree: the number of children a node has
  - binary trees have a maximum degree of 2

### vertex removal pseudocode

- if vertex has no children:
  - remove vertex
  - update parent's reference to null
- if vertex has one child:
  - connect parent to child
  - remove vertex
- if vertex has two children:
  - find successor (smallest value in right subtree)
  - copy successor's value to vertex
  - remove successor

### tree traversal order

> [!TIP]
> the name of each traversal order describes when the root is visited

- pre-order: root is visited first
- in-order: root is visited in the middle-ish
- post-order: root is visited last

- pre-order
  1. visit parent node
    - root gets visited first
  2. visit node on left subtree then visit the rest of the left subtree
  3. visit node on right subtree then visit the rest of the right subtree
    - rightmost node gets visited last
- in-order
  1. visit left subtree in in-order
    - leftmost leaf gets visited first
  2. visit parent node
    - root gets visited in the middle
  3. visit right subtree in in-order
    - rightmost node gets visited last
- post-order
  1. visit left subtree in post-order
    - leftmost leaf gets visited first
  2. visit right subtree in post-order (on the same level as previous step)
  3. visit parent node
    - root gets visited last

### binary tree

> tree where each node has at most 2 children

- children don't follow any particular order based on the value of the node

#### complete binary tree

- binary tree where:
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

### B-tree

> self-balancing search tree

- efficiently manages large amounts of sorted data stored on disk
- commonly used in databases and file systems

properties of B-tree of order _m_:

- each node can have at most m children
