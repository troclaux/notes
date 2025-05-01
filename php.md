
# PHP

## properties

- primarily used for web development
- it is embedded in html

- dynamically typed
- interpreted programming language
- has garbage collection
- doesn't have hoisted declarations
- function-scoped variables
- mutable strings

## basic syntax

php code is written between `<?php` and `?>` tags:

```php
<?php
echo "Hello, world!";
?>
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
```

## superglobals

> special variables available anywhere

- `$_GET`, `$_POST`: form data
- `$_SESSION`, `$_COOKIE`: user/session data
- `$_SERVER`: server environment
