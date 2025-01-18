# Shell

> program that interacts with the kernel to access and manage system resources

## basic commands

- `!!`: run last command
- `$(...)`: command substitution
  - `nvim $(find . | fzf)`: open a file from a fuzzy find in neovim
  - only works in strings with double quotes
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
- `?`: match any single ocurrence of any character
- `[abcd]`: matches any character inside the square brackets
- ranges:
  - `[A-Z]`: matches strings that go from uppercase A to uppercase Z
  - `[A-Za-z0-9]`: matches any alphanumerical string
- `[!abcd]` or `[^abcd]`: matches any character that is NOT inside the square brackets
- `{}`: match a group of names/wildcard patterns
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
- `||` (logical OR): pipe stdout to the next command
  - `ls this_file_does_not_exist.txt || echo 'file not found'`
- `&&` (logical AND): runs the next command only if the previous command was successful
  - `command_1 && command_2 && command_3`
- `2>`: redirect stderr
  - `command 2> error.log`

> [!TIP]
> combine multiple outputs and pipe everything into a single `fzf` by surrounding everything with parenthesis and separating each command with `;`

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
  - 1: error
- `exit 1`: exit script with status 1
  - 0: success

### variables

```bash
greeting=Hello
name=Tux
echo "$greeting $name"
```

- you can use the result of commands for another command

```bash
cd "$(find "$HOME" -type d | fzf)"
```

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

1. checks if directory ~/.config/nvim exists
2. recursively change ownership of directory contents

> [!IMPORTANT]
> single brackets `[` and double `[[` brackets behave differently
> `[` is in POSIX, `[[` isn't in POSIX
> [explanation](https://unix.stackexchange.com/questions/49007/when-should-i-use-vs-in-bash-single-vs-double-brackets)

```bash
#!/bin/bash

if [[ condition ]]
then
  statement
elif [[ condition ]]; then
  statement
else
  do this by default
fi

```

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

### loops

[loop through files](https://www.digitalocean.com/community/tutorials/workflow-loop-through-files-in-a-directory)

```bash
##!/bin/sh

INPUT_STRING=hello
while [ "$INPUT_STRING" != "bye" ]
do
  echo "Please type something in (bye to quit)"
  read INPUT_STRING
  echo "You typed: $INPUT_STRING"
done
```

```bash
##!/bin/bash

for i in {1..5}
do
  echo $i
done
```

```bash
for FILE in *; do cp $FILE "$FILE.bak"; done;
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

## zsh

> [!TIP]
> add private credentials like API keys or personal info in `~/.zshenv` as environment variables
> `export MY_CREDENTIAL="my_email@mail.com`

- list all keybindings:
  - `bindkey`
- creating zsh keybinds:
  - `bindkey -s '^h' "nvim . -c 'Telescope find_files'\n"`
  - `bindkey -s '^b' "!!\n\n"`
  - `bindkey '^Y' autosuggest-accept`

## oh-my-zsh

- `omz update`: update oh-my-zsh
- `zsh_stats`: list most used commands
- `-`: go to previous directory
- `d`: list last visited directories
