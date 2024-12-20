
# git

> version control system (VCS) design to track changes in files

- allows multiple developers to work on the same project without interfering in each others changes

[github](./github.md)

## jargon

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
  - Is set up by `git init` or `git clone`
- main working tree: when you run `git init`, git creates the main working tree

## What I learned

- Upstream vs downstream branches
  - Upstream branch is the branch that the current branch is tracking
  - Downstream branch is the branch that is tracking the current branch
- How to use SHA-1 hash
  - Can be used to see the contents of a commit

```bash
git cat-file -p <hash>
```

- Search commits for a specific string


```bash
- git log -S
```
The git log command isn't only useful for your local repo. You can log the commits of a remote repo as well!

```bash
git log remote/branch
```

```bash
git log origin/primeagen
```

- Set the default upstream branch for the master branch

```bash
- git branch --set-upstream-to=origin/master master
```

- Show history of commits in a graph

```bash
- git log --oneline --graph
```

### git switch

- Switch to a branch

```bash
- git switch <branch>
```

- Create a new branch and switch to it

```bash
- git switch -c <branch>
```

- Create a new branch from a specific commit

```bash
- git switch -c <branch> <commit_hash>
```

### git remote

- Show the remote repositories

```bash
- git remote -v
```

- Add a remote repository

```bash
- git remote add <name> <url>
```

- The remote can be local or on the internet

```bash
- git remote add origin ~/Documents/remote-git
```

- Undo the last commit while keeping the changes

```bash
git reset --soft HEAD~1
```

### git stash

- Show the list of stashes

```bash
git stash list
```

- Create a stash with a message

```bash
git stash -m 'message'
```

- Apply the last stash

```bash
git stash pop --index <stash index>
```

- Apply a commit to the current branch

```bash
git cherry-pick <commit>
```

- Show the history of the HEAD

```bash
git reflog
```

- Create an alias for a command

```bash
git config --global alias.<alias> <command>
```

- Show the git configuration

```bash
git config --list
```

- Show the commits that match the pattern
  - `-p` shows the diff

```bash
git log --grep <pattern> -p
```

### git bisect

- Start the bisect process

```bash
git bisect start
```

- Mark the current commit as bad 

```bash
git bisect bad
```

- Mark a commit as good

```bash
git bisect good <commit>
```

- Stop the bisect process

```bash
git bisect reset
```

- Automatically find the commit that introduced a bug
  - `--run` is the command to run

```bash
git bisect run <command> --run
```

- Create a new worktree

```bash
git worktree add ../<new-dir/branch
```

- Rebase from branch1 onto branch2

`from branch1`
```bash
git rebase <branch2>
```
*WARNING*
- You should never rebase a public branch (like main) onto anything else. Other developers have it checked out, and if you change its history, you'll cause a lot of problems for them.
- However, with your own branch, you can rebase onto other branches (including a public branch like main) as much as you want.

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

### examples

```bash
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

## Questions

- merging branch with upstream problems
  - i dont understand why it is necessary to create 2 merge commits to fix the different past of the branch when rebasing
- should i rebase private branch into public branch before pushing?


