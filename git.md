[github](./github.md)

# git

> version control system (VCS) designed to track changes in files

- allows multiple developers to work on the same project without interfering in each others changes

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

## basic commands

- SHA-1 hash can be used to see the contents of a commit: `git cat-file -p <hash>`
- search code code/content changes of each commit: `git log -S <pattern>`
- log the commits of remote repo: `git log remote/branch`, `git log origin/primeagen`
- set the default upstream branch for the master branch: `git branch --set-upstream-to=origin/master master`
- show history of commits in a graph: `git log --oneline --graph`

- search commit messages: `git log --grep="fix container image push"`
- view git diff using commit hash: `git show abc123def456`

### git switch

- switch to a branch: `git switch <branch>`
- create a new branch and switch to it: `git switch -c <branch>`
- create a new branch from a specific commit: `git switch -c <branch> <commit_hash>`

### git remote

- show the remote repositories: `git remote -v`
- add a remote repository: `git remote add <name> <url>`
- the remote can be local or on the internet: `git remote add origin ~/Documents/remote-git`
- undo the last commit while keeping the changes: `git reset --soft HEAD~1`

### git stash

- Show the list of stashes: `git stash list`
- create a stash with a message: `git stash -m 'message'`
- apply the last stash: `git stash pop --index <stash index>`
- apply a commit to the current branch: `git cherry-pick <commit>`
- show the history of the HEAD: `git reflog`
- create an alias for a command: `git config --global alias.<alias> <command>`
- show the git configuration: `git config --list`
- show the commits that match the pattern: `git log --grep <pattern> -p`
  - `-p` shows the diff

### git bisect

- start the bisect process: `git bisect start`
- mark the current commit as bad: `git bisect bad`
- mark a commit as good: `git bisect good <commit>`
- stop the bisect process: `git bisect reset`
- Automatically find the commit that introduced a bug: `git bisect run <command> --run`
  - `--run` is the command to run
- create a new worktree: `git worktree add ../<new-dir/branch`
- Rebase from branch1 onto branch2 (run this command on branch1): `git rebase <branch2>`

> [!WARNING]
> You should never rebase a public branch (like main) onto anything else. Other developers have it checked out, and if you change its history, you'll cause a lot of problems for them.
> However, with your own branch, you can rebase onto other branches (including a public branch like main) as much as you want.

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

## questions

- merging branch with upstream problems
  - i dont understand why it is necessary to create 2 merge commits to fix the different past of the branch when rebasing
- should i rebase private branch into public branch before pushing?
