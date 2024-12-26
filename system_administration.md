
# system administration

> 

## user management

- `useradd`: add a new user
- `userdel`: delete a user
- `usermod`: modify a user
- `passwd`: change a user's password

## group management

- `groupadd`: add a new group
- `groupdel`: delete a group
- `groupmod`: modify a group

## sudo configuration

- `sudo`: run a command as another user
- `visudo`: edit the sudoers file


## scheduling tasks with cron

[cron](/cli_tools.md#rsync)

- `crontab`: schedule tasks to run at specific times

### mounting and unmounting filesystems

> process of attaching a filesystem accessible to a specific point in your system's directory tree

```bash
# Basic mount command syntax
mount [device] [mount point]

# Example: mounting a USB drive
mount /dev/sdb1 /mnt/usb

# View mounted filesystems
mount
df -h  # Shows mounted filesystems with human-readable sizes

# Read-only mount
mount -o ro /dev/sdb1 /mnt/usb

# Read-write mount
mount -o rw /dev/sdb1 /mnt/usb

# Unmounting
umount /mnt/usb
```

- mount point: empty directories where filesystems are attached
- common locations:
  - `/mnt`: temporary mounts
  - `/media`: removable media
  - `/`: root filesystem

