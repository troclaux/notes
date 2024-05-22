# Back up with rsync

### Command to back up:

```shell
rsync -av --progress --delete /source/directory/ /backup/directory/
```

### Backing up home directory

```shell
rsync -av --progress --delete /home/troclaux/ /mnt/backup/troclaux-backup/
```

### To back up regularly:

1. Edit 'crontab' file:
  ```shell
  crontab -e
  ```
2. Add a Cron job to back up every Sunday at 2 AM:
  ```shell
  0 15 * * 0 rsync -av --delete /home/arthur/ /mnt/backup/arthur-backup/
  ```


### Restore home directory

```shell
rsync -av --delete /mnt/backup/arthur-backup/ /home/arthur/
```

### Restore with a Dry Run

Allows you to see what changes will be made without actually applying them
```shell
sudo rsync -av --delete --dry-run /mnt/backup/troclaux-backup/ /home/troclaux/
```
### Verifying the Cron Job

After saving the crontab file, you can verify that the cron job has been added correctly by listing the current cron jobs:
```shell
crontab -l
```
