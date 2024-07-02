
## Wildcard matching

- *: matches zero or more characters except slashes (/)
- **: matches zero or more directories
- ?: match any single ocurrence of anycharacter
- [abcd]: matches any character inside the square brackets
- ranges:
  - [A-Z]: matches strings that go from uppercase A to uppercase Z
  - [A-Za-z0-9 ]: matches any alphanumerical string
- [!abcd]: matches any character that is NOT inside the square brackets
- {}: match a group of names/Wildcard patterns

e.g.:
```shell
mkdir dir{1,2,3}
```
is equivalent to:
```shell
mkdir dir1 dir2 dir3
```

## Command separators

- (>)
  - replace contents of file after the > with the output of command before the >
  - ```echo "delete everything and write this sentence" > ~/Documents/new_file.txt```
- (>>)
  - Append output of previous command to the file after the separator
  - ```echo "append this at the end of file" >> ~/Documents/file.txt```
- ( \ )
  - split shell command into multiple lines
- (;)
  - Run multiple commands at once regardless of whether previous command was successful or not
  - **Default behaviour between commands in a shell script when there isn't separators between**
  - ```command1; command2; command3```
- (|)
  - forward the output of a cli program as the input for the next program
  - `cat file.txt | wc -l`
- (||)
  - Executes the second command only if the first command fails
  - `ls this_file_does_not_exist.txt || echo 'file not found'`
- (&&)
  - Runs the next command only if the previous command was successful
  - ```command_1 && command_2 && command_3```
- (2>)
  - redirect stderr
  - ```command 2> error.log```

## Shell scripting

### Variables

```shell
greeting=Hello
name=Tux
echo "$greeting $name"
```

- you can use the result of commands for another command

```shell
cd "$(find "$HOME" -type d | fzf)"
```

### Strings

- single quotes
  - Literal Interpretation
    - Single-quoted strings are taken literally
    - All characters within the single quotes are preserved exactly as they are
  - No Expansion
    - Variables, command substitutions, and special characters are not expanded or interpreted
- double quotes
  - Expansion Allowed
    - Double-quoted strings allow variable expansion, command substitution, and certain escape sequences
  - Partial Interpretation
    - While most special characters are preserved, $, `, and \ (when followed by certain characters) are interpreted


### Read user input

```shell
read -p "Enter your age" variable_name
```


### if else statement

The script belows does the following:
1. checks if directory ~/.config/nvim exists
2. recursively change ownership of directory contents

```shell
if [[ condition ]]
then
  statement
elif [[ condition ]]; then
  statement
else
  do this by default
fi
```

```shell
if [ -d "~/.config/nvim" ]; then
  sudo chown -R $USER ~/.config/nvim/
  echo "Successfully changed ownership of ~/.config/nvim/ directory"
else
  echo "Failed to change ownership of ~/.config/nvim. Directory does not exist"
fi
```

| OPERATION | SYNTAX | EXPLANATION |
| --- | --- | --- |
| Equality | num1 -eq num2 | is num1 equal to num2 |
| Greater than or equal to | num1 -ge num2 | is num1 greater than equal to num2 |
| Greater than | num1 -gt num2 | is num1 greater than num2 |
| Less than or equal to | num1 -le num2 | is num1 less than equal to num2 |
| Less than | num1 -lt num2 | is num1 less than num2 |
| Not equal to | num1 -ne num2 | is num1 not equal to num2 |

### while loops

```shell
#!/bin/sh

INPUT_STRING=hello
while [ "$INPUT_STRING" != "bye" ]
do
  echo "Please type something in (bye to quit)"
  read INPUT_STRING
  echo "You typed: $INPUT_STRING"
done
```

### for loops

```shell
#!/bin/bash

for i in {1..5}
do
  echo $i
done
```

you can also use loops simplify operations on multiple files with similar names

```shell
for i in {5..7}; do touch cnu_eixo_5_aula_0$i.md; done
```
