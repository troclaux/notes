
# C#

## properties

- object-oriented programming language
- statically typed language
- compiled language
- has garbage collection
- function-scoped variables
- no hoisted declarations
- immutable strings

```csharp
using System;

class Program
{
    static void Main()
    {
        Console.WriteLine("Hello, World!");
    }
}
```

## data types

- `int`: integer numbers
- `float`, `double`: floating point numbers
- `bool`: true or false
- `string`: text
- `char`: a single character

```csharp
int age = 25;
string name = "Alice";
bool isStudent = true;
```

## if else

```csharp
if (age >= 18)
{
    Console.WriteLine("Adult");
}
else
{
    Console.WriteLine("Minor");
}
```

## loops

```csharp
for (int i = 0; i < 5; i++)
{
    Console.WriteLine(i);
}
```

## functions

```csharp
static int Add(int a, int b)
{
    return a + b;
}
```

## classes and objects

```csharp
class Dog
{
    public string Name;
    public void Bark()
    {
        Console.WriteLine($"{Name} says: Woof!");
    }
}

class Program
{
    static void Main()
    {
        Dog myDog = new Dog();
        myDog.Name = "Buddy";
        myDog.Bark();  // Buddy says: Woof!
    }
}
```

