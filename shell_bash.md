
- [CLI tools](./cli_tools.md)

# shell

> program that interacts with the kernel to access and manage system resources

## basic commands

- `!!`: run last command
- `$(...)`: command substitution
  - `nvim $(find . | fzf)`: open a file from a fuzzy find in neovim
  - quoting the substitution (e.g. `"$()"`) prevents word splitting and pathname expansion, which is usually desirable
    - `name="Arthur" && echo 'Hello, '"$name"'! How are you today?'  # Output: Hello, Arthur! How are you today?`

- `ls`: list items in current directory
- `pwd`: print present working directory
- `touch`: create new file
- `mkdir`: create new directory
- `cd`: change to directory
- `mv`: rename or move file
- `rm`: remove file or directory
  - `rm -r`: remove directory and its contents
  - `rm -rdf`: remove directory and its contents without confirmation
- `cp`: copy file or directory
- `chmod`: change file permissions
  - e.g. `chmod +x script.sh`
- `chown`: change ownership of file or directory
  - e.g. `sudo chown -R $USER ~/.config/nvim/`

> [!IMPORTANT]
> path behaviour changes depending on whether there is a trailing slash
> with `/`: refers to the contents of the directory
> without `/`: refers to the directory itself
> behaviour present in `cp`, `mv`, `scp`, `rsync`, `ln`, `tar`, `zip`, `find`, `curl`, `du`

## wildcard matching

- `*`: matches zero or more characters except slashes `/`
- `**`: matches zero or more directories
- `?`: match any single occurrence of any character
- `[abcd]`: matches any character inside the square brackets
- ranges:
  - `[A-Z]`: matches strings that go from uppercase A to uppercase Z
  - `[A-Za-z0-9]`: matches any single alphanumerical string
- `[!abcd]` or `[^abcd]`: matches any character that is NOT inside the square brackets
- `{}` (brace expansion): generates arbitrary strings (not a pattern match)
  - `{3..9}`: defines range inside curly braces

example:

```bash
mkdir dir{1,2,3}
mkdir dir{1..3}
```

is equivalent to:

```bash
mkdir dir1 dir2 dir3
```

more examples:

```bash
ls *.pdf *.docx
ls *[0-9]?.{jpg,png}
```

## command separators

- `>`: replace contents of file after the `>` with the output of command before the `>`
  - `echo "delete everything and write this sentence" > ~/Documents/new_file.txt`
- `>>`: Append output of previous command to the file after the separator
  - `echo "append this at the end of file" >> ~/Documents/file.txt`
- `\`: split shell command into multiple lines
- `;`: run commands sequentially regardless of whether previous command was successful or not
  - default behaviour between commands in a shell script when there isn't separators between
  - `command1; command2; command3`
- `|` (pipe): forward/pipe the output of a cli program as the input for the next program
  - `cat file.txt | wc -l`
- `|&`: pipe both stdout and stderr to the next command
  - shorthand for `2>&1 |`
  - bash-specific (not POSIX)
- `||` (logical OR): run next command only if the previous command fails (exit status != 0)
  - `ls this_file_does_not_exist.txt || echo 'file not found'`
- `&&` (logical AND): runs the next command only if the previous command was successful
  - `command_1 && command_2 && command_3`
- `2>`: redirect stderr
  - `command 2> error.log`

> [!TIP]
> join multiple outputs and pipe everything into a single `fzf` by surrounding everything with parenthesis and separating each command with `;`

```bash
(ls /usr/share/applications/ ; flatpak list | awk '{print $2 ".desktop"}') | fzf
```

- `#`: comment
- `#!/bin/sh`: shebang
  - used to specify the interpreter that will run the script
