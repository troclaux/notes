# back up with rsync

to back up:

```bash
rsync -av --progress --delete ~/source/directory ~/backups
```

backing up home directory from another storage device:

```bash
rsync -av --progress --delete /mnt/backups/directory ~/backups
```

to back up regularly on a specific schedule:

1. edit 'crontab' file:
  ```bash
  crontab -e
  ```
2. add a cron job to back up every sunday at 2 am:
  ```bash
  0 15 * * 0 rsync -av --delete ~/.config/nvim ~/backups
  ```

to restore:

```bash
rsync -av --delete ~/backups/nvim ~/.config
```

```bash
rsync -av --delete ~/backups/troclaux /home
```

to restore with a dry run:

allows you to see what changes will be made without actually applying them
```bash
sudo rsync -av --delete --dry-run ~/backups/nvim ~/.config
```

## verifying the cron job

after saving, verify that the cron job has been added correctly by listing the current cron jobs:
```bash
crontab -l
```

example of cron job command:

```bash
0 15 * * 2 rsync -av --delete ~/Music/chill_game_ost ~/backups >> ~/backups/logs/chill_game_ost.log 2>&1
```
