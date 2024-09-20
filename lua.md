
## basics

- dynamically typed
- high-level
- designed for embedded use in applications
  - complete language is small
- easy to read and understand
- blocks are NOT defined by indentation
- garbage collection
- lua is NOT object oriented
- immutable strings
- parentheses in `if` statements are optional
- can also be used for functional programming

## data types

- `nil`: represents null or no value
- `boolean`: true or false
- `number`
- `string`
- `table`: works like an array, dictionary or object

```lua

x = 10            -- number
y = "Hello"       -- string
is_active = true  -- boolean
z = nil           -- nil

fruits = {"apple", "banana", "cherry"}
print(fruits[1])  -- Output: apple

person = {name = "John", age = 30}
print(person["name"])  -- Output: John

```

## string operations

```lua

str1 = "Hello"
str2 = "World"
result = str1 .. " " .. str2  -- Concatenation with a space
print(result)  -- Output: Hello World

str = "Hello"
-- calculates string length
print(#str)  -- Output: 5

-- s: The string.
-- i: The starting index (1-based).
-- j: The ending index (optional). If not specified, it extracts up to the end.
string.sub(s, i, j)

```

## control flow statements

### conditionals

```lua

x = 15

if x > 20 then
  print("x is greater than 20")
elseif x > 10 then
  print("x is greater than 10")
else
  print("x is 10 or less")
end

```

### while

```lua
i = 0

while i < 5 do
  print(i)
  i = i + 1
end
```

### loops

```lua

i = 0
while i < 5 do
  print(i)
  i = i + 1
end

for i = 1, 5 do
  print(i)
end

```

## functions

```lua

function greet(name)
  return "Hello, " .. name
end

print(greet("Alice"))  -- Output: Hello, Alice

```

## modules

> used to organize code into reusable libraries

creating a module:
```lua

-- mymodule.lua
local M = {}

function M.sayHello()
  print("Hello from module!")
end

return M

```

using a module:
```lua

local mymodule = require("mymodule")
mymodule.sayHello()  -- Output: Hello from module!

```
