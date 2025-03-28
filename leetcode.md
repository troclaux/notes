
# LeetCode

## General Tips

- read the problem carefully
- clarify any doubts
  - what is the type of the input?
- pseudocode => brute-force => optimize
  - what you are trying to do?
- handle edge cases
- when using recursion, handle base cases
- read documentation
- recognize common problem-solving patterns:
  - sliding window
  - two pointers
  - fast-slow pointers
- recognize when to use algorithms
  - divide and conquer
  - dynamic programming
  - sorting
  - recursion
- ask LLM
  - use reasoning models for complex problems

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

### memoization

> store results of expensive function calls and returns cached result when the same input occurs again

- can be used in fibonacci sequence

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

### longest increasing subsequence

## Specific Problems and Tips

### 1 Two Sum

- Use `map` to easily search elements
- Define the complement then search it
- If you want to get the complement but don't want to get the first occurrence, you can create a map where you only add an element after you find that the complement is equal to the number
- Example: `[1, 2, 2, 3]` target is 4
  - 1 in map | 2, 2, 3 remain

### 14 Longest Common Prefix

- If it's a common prefix, restrict the length of the loop
  - The maximum length of the prefix is the length of the shortest word

### 21 merge two sorted lists

- If you don't wish to backtrack to get to the first element of a linked list, save it at the start of the program
- Don't forget the ternary operator

```python
a = b if a > b else c
```

### 22 remove element

- to replace elements in-place, track index positions with an offset variable
- beware off-by-1 errors

### 58 length of last word

- One line Python solution
- Use `list[-1]`

### 88 Merge Sorted Array

- when it's hard to manipulate an array, consider reversing the order in which you manipulate it

### 100 Same Tree

- Recursion is a way to call future iterations of DFS without having to create a data structure to store memory of all iterations

### 101 Symmetric Tree

- Identify points where the solution should change behavior
  - That way, you can define the functions that will compute each part of the solution
- Pay attention to the return type of the function; it normally helps you start building the logic behind a recursive solution

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

### 141 linked list cycle

- turtle and hare solution
  - If you create 2 pointers: 1 fast (hare), 1 slow (turtle), and you are looking for a circle, the hare eventually will reach/be the turtle

### 217 contains duplicate

- create `set`
- check if there's duplicate in set
- if not, add the current element to set

### 202 happy number

- If you want to separate the digits of a number in a list, you can convert the number to a string and `.append(int(char))` into the list

```python
n = 223 => digits = [2, 2, 3]
```

### 219 contains duplicate ii

- Only store what is necessary
  - If you need to check only the last difference during iteration, don't save all the differences in a list

### 222 Count Complete Tree Nodes

- The base cases tell you how to start the function:
  - What happens when input is empty?
  - What happens when input is one element?
  - The recursion happens when?
  - The function's input format is a hint on how to solve the problem

### 226 Invert Binary Tree

- In Python, you don't need a buffer to switch 2 variables' values

```python
a, b = b, a
```
