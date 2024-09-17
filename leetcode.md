
# LeetCode

## General Tips

- No need to implement
  - `map`
  - `set`

- May need to implement
  - `linked list`

## Specific Problems and Tips

### 88 Merge Sorted Array
- When it's hard to manipulate an array, sometimes you can reverse the order in which you manipulate it

### Two Sum
- Use `map` to easily search elements
- Define the complement then search it
- If you want to get the complement but don't want to get the first occurrence, you can create a map where you only add an element after you find that the complement is equal to the number
- Example: `[1, 2, 2, 3]` target is 4
  - 1 in map | 2, 2, 3 remain

### IsPalindrome
- Convert number to str
- If the number is negative, it's not a palindrome
- For `i in length//2`:
  - You don't need to pass through all characters

### ContainsDuplicate
- Create `set`
- Check if there's duplicate in set
- If not, add the current element to set

### LengthLastWord
- One line Python solution
- Use `list[-1]`

### LongestCommonPrefix
- If it's a common prefix, restrict the length of the loop
  - The maximum length of the prefix is the length of the shortest word

### 202 Happy Number
- If you want to separate the digits of a number in a list, you can convert the number to a string and `.append(int(char))` into the list

```python
n = 223 => digits = [2, 2, 3]
```

### 219 Contains Duplicate II
- Only store what is necessary
  - If you need to check only the last difference during iteration, don't save all the differences in a list

### 141 Linked List Cycle
- Turtle and hare solution
  - If you create 2 pointers: 1 fast (hare), 1 slow (turtle), and you are looking for a circle, the hare eventually will reach/be the turtle

### 21 Merge Two Sorted Lists
- If you don't wish to backtrack to get to the first element of a linked list, save it at the start of the program
- Don't forget the ternary operator

```python
a = b if a > b else c
```

### 103 Maximum Depth of Binary Tree
- If not node: return depth

```python
max(recursive_function(), recursive_function())
```

### 100 Same Tree
- Recursion is a way to call future iterations of DFS without having to create a data structure to store memory of all iterations

### 226 Invert Binary Tree
- In Python, you don't need a buffer to switch 2 variables' values

```python
a, b = b, a
```

### 101 Symmetric Tree
- Identify points where the solution should change behavior
  - That way, you can define the functions that will compute each part of the solution
- Pay attention to the return type of the function; it normally helps you start building the logic behind a recursive solution

### 222 Count Complete Tree Nodes
- The base cases tell you how to start the function:
  - What happens when input is empty?
  - What happens when input is one element?
  - The recursion happens when?
  - The function's input format is a hint on how to solve the problem
