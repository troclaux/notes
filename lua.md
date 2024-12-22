
# lua

## basics

- dynamically typed
- high-level
- designed for embedded use in applications
  - complete language is small
- easy to read and understand
- blocks are NOT defined by indentation
- garbage collection
- functions are values (i.e. first-class citizens)
- doesn't have built-in class-based OOP, but supports object-oriented programming through prototype-based approach:
  - tables and metatables
  - can implement objects, inheritance, and polymorphism
- immutable strings
- parentheses in `if` statements are optional
- can also be used for functional programming

## data types

- `nil`: represents null or no value
- `boolean`: true or false
- `number`
- `string`
- `table`: used for arrays, dictionaries and objects

```lua
x = 10            -- number
y = "Hello"       -- string
is_active = true  -- boolean
z = nil           -- nil

fruits = {"apple", "banana", "cherry"}
print(fruits[1])  -- Output: apple

person = {name = "John", age = 30}
local person = {
    name = "John",
    greet = function(self)
        print("Hello, my name is " .. self.name)
    end
}

print(person["name"])  -- Output: John

-- Method call
person:greet()  -- Output: Hello, my name is John
```

## string manipulation

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

-- removes leading and trailing whitespaces
str = str:gsub('^%s+', ''):gsub('%s+$', '')
```

- `string.gsub(input_string, pattern, replacement)`
  - `input_string`: the string to modify
  - `pattern`: the pattern to search for
  - `replacement`: the string to replace the pattern with

## if else

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

### loops

```lua
i = 0
while i < 5 do
  print(i)
  i = i + 1
end
```

```lua
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
