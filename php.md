
# PHP

## properties

- dynamically typed
- weakly typed language
- interpreted programming language
- has garbage collection
- doesn't have hoisted declarations
- function-scoped variables
- mutable strings
- each statement must end with semicolon (`;`)

- primarily used for web development
- it is embedded in html
- PHP is acronym for "Php: Hypertext Preprocessor"
- uses `.php` file extension

## basic syntax

php code is written between `<?php` and `?>` (optional) tags:

```php
<?php
echo "Hello, world!";
?>
```

> [!NOTE]
> You don’t need to close the `<?php` tag in pure PHP files because:
> the PHP interpreter assumes the file ends with PHP code
> leaving it open avoids accidental whitespace or output after ?>, which can break things like headers or sessions

| Feature               | `echo`            | `print`          |
| --------------------- | ----------------- | ---------------- |
| Returns a value?      | ❌ No              | ✅ Yes (always 1) |
| Multiple arguments?   | ✅ Yes             | ❌ No             |
| Performance           | ✅ Slightly better | —                |
| Usable in expressions | ❌ No              | ✅ Yes            |

html + php integration example:

```html
<!DOCTYPE html>
<html>
<head>
    <title>PHP Example</title>
</head>
<body>

    <h1>Welcome!</h1>

    <?php
        $name = "Alice";
        echo "<p>Hello, $name!</p>";
    ?>

</body>
</html>
```

generating html dynamically:

```html
<ul>
<?php
$fruits = ["Apple", "Banana", "Cherry"];
foreach ($fruits as $fruit) {
    echo "<li>$fruit</li>";
}
?>
</ul>
```

## data types

- `integer`
- `float`
- `boolean`: case insensitive (e.g. both `true` and `TRUE` are valid)
- `string`
- `array`
- `object`
- `null`
- `resource`: special data type to represent external resources (things that are not native php data types)
    - e.g. files, database connections

> [!IMPORTANT]
> all variable names must start with `$`

> [!WARNING]
> variable names can't start with a number or `-`

```php
$age = 30;
$age++;
$price = 19.99;
$name = "Alice";
$isAdmin = true;
$isLoggedIn = false;

$colors = ["red", "green", "blue"];

$user = [
    "name" => "Alice",
    "age" => 30
];

$matrix = [
    [1, 2],
    [3, 4]
];

class Car {
    public $brand = "Toyota";
}

$car = new Car();
$middleName = null;
$file = fopen("example.txt", "r");
```

### string operations

- get string length: `echo strlen("hello"); // 5`

concatenate with `.`:

```php
$name = "John";
$greeting = "Hello, " . $name; // "Hello, John"
```

string interpolation (requires double quotes):

```php
$name = "Alice";
echo "Hi, $name!";  // Outputs: Hi, Alice!
```

get substring:

```php
$str = "abcdef";
echo substr($str, 1, 3); // "bcd"
```

### type checking

- `is_int()`
- `is_string()`
- `is_array()`
- `is_object()`
- `is_null()`
- `is_bool()`

```php
$val = 42;
if (is_int($val)) {
    echo "It's an integer";
}
```

### type casting

```php
$val = "10";
$intVal = (int)$val;      // 10
$floatVal = (float)$val;  // 10.0
$boolVal = (bool)$val;    // true
$x = (string)100;         // "100"
```

## if else

> [!WARNING]
> avoid `and` and `or` in expressions
> because both have lower precedence than `=`, which can lead to unexpected behaviour.
> use `&&` and `||` instead
> use `and` and `or` in control flow (`if`, `while`, etc)

```php
<?php
$grade = 85;

if ($grade >= 90) {
    echo "Excellent!";
} elseif ($grade >= 60) {
    echo "You passed.";
} else {
    echo "You failed.";
}
?>
```

## loops

```php
for ($i = 0; $i < 5; $i++) {
    echo $i;
}

$colors = ["red", "green", "blue"];
foreach ($colors as $color) {
    echo $color;
}
```

## functions

```php
function greet($name) {
    return "Hello, $name!";
}

echo greet("Alice");
```

- `range(3, 5);`: returns array `[3, 4, 5]`
- `var_dump($var1)`: display info about 1 or more variables
- `sqrt($num)`
- `pow($num)`
- `round($num)`
- `floor($num)`
- `ceil($num)`

## classes and objects

use `->` to access properties and methods of an object

```php
class Person {
    public $name = "Alice";

    public function sayHello() {
        return "Hello, I'm $this->name";
    }
}

$p = new Person();
echo $p->name;           // Outputs: Alice
echo $p->sayHello();     // Outputs: Hello, I'm Alice
```

## superglobals

> special variables available anywhere

- `$_GET`, `$_POST`: form data
- `$_SESSION`, `$_COOKIE`: user/session data
- `$_SERVER`: server environment
- `$_GLOBALS`: superglobal associative array that stores all global variables
    - associative array: an array where keys are strings instead of numeric indexes
    - `echo $GLOBALS['x'];  // Access global variable $x`
- `$_ENV`: superglobal associative array in PHP that contains environment variables passed to the current script
    - `echo $_ENV['APP_ENV'];  // Output: production`

## connecting with a database

### php + mysql

```php
$host = "localhost";
$user = "root";
$password = "";
$dbname = "my_database";

$conn = new mysqli($host, $user, $password, $dbname);

// Check connection
if ($conn->connect_error) {
    die("Connection failed: " . $conn->connect_error);
}
```

run sql query:

```php
$sql = "SELECT id, name FROM users";
$result = $conn->query($sql);

if ($result->num_rows > 0) {
    // Loop through rows
    while ($row = $result->fetch_assoc()) {
        echo "ID: " . $row["id"] . " - Name: " . $row["name"] . "<br>";
    }
} else {
    echo "No results";
}
```

insert data:

```php
$name = "Alice";
$sql = "INSERT INTO users (name) VALUES ('$name')";

if ($conn->query($sql) === TRUE) {
    echo "New record created";
} else {
    echo "Error: " . $conn->error;
}
```

use `prepare` statements to avoid sql injection:

```php
$stmt = $conn->prepare("INSERT INTO users (name) VALUES (?)");
$stmt->bind_param("s", $name);
$stmt->execute();
```

- closing the connecting: `$conn->close();`

