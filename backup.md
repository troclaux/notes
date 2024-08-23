# back up with rsync

### to back up

```shell
rsync -av --progress --delete ~/source/directory ~/backups
```

### backing up home directory from another storage device

```shell
rsync -av --progress --delete /mnt/backups/directory ~/backups
```

### to back up regularly on a specific schedule

1. edit 'crontab' file:
  ```shell
  crontab -e
  ```
2. add a cron job to back up every sunday at 2 am:
  ```shell
  0 15 * * 0 rsync -av --delete ~/.config/nvim ~/backups
  ```

### to restore

```shell
rsync -av --delete ~/backups/nvim ~/.config
```

```shell
rsync -av --delete ~/backups/troclaux /home
```

### restore with a dry run

allows you to see what changes will be made without actually applying them
```shell
sudo rsync -av --delete --dry-run ~/backups/nvim ~/.config
```

### verifying the cron job

after saving, verify that the cron job has been added correctly by listing the current cron jobs:
```shell
crontab -l
```

### example of cron job command

```shell
0 15 * * 2 rsync -av --delete ~/Music/chill_game_ost ~/backups >> ~/backups/logs/chill_game_ost.log 2>&1
```
