
# Algorithms

![Big O complexity chart](./images/big_o_complexity_chart.jpg)

| notation   | name    |
|--------------- | --------------- |
| O(1)   | constant   |
| O(log n)   | logarithm   |
| O(n)   | linear   |
| O(n * log n)   | ---   |
| O(n²) | quadratic |
| O(n³) | cubic |
| O(n^c) | polinomial |
| O(c^n) | exponential |
| O(n!) | factorial |

- adaptive: faster for partially sorted data sets
- stable: does not change the relative order of elements with equal keys
- in-place: only requires a constant amount of memory
- online: can sort a list as it receives it


## bubble sort

- summary
  - switch pair of elements n² times so that the bigger element goes to the right
  - easy to implement
  - slow
  - low memory usage
- time complexity
  - worst case: O(n²)
  - average case: O(n²)
  - best case: O(n)
- space complexity
  - O(1)

```python
def bubblesort(nums):
    for j in range(len(nums)):
        for i in range(len(nums)-1-j):
            print(nums)
            if nums[i] > nums[i+1]:
                buffer = nums[i+1]
                nums[i+1] = nums[i]
                nums[i] = buffer
    return nums
```


## selection sort

- summary
  - slow
  - OBS: SeLectiOn sort => SLOw
  - OBS: Select Smallest and Swap
  - in-place
  - not stable
  - simple implementation
- pseudocode
  - split values into sorted and unsorted
    - sorted section begins empty
  - select smallest value in all unsorted section
  - swap smallest value with rightmost/last position of sorted section
  - repeat
- time complexity
  - worst case: O(n²)
  - average case: O(n²)
  - best case: O(n²)
- space complexity
  - O(1)


## insertion sort

- summary
  - OBS: (insert) unsorted element into correct position
  - slow
  - simple implementation, easy to write
  - faster than other simple sorting algorithms like bubble sort
  - adaptive: faster for partially sorted data sets
  - stable: does not change the relative order of elements with equal keys
  - in-place: only requires a constant amount of memory
  - online: can sort a list as it receives it
- pseudocode
  - separate the array into 2 parts: sorted and unsorted section
  - pick the first element of the unsorted and insert it into correct position in sorted section
  - put the unsorted element in the correct position in the sorted section
  - sort the next element of the sorted section
- time complexity
  - worst case: O(n²)
  - average case: O(n²)
  - best case: O(n²)
- space complexity
  - O(1)

```python
def insertion_sort(nums):
    for i in range(len(nums)):
        j = i
        while j > 0 and nums[j - 1] > nums[j]:
            nums[j], nums[j - 1] = nums[j - 1], nums[j]
            j -= 1
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
  - best case: O(n * log n)
  - average case: O(n * log n)
  - worst case: O(n²)
- space complexity
  - O(log(n))

```python

def quicksort(nums):
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
    return quicksort(less) + [pivot] + quicksort(greater)

```

## merge sort

- summary
  - fast
  - uses a lot of memory
  - divide and conquer
- pseudocode
  - splits unsorted array recursively in two
  - after splitting, it merges 2 arrays by creating a third and iterating through the 2 initial arrays, comparing the elements and adding the smallest
- time complexity
  - best case: O(n * log n)
  - average case: O(n * log n)
  - worst case: O(n * log n)
- space complexity
  - O(n)


## shell sort

- summary
  - improvement of Insertion sort
  - gap
- pseudocode
- time complexity
  - best case: O(n * log n)
  - average case: depende do gap
  - worst case: O(n²)


## heap sort

- summary
  - uses heap data structure
- pseudocode
- time complexity
  - best case: O(n * log n)
  - average case: O(n * log n)
  - worst case: O(n * log n)


# data search

## binary search

  - sort array
  - compare middle value with target value
  - if middle is equal to target, stop and return current position
  - if middle is bigger than target, discard the half with the bigger numbers and binary search the remaining array
  - if middle is smaller than target, discard the half with the smaller numbers and binary search the remaining array
