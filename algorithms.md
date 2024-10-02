
# algorithms

## big O notation

![Big O complexity chart](./images/big_o_complexity_chart.jpg)

[Big O cheat sheet](https://www.bigocheatsheet.com/)

[Big O notation explained](https://yourbasic.org/algorithms/big-o-notation-explained/)

| notation   | name    |
|--------------- | --------------- |
| O(1) | constant |
| O(log n) | logarithmic |
| O(n) | linear |
| O(n * log n) | --- |
| O(n²) | quadratic |
| O(n³) | cubic |
| O(n^c) | polinomial |
| O(c^n) | exponential |
| O(n!) | factorial |

- adaptive: faster for partially sorted data sets
- stable: does not change the relative order of elements with equal keys
- in-place: algorithm that operates on its input data structure without needing an auxiliary data structure
- online: can sort a list as it receives it
- space complexity: measure of the amount of memory used by an algorithm as the input size grows
- time complexity: measure of the runtime of an algorithm as the input size grows
  - best case => Ω()
  - average case => ϴ()
  - worst case => O()

## techniques

[dynamic programming](https://yourbasic.org/algorithms/dynamic-programming-explained/)

- dynamic programming algorithm: solves a complex problem by dividing it into simpler subproblems, solving each of those just once, and storing their solutions
- memoization: optimization technique used to speed up programs by storing the results of expensive function calls and returning the cached result when the same inputs occur again
- tabulation: approach where you solve a dynamic programming problem by first filling up a table, and then compute the solution to the original problem based on the results in this table

## sorting algorithms

### bubble sort

- summary
  - switch pair of elements n² times so that the bigger element goes to the right
  - easy to implement
  - slow
  - low memory usage
- time complexity
  - best case: Ω(n)
  - average case ϴ(n²)
  - worst case: O(n²)
- space complexity
  - O(1)

```python
def bubble_sort(nums):
    for j in range(len(nums)):
        for i in range(len(nums)-1-j):
            if nums[i] > nums[i+1]:
                nums[i], nums[i+1] = nums[i+1], nums[i]
    return nums
```


### selection sort

- summary
  - slow
  - OBS: SeLectiOn sort => SLOw
  - OBS: SELECTION Sort => SELECT Smallest and Swap
  - find smallest value and swap with rightmost value in sorted section
  - in-place
  - not stable
  - simple implementation
- pseudocode
  - split values into sorted and unsorted
    - sorted section begins empty
  - find smallest value in all unsorted section
  - swap smallest value with rightmost value of sorted section
  - repeat until array is sorted
- time complexity
  - best case: Ω(n²)
  - average case ϴ(n²)
  - worst case: O(n²)
- space complexity
  - O(1)

```python
def selection_sort(nums):
    if len(nums) <= 1:
        return nums
    for i in range(len(nums)):
        min_idx = i
        for j in range(i + 1, len(nums)):
            if nums[j] < nums[min_idx]:
                min_idx = j
        nums[i], nums[min_idx] = nums[min_idx], nums[i]
    return nums
```

### insertion sort

- summary
  - OBS: (insert) unsorted element into correct position
  - slow
  - simple implementation, easy to write
  - get first unsorted element, then insert it in correct position in sorted section
  - faster than other simple sorting algorithms like bubble sort
  - adaptive: faster for partially sorted data sets
  - stable: does not change the relative order of elements with equal keys
  - in-place: only requires a constant amount of memory
  - online: can sort a list as it receives it
- pseudocode
  - iterate over the array and position the current element in the correct position
  - separate the array into 2 parts: sorted and unsorted section
  - pick the first element of the unsorted and insert it into correct position in sorted section
  - put the unsorted element in the correct position in the sorted section
  - sort the next element of the sorted section
- time complexity
  - best case: Ω(n²)
  - average case ϴ(n²)
  - worst case: O(n²)
- space complexity
  - O(1)

```python
def insertion_sort(nums):
    for i in range(1, len(nums)):
        key = nums[i]
        j = i - 1
        while j >= 0 and key < nums[j]:
            nums[j + 1] = nums[j]
            j -= 1
        nums[j + 1] = key
    return nums
```

## quick sort

- summary
  - divide and conquer
  - pick a pivot element and partition the array around the pivot
  - recursively sort the sub-arrays
    - all elements less than the pivot are on the left
    - all elements greater than the pivot are on the right
    - the pivot is in the correct position
  - in-place
  - fast (but not always)
  - high memory usage
- pseudocode
  - select an element that will be the pivot
  - move all elements that are smaller than the pivot on the left of the pivot
  - move all elements that are greater than the pivot on the right of the pivot
  - run quicksort on the left segment of the array
  - run quicksort on the right segment of the array
- time complexity
  - best case: Ω(n * log n)
  - average case ϴ(n * log n)
  - worst case: O(n²)
- space complexity
  - O(log(n))

```python
def quick_sort(nums):
    if len(nums) <= 1:
        return nums
    pivot = nums[0]
    less = []
    greater = []
    for i in nums[1:]:
        if i > pivot:
            greater.append(i)
        else:
            less.append(i)
    return quick_sort(less) + [pivot] + quick_sort(greater)
```

### merge sort

- summary
  - fast
  - uses a lot of memory
  - divide and conquer
- pseudocode
  - splits unsorted array recursively in two
  - after splitting, it merges 2 arrays by creating a third and iterating through the 2 initial arrays, comparing the elements and adding the smallest
- time complexity
  - best case: Ω(n * log n)
  - average case ϴ(n * log n)
  - worst case: O(n * log n)
- space complexity
  - O(n)

```python
def merge_sort(arr):
    if len(arr) <= 1:
        return arr
    mid = len(arr) // 2
    left = arr[:mid]
    right = arr[mid:]
    return merge(merge_sort(left), merge_sort(right))


def merge(arr1, arr2):
    res = []
    i = 0
    j = 0
    while i < len(arr1) and j < len(arr2):
        if arr1[i] < arr2[j]:
            res.append(arr1[i])
            i += 1
        else:
            res.append(arr2[j])
            j += 1
    res.extend(arr1[i:])
    res.extend(arr2[j:])
    return res
```

### heap sort

- summary
  - uses heap data structure
- pseudocode
- time complexity
  - best case: Ω(n * log n)
  - average case ϴ(n * log n)
  - worst case: O(n * log n)


## search algorithms

### binary search

- efficient search on sorted arrays
- compare middle value with target value
- if middle is equal to target, stop and return current position
- if middle is bigger than target, discard the half with the bigger numbers and binary search the remaining array
- if middle is smaller than target, discard the half with the smaller numbers and binary search the remaining array

### depth-first search (DFS)

- summary
  - explore as far as possible along each branch before backtracking
  - uses a stack
  - recursive implementation


### breadth-first search (BFS)

- summary
  - explore all neighbor nodes at the present depth prior to moving on to the nodes at the next depth level
  - uses a queue
  - iterative implementation

## minimum spanning tree algorithms

### kruskal's algorithm

- summary
  - find the minimum spanning tree in a graph
  - greedy algorithm
  - works with positive and negative weights

### prim's algorithm

- summary
  - find the minimum spanning tree in a graph
  - greedy algorithm
  - works with positive weights

## shortest path algorithms

### dijkstra's algorithm

- summary
  - find the shortest path between source node and all other nodes in a graph
  - greedy algorithm
  - works with positive weights
- time complexity
  - O(V*E)
    - V: number of vertices
    - E: number of edges

### bellman-ford algorithm

- summary
  - find the shortest path between nodes in a graph
  - works with negative weights
  - slower than dijkstra's algorithm
- time complexity
  - O(V²)
    - V: number of vertices

### a* algorithm

- summary
  - find the shortest path between nodes in a graph
  - uses heuristics to find the best path
  - uses a priority queue

