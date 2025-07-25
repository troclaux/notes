# Neovim

- stay in search mode in Vim after searching for a pattern, instead of pressing enter:
  - ctrl-g
  - ctrl-t

- indent and unindent line:
  - ctrl-t
  - ctrl-d

- use `:Ctrl+f` to navigate Vim's command history with familiar motions, such as `j` and `k`, to quickly recall and reuse previous commands

- while in insert mode, use `Alt` to execute normal mode commands
- repeat the last recorded register count times: `Q`
- insert the output of a terminal command: `:.!<terminal command>`

## registers

- view register contents: `:reg`
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

## macro

- replay last macro: `Q` or `@@`
- replay macro in register `a` 10 times: `10@a`
- use different **lowercase letters** (e.g. `qa`, `qb`) for separate macros

> [!TIP]
> Uppercase registers append in macros.
> When recording macros in Neovim: `qw` records to register `w` (overwrites it), but `qW` **appends** to register `w`.
> So `qW` doesn't create a new macro, it adds to the existing one in `w`.

## range

range can be specified before most commands

- line
  - `:21s/old/new/g`
- last line
  - `:$s/old/new/g`
- multiple lines
  - `:4, 6!wc -w`
    - count all words in the 3 lines specified and replaces current line with the result
- current line
  - `:.s/old/new/g`
  - `:.!wc -w`
- current line to end
  - `:.,$s/old/new/g`
- all lines
  - `:%s/old/new/g`

## substitute command

after running the command `:%s/old/new/c`, the prompt `replace with new (y/n/a/q/l/^E/^Y)?` appears, which provides several options:

- `y (yes)`: replace the current instance of the old word with the new word
- `n (no)`: do not replace the current instance of the old word, and move on to the next instance
- `a (all)`: replace all remaining instances of the old word with the new word, without prompting again
- `q (quit)`: stop the substitution process entirely
- `l (last)`: replace the current instance of the old word with the new word, and stop prompting for replacements
- `^E (scroll down)`: scroll the screen down one line
- `^Y (scroll up)`: scroll the screen up one line

- you can use other delimiters for the substitute command:
  - `s#old#new#` (using # instead of /)
  - `s|old|new|` (using | instead of /)
  - `s@old@new@` (using @ instead of /)
  - `s^old^new^` (using ^ instead of /)

### format

- format: `:range s/old/new/cgil`
  - `c`: confirm each substitution
  - `g`: replace all occurrences in the line
  - `i`: ignore case for pattern
  - `l`: don't ignore case for pattern

### grouping and backreferences

- `\0`: whole matched pattern
- `\1`: matched pattern in the first pair of `\(\)`
- `\2`: matched pattern in the second pair of `\(\)`

# plugins

## list of enabled lazyvim extras

- ai.copilot    blink-cmp-copilot  copilot-cmp  copilot.lua  blink.cmp  lualine.nvim  nvim-cmp
- ai.copilot-chat  CopilotChat.nvim  edgy.nvim
- coding.blink  blink.cmp  friendly-snippets  blink.compat  catppuccin
- editor.telescope  dressing.nvim  nvim-lspconfig  telescope-fzf-native.nvim  telescope.nvim
- util.dot    mason.nvim  nvim-lspconfig  nvim-treesitter

## vim fugitive

- commit amend: `ca`
- git stash: `czz`
- git unstash: `czA`

