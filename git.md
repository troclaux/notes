[github](./github.md)

# git

> version control system (VCS) designed to track changes in files

- allows multiple developers to work on the same project without interfering in each others changes

- branch: allow you to work on different versions of your project at the same time without mixing up new features or bug fixes until they’re ready
- commit: snapshots of your project’s history
  - each commit records changes to one or more files and includes a short message describing what changed
- push: sends your local commits to github
- pull: fetches and merges changes from github to your local machine
- fetch: retrieves changes from github but doesn’t merge them automatically, giving you a chance to review before merging

- HEAD: current branch
- SHA-1 hash: unique identifier for a commit
- upstream branch: branch that the current branch is tracking
- downstream branch: branch that is tracking the current branch
- index: staging area
- work tree/working tree: set of files that represent your project
  - set up with `git init` or `git clone`
- main working tree: when you run `git init`, git creates the main working tree

## cli commands

- `git cat-file -p <hash>`: SHA-1 hash can be used to see the contents of a commit
  - inspect the contents of a git object (like a commit, tree, blob or tag) in a human-readable format
- `git show abc123def456`: show git diff using commit hash
- `git log --grep="fix container image push"`: search commit messages
- `git log -S <pattern>`: search code code/content changes of each commit
- `git log -L :my_function:app.py`: show every commit that modified `my_function()` in `app.py`, with diffs
- `git log origin/myfeature`: log the commits of remote repo
- `git log --oneline --graph`: show history of commits in a graph
- `git blame <file>`: show commit hash, author, timestamp for each line of code
- `git branch --set-upstream-to=origin/master master`: set the default upstream branch for the master branch

### git switch

- `git switch <branch>`: switch to a branch
  - create a new branch and switch to it: `git switch -c <branch>`
  - create a new branch from a specific commit: `git switch -c <branch> <commit_hash>`

### git remote

- `git remote -v`: show the remote repositories
- `git remote add <name> <url>`: add a remote repository
- `git remote add origin ~/Documents/remote-git`: the remote can be local or on the internet
- `git reset --soft HEAD~1`: undo the last commit while keeping the changes

### git stash

- `git stash list`: show the list of stashes
- `git stash -m 'message'`: create a stash with a message
- `git stash pop --index <stash index>`: apply the last stash
- `git cherry-pick <commit>`: apply a commit to the current branch
- `git reflog`: show the history of the HEAD
- `git config --global alias.<alias> <command>`: create an alias for a command
- `git config --list`: show the git configuration
- `git log --grep <pattern> -p`: show the commits that match the pattern
  - `-p` shows the diff

### git bisect

- `git bisect start`: start the bisect process
- `git bisect bad`: mark the current commit as bad
- `git bisect good <commit>`: mark a commit as good
- `git bisect reset`: stop the bisect process
- `git bisect run <command> --run`: Automatically find the commit that introduced a bug
  - `--run` is the command to run
- `git worktree add ../<new-dir/branch`: create a new worktree
- `git rebase <branch2>`: Rebase from branch1 onto branch2 (run this command on branch1)

> [!WARNING]
> You should never rebase a public branch (like main) onto anything else. Other developers have it checked out, and if you change its history, you'll cause a lot of problems for them.
> However, with your own branch, you can rebase onto other branches (including a public branch like main) as much as you want.

## branching strategy

### git flow

- branches:
  - `main`: production-ready code
  - `develop`: integration branch for features
  - `feature/*`: new features (branched from `develop`)
  - `release/*`: prep for release (branched from `develop`)
  - `hotfix/*`: urgent production fixes (branched from `main`)

worflow:

1. create `feature/login-system` => merge into `develop`
1. create `release/v1.0` => test, bugfix => merge into `main` and `develop`
1. create `hotfix/urgent-bug` => merge into `main` and `develop`

- pros: separates environments clearly
- cons: overhead for small teams, many merges

### github flow

TODO

### gitlab flow

TODO

### trunk-based

TODO

## gitignore

> specifies which files and directories git should ignore

- files that should typically be ignored:
  - build artifacts and compiled code
  - dependencies and package directories
  - environment and config files with secrets
  - system files (.DS_Store, Thumbs.db)
  - ide files and directories
  - log files
  - temporary files

```gitignore
# Ignore single file
secret.txt

# Ignore directory
node_modules/
dist/

# Ignore files by extension
*.log
*.tmp

# Ignore files in any directory
**/*.pyc

# Negate ignore (don't ignore)
!important.log

# Ignore files in root only
/config.json

# Ignore files in specific directory
logs/*.log
```

## questions

- merging branch with upstream problems
  - i dont understand why it is necessary to create 2 merge commits to fix the different past of the branch when rebasing
- should i rebase private branch into public branch before pushing?
