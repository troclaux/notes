# Back up with rsync

### To back up

```shell
rsync -av --progress --delete ~/source/directory ~/backups
```

### Backing up home directory from another storage device

```shell
rsync -av --progress --delete /mnt/backups/directory ~/backups
```

### To back up regularly:

1. Edit 'crontab' file:
  ```shell
  crontab -e
  ```
2. Add a Cron job to back up every Sunday at 2 AM:
  ```shell
  0 15 * * 0 rsync -av --delete ~/.config/nvim ~/backups
  ```

### To restore

```shell
rsync -av --delete ~/backups/nvim ~/.config
```

```shell
rsync -av --delete ~/backups/troclaux /home
```

### Restore with a Dry Run

Allows you to see what changes will be made without actually applying them
```shell
sudo rsync -av --delete --dry-run ~/backups/nvim ~/.config
```

### Verifying the Cron Job

After saving, verify that the cron job has been added correctly by listing the current cron jobs:
```shell
crontab -l
```
