# Git jargon
- HEAD
  - The current branch
- SHA-1 hash
  - A unique identifier for a commit
- Upstream branch
  - The branch that the current branch is tracking
- Downstream branch
  - The branch that is tracking the current branch
- Index
  - The staging area
- Work tree/working tree
  - Set of files that represent your project
  - Is set up by `git init` or `git clone`

# What I learned
- Upstream vs downstream branches
- How to use SHA-1 hash
  - `git cat-file -p <hash>`
    - Can be used to see the contents of a commit
- `git log -S`
- `git branch --set-upstream-to=origin/master master`
  - Set the default upstream branch for the master branch
- `git log --oneline --graph`
- `git remote -v`
  - Show the remote repositories
- `git remote add <name> <url>`
  - Add a remote repository
  - The remote can be local or on the internet
    - eg: `git remote add origin ~/Documents/remote-git`
- `git reset --soft HEAD~1`
  - Undo the last commit while keeping the changes
- `git stash list`
  - Show the list of stashes
- `git stash -m 'message'`
  - Create a stash with a message
- `git stash pop --index <stash index>`
  - Apply the last stash
- `git cherry-pick <commit>`
  - Apply a commit to the current branch
- `git reflog`
  - Show the history of the HEAD
- `git config --global alias.<alias> <command>`
  - Create an alias for a command
- `git config --list`
  - Show the git configuration
- 

# Questions
- merging branch with upstream problems
  - i dont understand why it is necessary to create 2 merge commits to fix the different past of the branch
- should i rebase private branch into public branch before pushing?


