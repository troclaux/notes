
# Algorithms

![Big O complexity chart](./images/big_o_complexity_chart.jpg)

## Bubble Sort
- Summary
  - Switch pair of elements n² times so that the bigger element goes to the right
  - Easy to implement
  - Slow
- Time complexity
  - O(n²)

## Quick Sort
- Summary
  - Not always fast (worst case is n²)
- Pseudocode
  - Select an element that will be the pivot
  - Move all elements that are smaller than the pivot on the left of the pivot
  - Move all elements that are greater than the pivot on the right of the pivot
  - Run quicksort on the left segment of the array
  - Run quicksort on the right segment of the array
- Time complexity
  - Average: O(n * log n)
  - Worst case: O(n²)

## Merge Sort
- Summary
  - Fast
  - Uses a lot of memory
  - Divide and conquer
- Pseudocode
  - Splits unsorted array recursively in two
  - After splitting, it merges 2 arrays by creating a third and iterating through the 2 initial arrays, comparing the elements and adding the smallest
- Time complexity
  - Average: O(n * log n)

## Insertion Sort
- Summary
  - Slow
  - Simple implementation, easy to write
  - Fast for very small data sets
  - Faster than other simple sorting algorithms like Bubble Sort
  - Adaptive: Faster for partially sorted data sets
  - Stable: Does not change the relative order of elements with equal keys
  - In-Place: Only requires a constant amount of memory
  - Online: Can sort a list as it receives it
- Pseudocode
  - Separate the array into 2 parts: sorted and unsorted section
  - Pick the first element of the unsorted and compare with the last of the sorted
  - Put the unsorted element in the correct position in the sorted section
  - Pick the next element of the sorted section

# Data Structures

## Linked List

## Queue
