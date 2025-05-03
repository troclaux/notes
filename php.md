
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

php code is written between `<?php` and `?>` tags:

```php
<?php
echo "Hello, world!";
?>
```

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
- `boolean`
- `string`
- `array`
- `object`
- `null`

```php
$age = 30;
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

## if else

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

## superglobals

> special variables available anywhere

- `$_GET`, `$_POST`: form data
- `$_SESSION`, `$_COOKIE`: user/session data
- `$_SERVER`: server environment

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

