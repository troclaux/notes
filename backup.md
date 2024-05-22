# Back up with rsync

### Command to back up:

```shell
rsync -av --progress --delete ~/source/directory ~/backups
```

### Backing up home directory from another storage device

```shell
rsync -av --progress --delete ~/backups /mnt/backups/troclaux-backups
```

### To back up regularly:

1. Edit 'crontab' file:
  ```shell
  crontab -e
  ```
2. Add a Cron job to back up every Sunday at 2 AM:
  ```shell
  0 15 * * 0 rsync -av --delete ~/.config/nvim ~/backups/nvim
  ```

### Restore home directory

```shell
rsync -av --delete ~/backups/nvim ~/.config/nvim/
```

### Restore with a Dry Run

Allows you to see what changes will be made without actually applying them
```shell
sudo rsync -av --delete --dry-run /mnt/backups/directory_backup ~/backups/directory_backup
```

### Verifying the Cron Job

After saving the crontab file, you can verify that the cron job has been added correctly by listing the current cron jobs:
```shell
crontab -l
```
