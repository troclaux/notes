
# C#

## properties

- object-oriented programming language
- statically typed language
- strongly typed language
- compiled language
- has garbage collection
- function-scoped variables
- no hoisted declarations
- immutable strings
- each statement must end with semicolon (`;`)

- created by microsoft
- designed for the .NET platform
- comment with `//` and `/* */`

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

- `using System;`: imports basic classes from standard library (e.g. `Console.WriteLine("hello");`)
- `static void Main()`: entry point of c# program, where the execution starts
    - `static`: no need to create an object to call this method (belongs to the class)
    - `void`: this method returns nothing

- a c# program is composed by 3 parts:
    - namespaces: used to organize code and avoid naming conflicts
    - classes
    - methods


```csharp
using System;

namespace MyApp {
    class Program {
        static void Main(string[] args) {
            Console.WriteLine("Hello, world!");
        }
    }
}
```

## data types

- `int`: integer numbers
- `double`: floating point numbers
- `bool`: true or false
- `char`: a single character
- `string`: text
- `object`: base type of all types in c#
- `array`: fixed-size sequence of elements
- `decimal`: high-precision number (e.g. money), requires you to type `m` after the number
- `var`: implicitly typed (type inferred at compile time)

```csharp
int age = 25;
float height = 1.75f;
double weight = 70.5;
decimal price = 19.99m;
bool isStudent = true;
char grade = 'A';
string name = "Alice";
object data = "some value";

int[] numbers = new int[3];                    // Array of 3 integers (default values 0)
int[] vetor = { 1, 2, 3 };
string[] names = { "Alice", "Bob", "Carol" };  // Array with 3 strings

var score = 95;                                // Inferred as int
int? year = null;                              // Nullable int
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

## string operations

- declaring strings: `string name = "Alice";`
- concatenation: `string fullName = "Alice" + " " + "Smith";`
- interpolation: `string greeting = $"Hello, {name}!";`
- get string length: `int len = name.Length;`
- get char from string: `char firstLetter = name[0]; // 'A'`
- get substring: `string part = name.Substring(1, 3); // "lic"`
- convert to lowercase/uppercase:
    - `string upper = name.ToUpper(); // "ALICE"`
    - `string lower = name.ToLower(); // "alice"`
- contains: `bool hasA = name.Contains("A");       // true`

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

```csharp
string[] names = { "Alice", "Bob", "Charlie" };

foreach (string name in names)
{
    Console.WriteLine(name);
}
```

## classes and objects

`:` == inherits from

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

```csharp
class Animal  // Base class
{
    public void Eat()
    {
        Console.WriteLine("This animal eats food.");
    }
}

class Dog : Animal  // Derived class
{
    public void Bark()
    {
        Console.WriteLine("The dog barks.");
    }
}

class Program
{
    static void Main()
    {
        Dog myDog = new Dog();
        myDog.Eat();   // Inherited from Animal
        myDog.Bark();  // Defined in Dog
    }
}
```

## collections

- `List<T>`

## keywords

- `public`: no restrictions
- `private`: only within the same class
- `protected`: in the same class or any child class
- `internal`: in the same assembly/project only
- `protected internal`: in the same assembly or any derived class, even if it's in another assembly
- `private protected`: in the same class or derived classes but only within the same assembly

---

## running c# on the cli

1. write code in a file that ends with `.cs` that has a `Program` class (e.g. `Program.cs`)
1. create `.csproj` file: `dotnet new console -o .`
1. build: `dotnet build`
1. run: `dotnet run`

