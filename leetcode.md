
# leetcode

## general tips

- read the problem carefully
- clarify any doubts
  - what is the type of the input?
- pseudocode => brute-force => optimize
  - what you are trying to do?
- handle edge cases
- when using recursion, handle base cases
- read documentation
- recognize common problem-solving patterns:
  - sliding window: use two pointers
  - fast-slow pointers
- recognize when to use algorithms
  - divide and conquer
  - dynamic programming
  - sorting
  - recursion
- ask LLM
  - use reasoning models for complex problems
- when greedy solution fails, recursion with memoization is a powerful fallback to explore all paths and find a global solution

## debugging

- pay attention to the function signature
  - what is the output?
  - what is the input?
  - what is the purpose of each input?
  - are you using the correct types?

## techniques

### divide and conquer

> break a complex problem into smaller sub-problems, solve each independently, and combine the solutions to solve the original problem

### dynamic programming

> break a problem into smaller sub-problems, solve each only once, and store the solutions to sub-problems to avoid redundant computation and optimize the solution

- typical example: fibonacci sequence
- normally uses a matrix to store solutions

### memoization

> store results of expensive function calls and returns cached result when the same input occurs again

- `@lru_cache(None)`: built-in python decorator that caches the results of function calls with a specific argument

```python
from functools import lru_cache

@lru_cache(maxsize=None)
def fibonacci(n):
    if n <= 1:
        return n
    return fibonacci(n-1) + fibonacci(n-2)
```

without `lru_cache`:

```python
def fibonacci(n, memo = {}):
    if n in memo:
        return memo[n]
    if n == 0:
        return 0
    elif n == 1:
        return 1
    else:
        result = fibonacci(n-1, memo) + fibonacci(n-2, memo)
        memo[n] = result
        return result
```

### knapsack problem

TODO

## Specific Problems and Tips

### 1 Two Sum

- use `map` to easily search elements
- define the complement then search it
- if you want to get the complement but don't want to get the first occurrence, you can create a map where you only add an element after you find that the complement is equal to the number
- example: `[1, 2, 2, 3]` target is 4
  - 1 in map | 2, 2, 3 remain

### 3 longest substring without repeating characters

TODO

### 14 Longest Common Prefix

- if it's a common prefix, restrict the length of the loop
  - the maximum length of the prefix is the length of the shortest word

### 21 merge two sorted lists

- If you don't wish to backtrack to get to the first element of a linked list, save it at the start of the program
- Don't forget the ternary operator

```python
a = b if a > b else c
```

### 22 remove element

- to replace elements in-place, track index positions with an offset variable
- beware off-by-1 errors

### 49 group anagrams

TODO

### 58 length of last word

- One line Python solution
- Use `list[-1]`

### 88 Merge Sorted Array

- when it's hard to manipulate an array, consider reversing the order in which you manipulate it

### 100 Same Tree

- recursion is a way to call future iterations of DFS without having to create a data structure to store memory of all iterations

### 101 Symmetric Tree

- identify points where the solution should change behavior
  - that way, you can define the functions that will compute each part of the solution
- pay attention to the return type of the function
  - it normally helps you start building the logic behind a recursive solution

### 103 Maximum Depth of Binary Tree

- If not node: return depth

```python
max(recursive_function(), recursive_function())
```

### 125 valid palindrome

- Convert number to str
- If the number is negative, it's not a palindrome
- For `i in length//2`:
  - You don't need to pass through all characters

### 136 single number

- linear time complexity
- constant extra space
  - the memory used can't grow with the size of input
  - fixed extra memory used
  - not allowed: arrays/dictionaries/sets that grow with the input
- bitwise operations can be useful
- the solution uses xor
  - 0 XOR num = num
  - num XOR num = 0
  - 8000 >> 3 = 1000
    - 8000/(2^3)
  - 1000 << 3 = 8000
    - 1000 * (2^3)

### 141 linked list cycle

- turtle and hare solution
  - If you create 2 pointers: 1 fast (hare), 1 slow (turtle), and you are looking for a circle, the hare eventually will reach/be the turtle

### 217 contains duplicate

- create `set`
- check if there's duplicate in set
- if not, add the current element to set

### 202 happy number

- if you want to separate the digits of a number in a list, you can convert the number to a string and `.append(int(char))` into the list

```python
n = 223 => digits = [2, 2, 3]
```

### 219 contains duplicate ii

- only store what is necessary
  - if you need to check only the last difference during iteration, don't save all the differences in a list

### 222 Count Complete Tree Nodes

- the base cases tell you how to start the function:
  - what happens when input is empty?
  - what happens when input is one element?
  - the recursion happens when?
  - the function's input format is a hint on how to solve the problem

### 226 Invert Binary Tree

- in python, you don't need a buffer to switch 2 variables' values

```python
a, b = b, a
```
