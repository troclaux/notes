# Neovim

- Stay in search mode in Vim after searching for a pattern, instead of pressing enter:
  - ctrl-g
  - ctrl-t

- indent and unindent line:
  - ctrl-t
  - ctrl-d

- use `:Ctrl+f` to navigate Vim's command history with familiar motions, such as `j` and `k`, to quickly recall and reuse previous commands

## registers

- black hole register
  - discards any text you yank or delete into it
  - to delete without saving: `"_d`
- read-only registers
  - current file name: `%`
  - run last command: `:@:`
  - print directory/name of file: `:echo @%`
  - print name of file ('tail'): `:echo expand('%:t')`
  - print full path: `:echo expand('%:p')`
  - print directory containing file ('head'): `:echo expand('%:p:h')`

## range

range can be specified before most commands

- line
  - :21s/old/new/g
- last line
  - :$s/old/new/g
- multiple lines
  - :4, 6!wc -w
    - count all words in the 3 lines specified and replaces current line with the result
- current line
  - :.s/old/new/g
  - :.!wc -w
- current line to end
  - :.,$s/old/new/g
- all lines
  - :%s/old/new/g


## substitute command

after running the command `:%s/old/new/c`, the prompt "replace with new (y/n/a/q/l/^E/^Y)?" appears, which provides several options:
- y (yes): replace the current instance of the old word with the new word
- n (no): do not replace the current instance of the old word, and move on to the next instance
- a (all): replace all remaining instances of the old word with the new word, without prompting again
- q (quit): stop the substitution process entirely
- l (last): replace the current instance of the old word with the new word, and stop prompting for replacements
- ^E (scroll down): scroll the screen down one line
- ^Y (scroll up): scroll the screen up one line

- you can use other delimiters for the s command:
  - `s#old#new#` (using # instead of /)
  - `s|old|new|` (using | instead of /)
  - `s@old@new@` (using @ instead of /)
  - `s^old^new^` (using ^ instead of /)

### format

- format: `:range s/old/new/cgil`
  - c: confirm each substitution
  - g: replace all occurrences in the line
  - i: ignore case for pattern
  - l: don't ignore case for pattern

### grouping and backreferences

- `\0`: whole matched pattern
- `\1`: matched pattern in the first pair of `\(\)`
- `\2`: matched pattern in the second pair of `\(\)`

# plugins

## vim fugitive

- commit amend: `ca`
- git stash: `czz`
- git unstash: `czA`
