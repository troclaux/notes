
Stay in search mode in Vim after searching for a pattern, instead of pressing enter:
- ctrl-g
- ctrl-t

indent and unindent line:
- ctrl-t
- ctrl-d


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


## substitute command :%s/

After you run the command :%s/oldword/newword/c in Vim, the prompt "replace with nova palavra (y/n/a/q/l/^E/^Y)?" appears, which provides you with several options to respond to each replacement:
- y (yes): replace the current instance of the old word with the new word
- n (no): do not replace the current instance of the old word, and move on to the next instance
- a (all): replace all remaining instances of the old word with the new word, without prompting again
- q (quit): stop the substitution process entirely
- l (last): replace the current instance of the old word with the new word, and stop prompting for replacements
- ^E (scroll down): scroll the screen down one line
- ^Y (scroll up): scroll the screen up one line

# plugins

## vim fugitive

- commit amend:
  - :ca
- git stash
  - :czz
- git unstash
  - :czA
