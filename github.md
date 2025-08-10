
# github

> web-based platform that hosts git repositories online

[git](./git.md)

## basics

- repository (repo): storage location for a project containing all its files, history, and configuration
  - can contain code, documentation, images, and other project-related files
  - tracks all changes made to files in your project
- branch: allow you to work on different versions of your project at the same time without mixing up new features or bug fixes until they’re ready
- fork: personal copy of someone else's project
- pull request: submit changes from one branch to another
  - typically used to propose changes to a project
  - allows for discussion and review before merging changes
- issues: github feature that tracks tasks, enhancements, or bugs
  - can be used to:
    - track bugs and feature requests
    - assign work to team members
    - organize project tasks
    - collect user feedback

- upstream vs downstream: describe the direction of data flow between repos (normally between forks of a repo)
  - upstream repository: original source of code, where a copy (fork) came from
  - downstream repository: code that is receiving updates (fork)

## github account authentication with SSH keys

```bash
rm ~/.ssh/id_ed25519 && rm ~/.ssh/id_ed25519.pub
```

1. generate SSH keys

```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
```

2. you will be prompted to choose file location to save the key, just press enter to accept default location

3. you will be prompted to choose password, insert your password

4. add SSH key to SSH agent

```bash
eval "$(ssh-agent -s)"
```

5. add SSH private key to SSH agent

reasons to add to SSH agent:
- you won't need to enter the password every time you use the key
- it's safer, because the agent will save the SSH keys encrypted in memory

```bash
ssh-add ~/.ssh/id_ed25519
```

6. add SSH key to github account

copy ssh public key

```bash
cat ~/.ssh/id_ed25519.pub | xclip -selection clipboard
```

navigate to Github's "SSH and GPG keys" section

select "new SSH key" button

choose title of key

for key type, choose "Authentication Key"

paste public SSH key

7. test SSH connection

```bash
ssh -T git@github.com
```

you should see a message like this:

```
Hi username! You've successfully authenticated, but GitHub does not provide shell access.
```

8. update dotfiles repo

```bash
rm ~/dotfiles/.ssh/id_ed25519 && rm ~/dotfiles/.ssh/id_ed25519.pub
```

```bash
cp ~/.ssh/id_ed25519 ~/.ssh/id_ed25519.pub ~/dotfiles/.ssh
```

> [!IMPORTANT]
> don't forget to encrypt keys before adding to github repo

```bash
ansible-vault encrypt ~/dotfiles/.ssh/id_ed25519
```

```bash
ansible-vault encrypt ~/dotfiles/.ssh/id_ed25519.pub
```

you can decrypt the SSH keys later if necessary:

```bash
ansible-vault decrypt ~/dotfiles/.ssh/id_ed25519
```

```bash
ansible-vault decrypt ~/dotfiles/.ssh/id_ed25519.pub
```