- `chmod +x script.sh`: make script executable
  - [chmod](/cli_tools.md#chmod)
- `./script.sh`: run script
- `echo $?`: print exit status of last command
- `exit 0`: exit script with status 0
  - learn about status codes [here](#exit-status)

## variables

```bash
greeting=Hello
name=Tux
echo "$greeting $name"
```

- you can use the result of commands for another command

```bash
cd "$(find "$HOME" -type d | fzf)"
```

### special variables

- `$#`: the number of arguments passed to the script
- `$?`: holds the exit status (return code) of the last command that was run
- `$@`: a list of all the arguments, each one preserved as a separate word (i.e., quoted separately)
- `$*`: all the arguments passed, but treated as a single string (unless quoted properly)
- `$0`: the name of the script file being executed
- `$1`: the first argument passed to the script
- `$2`: the second argument, and so on...

### strings

- single quotes
  - literal interpretation
    - single-quoted strings are taken literally
    - all characters within the single quotes are preserved exactly as they are
  - no expansion
    - variables, command substitutions, and special characters are not expanded or interpreted
- double quotes
  - expansion allowed
    - double-quoted strings allow variable expansion, command substitution, and certain escape sequences
  - partial interpretation
    - while most special characters are preserved, $, `, and \ (when followed by certain characters) are interpreted
- read user input: `read -p "Enter your age" variable_name`

## if else

The script belows does the following:

1. checks if directory `~/.config/nvim` exists
2. recursively change ownership of directory contents

> [!IMPORTANT]
> single brackets `[` and double `[[` brackets behave differently
> `[` is in POSIX, `[[` isn't in POSIX
> [explanation](https://unix.stackexchange.com/questions/49007/when-should-i-use-vs-in-bash-single-vs-double-brackets)

```bash
#!/bin/bash

if [ condition ]
then
  statement
elif [ condition ]; then
  statement
else
  do this by default
fi
```

- `num1 -eq num2`: num1 is equal to num2
- `num1 -ge num2`: num1 is greater than or equal to num2
- `num1 -gt num2`: num1 is greater than num2
- `num1 -le num2`: num1 is less than or equal to num2
- `num1 -lt num2`: num1 is less than num2
- `num1 -ne num2`: num1 is not equal to num2

practical example:

```bash
#!/bin/bash

if [[ $(id -u) -eq 0 ]]; then
  echo "You are running as root!"
else
  echo "You are not running as root."
fi
```

- check if a file or directory exists: `if [ -e "path/to/your/file" ]; then`
- check if a file exists: `if [ -f "path/to/your/file" ]; then`
- check if a directory exists: `if [ -d "path/to/your/directory" ]; then`
- check if a file exists and is not empty: `if [ -s "path/to/your/file" ]; then`

```bash
if [ -d "$HOME/.config/nvim" ]; then
  sudo chown -R $USER ~/.config/nvim/
  echo "Successfully changed ownership of ~/.config/nvim/ directory"
else
  echo "Failed to change ownership of ~/.config/nvim. Directory does not exist"
fi
```

### loops

[loop through files](https://www.digitalocean.com/community/tutorials/workflow-loop-through-files-in-a-directory)

```bash
#!/bin/bash

directory="$HOME/Documents"

for file in "$directory"/*; do
  if [ -f "$file" ]; then
    echo "Found file: $file"
  fi
done
```

```bash
#!/bin/sh

INPUT_STRING=hello
while [ "$INPUT_STRING" != "bye" ]
do
  echo "Please type something in (bye to quit)"
  read INPUT_STRING
  echo "You typed: $INPUT_STRING"
done
```

```bash
#!/bin/bash

for i in {1..5}
do
  echo $i
done
```

```bash
for FILE in *; do cp "$FILE" "$FILE.bak"; done
```

reminder: ranges are simpler than loops:

```bash
touch eixo_5/cnu_eixo_5_mq_aula_0{3..9}.md
```
- you can also use loops simplify operations on multiple files with similar names:
  - `for i in {5..7}; do touch cnu_eixo_5_aula_0$i.md; done`

### renaming files

- changing filenames in current directory:
  - `for file in ./cnu_*; do mv "$file" "${file#./cnu_}"; done`
- changing filenames in another directory:
  - `for file in pdfs/cnu_*; do mv "$file" "pdfs/${file#pdfs/cnu_}"; done`

## exit status

> number that a command or script returns when it finishes running

> this number tells whether the operation was successful or not

- exit status is also called return code
- every command in a shell script returns an exit status

- 0 = success
- 1-255 = error: different programs use different numbers to represent different types of errors

- `exit 0`: exit shellscript with `0`

## standard data streams

in unix-like systems, every process has three standard data streams:

| Stream   | File Descriptor | Purpose                           |
| -------- | --------------- | --------------------------------- |
| `stdin`  | `0`             | Standard Input (input to program) |
| `stdout` | `1`             | Standard Output (normal output)   |
| `stderr` | `2`             | Standard Error (error messages)   |

- `stdin`: input stream (default: keyboard)
- `stdout`: output stream for program output (default: terminal screen)
- `stderr`: output stream for error messages (default: also the terminal)

useful redirection examples:

```bash
command > out.txt           # send stdout to file
command 2> err.txt          # send stderr to file
command > out.txt 2>&1      # send both stdout and stderr to file
command &> combined.txt     # same as above (Bash shortcut)
command < input.txt         # read stdin from file
```

## POSIX

> Portable Operating Systm Interface

> standards define by IEEE to maintain compatibility between unix-like operating systems

- defines core APIs, command-line utilities and shell behaviour that compliant systems and scripts should support
- examples:
  - shell syntax: `test`, `if`, `while`, `for`, etc
  - commands like `ls`, `cp`, `mv`, `find`, `grep`, `awk`, `sed`
  - system calls like `fork()`, `exec()`, `open()`, `read()`, `write()`

## zsh

> [!TIP]
> add private credentials like API keys or personal info in `~/.zshenv` as environment variables
> `export MY_CREDENTIAL="my_email@mail.com"`

- list all keybindings: `bindkey`
- creating zsh keybinds:
  - `bindkey -s '^h' "nvim . -c 'Telescope find_files'\n"`
  - `bindkey -s '^b' "!!\n\n"`
  - `bindkey '^Y' autosuggest-accept`

## oh-my-zsh

- `omz update`: update oh-my-zsh
- `zsh_stats`: list most used commands
- `-`: go to previous directory
- `d`: list last visited directories
