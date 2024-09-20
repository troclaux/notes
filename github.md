
# github

[git](./git.md)

## use SSH keys to associate with your github account

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
