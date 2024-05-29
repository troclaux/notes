
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
- |: forward the output of a cli program as the input for the next program
	- `cat file.txt | wc -l`

e.g.:
```shell
mkdir dir{1,2,3}
```
is equivalent to:
```shell
mkdir dir1 dir2 dir3
```


### if else statement

The script belows does the following:
1. checks if directory ~/.config/nvim exists
2. recursively change ownership of directory contents

```shell
if [ -d "~/.config/nvim" ]; then
	sudo chown -R $USER ~/.config/nvim/
	echo "Successfully changed ownership of ~/.config/nvim/ directory"
else
	echo "Failed to change ownership of ~/.config/nvim. Directory does not exist"
fi
```
