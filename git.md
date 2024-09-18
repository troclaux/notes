
# git

## Jargon

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
- Main working tree
  - when you run `git init`, git creates the main working tree

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

## Questions

- merging branch with upstream problems
  - i dont understand why it is necessary to create 2 merge commits to fix the different past of the branch when rebasing
- should i rebase private branch into public branch before pushing?


