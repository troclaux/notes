
- [basic shell commands](./bash_shell.md)
- [git](./git.md)
- [docker](./docker.md)
- [kubernetes](./kubernetes.md)
- [psql](./postgresql.md)

# CLI tools

> Command Line Interface for typing commands

- CLI (Command Line Interface) vs terminal vs shell
  - CLI: interface for typing commands
    - e.g. bash, powershell, CMD
  - terminal: app that provides access to the shell via text input/output (frontend)
    - e.g. gnome terminall, kitty, alacritty
  - program that interprets commands and executes them
    - e.g. bash, zsh, fish

> [!IMPORTANT]
> path behaviour changes depending on whether there is a trailing slash
> with `/`: refers to the contents of the directory
> without `/`: refers to the directory itself
> behaviour present in `cp`, `mv`, `scp`, `rsync`, `ln`, `tar`, `zip`, `find`, `curl`, `du`

## echo

> print strings to output

- use " instead of ' for variable expansion or command substitution
- use -n to suppress the trailing newline
- use -e to enable interpretation of backslash escapes
- use -E to disable interpretation of backslash escapes

```bash

## Print a message with environment variables
echo "My path is $PATH"

## Print a message without the trailing newline
echo -n "Hello World"

```
## cat

> concatenate files and print on the standard output

- `cat file.txt` print the contents of a file
- `cat file1.txt file2.txt` concatenate multiple files
- `cat file1.txt file2.txt > file3.txt` concatenate multiple files and save to a new file
- `cat file1.txt file2.txt >> file3.txt` concatenate multiple files and append to an existing file
- `cat file1.txt file2.txt | less` concatenate multiple files and pipe to less
- `cat file1.txt file2.txt | grep "pattern"` concatenate multiple files and pipe to grep

## find

> print files, directories and subdirectories recursively in the directory tree

- `find ~/ ~/Documents ~/Downloads ~/learning -name "*.md"` search multiple directories
- `find . -name "*.txt"` find all files with .txt extension
- `find . -type f` find all files
- `find . -type d` find all directories
- `find . -empty` find all empty files
- `find . -size +10M` find all files larger than 10MB
- `find . -size -10M` find all files smaller than 10MB
- `find . -maxdepth 1` search only in the current directory
- `find . -mindepth 1` search all directories except the current directory
- `find . -exec` execute a command on the found files

## grep

- `grep -i` case insensitive
- `grep -v` invert match
- `grep -r` recursive search
- `grep -l` list files with matches
- `grep -c` count matches
- `grep -n` show line numbers
- `grep -E` extended regex

## wc

> counts the number of lines, words, and characters in a file

- `wc -l` counts the number of lines
- `wc -w` counts the number of words
- `wc -m` counts the number of characters
- `wc -c` counts the number of bytes

```bash
wc file.txt
```

## xdg-open

## sed

## curl

> transfer data from or to a server using URLs and various protocols

[curl tutorial](https://curl.se/docs/tutorial.html)

- use cases
  - testing APIs with HTTP requests
  - automation
  - debugging
  - file download
  - get my public IPv4: `curl -4 ifconfig.me`
  - get my public IPv6: `curl ifconfig.me`
- can use many protocols: HTTP, HTTPS, FTP, FTPS, etc

example of curl command:

```bash
curl -X POST http://example.com/resource -H "Content-Type: application/json" -d '{"key1":"value1","key2":"value2"}'
```

- `-X`: specify the request method (GET, POST, PUT, DELETE)
- `-H`: specify the request headers
- `-d`: specify the request data
  - URL-encoded data: `"param1=value1&param2=value2"`
  - JSON: `'{"key1":"value1","key2":"value2"}'`

```bash
curl -X POST -H "Content-Type: application/json" -d '{"body": "this is a social network post"}' http://localhost:8080/api/post -w "%{http_code}"
```
- `-w`: print the response code
- `%{http_code}`: print the response code
  - `%`: format specifier
  - `{http_code}`: built-in variable in curl that stores the response code

- curl built-in variables:
  - `http_code`: response code
  - `http_version`: HTTP version
  - `protocol`: protocol used
  - `response_size`: size of the response

```bash
curl -X PUT http://localhost:3000/api/exercises \
  -H "Content-Type: application/json" \
  -H "Cookie: authjs.session-token=bd65694e-a1bb-4a3f-b7cd-43642b7146fa" \
  -d '{ "id": 1,
    "name": "Push-up new",
    "description": "A basic exercise for upper body strength new"
  }'
```

### use cases

1. run basic GET request

```bash
curl http://example.com
```

2. test API

```bash
curl http://localhost:5000/api/v1/tasks
```

3. download files

```bash
curl -O http://example.com/file.zip
```

4. upload files

```bash
curl -F "file=@/path/to/file.zip" http://example.com/upload
```

5. get file from github

```bash
curl -sL https://gist.githubusercontent.com/2KAbhishek/9c6d607e160b0439a186d4fbd1bd81df/raw/244284c0b3e40b2b67697665d2d61e537e0890fc/Shell_Keybindings.md
```

## jq

> process JSON data

- use cases:
  - parse json
  - manipulate json
  - filter json

```bash
echo '{"name": "John", "age": 30}' | jq '.name'
# Output: "John"

echo '[{"name": "John", "age": 30}, {"name": "Jane", "age": 25}]' | jq '.[] | select(.age > 28)'
# Output: {"name": "John", "age": 30}
```

- frequently used with curl to parse json responses from http requests
  - e.g. `curl https://api.github.com/users/1 | jq '.name'`
- to get a specific field in an array, use: `.[]`

```bash
echo '[{"name": "John", "email": "john@example.com"}, {"name": "Jane", "email": "jane@example.com"}]' | jq '.[].name, .[].email'
# output:
# "John"
# "Jane"
# "john@example.com"
# "jane@example.com"
```

another example

```json
{"type": "foo", "values": [1, 2, 3, 4, 5]}
{"type": "foo", "values": [69, 420, 42, 69420]}
{"type": "bar", "values": {"a": 42, "b": 69}}
{"type": "bar", "values": {"a": 1337, "b": 420}}
{"type": "bar", "values": {"a": 111, "b": 222}}
```

```bash
cat json | jq -c 'select(.type == "foo") | .values | add'

# output:
# 15
# 69951
```

- select: filter all json with `type` of `foo` value
- read the `values` field and add all values

## ffmpeg

converts media formats:

```bash
ffmpeg -i input.avi output.mp4
```
- `-i`: specifies the input file

## yt-dlp

downloading a video:

```bash
yt-dlp --write-subs --sub-lang en --embed-subs https://www.youtube.com/watch?v=dQw4w9WgXcQ
```

- `--write-subs`:downloads subtitles if available
- `--sub-lang en`: specifies the language of the subtitles
- `--embed-subs`: embed subtitles in the video file

convert video files to audio-only files (requires ffmpeg and ffprobe):

```bash
yt-dlp -f bestaudio -x --audio-format flac https://www.youtube.com/watch?v=dQw4w9WgXcQ
```

- `-f bestaudio`: selects best available audio-only format
- `-x`: extracts the audio from the video file
- `--audio-format flac`: converts the extracted audio to FLAC format

## rclone

start rclone setup:

```bash
rclone config
```

1. `no remotes found, make a new one?`
1. insert `n` to create a new remote
1. `enter name for new remote`
1. insert name for remote (i usually use `gdrive`)
1. `type of storage to configure, choose a number from below...`
1. insert the number for Google Drive
1. Leading `~` will be expanded in the file name as will environment variables such as `${RCLONE_CONFIG_DIR}`. Enter a value. Press Enter to leave empty.
1. press `Enter`
1. `Edit advanced config?`
1. insert `n`
1. `Use web browser to automatically authenticate rclone with remote?`
1. insert `y` and login to your google account in the browser
1. `Configure this as a Shared Drive (Team Drive)?`
1. insert `n`
1. `Keep this "gdrive" remote?`
1. insert `y`
1. rclone will list all remotes, insert `q` to quit configuration

test remote by listing the contents of google drive:

```bash
rclone lsd gdrive:
```

copy system directory to Google Drive:

```bash
rclone copy ~/Music/chill_game_ost gdrive:chill_game_ost -P
```

- `-P`: show transfer progress

sync files and folders:

```bash
rclone sync ~/Music/chill_game_ost gdrive:chill_game_ost
```

> [!TIP]
> you can use cronjobs to regularly sync your folders

```bash
0 15 * * 2 rclone sync ~/Music/chill_game_ost gdrive:chill_game_ost >> ~/backups/chill_game_ost.log 2>&1
```

## rsync

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

### verifying the cron job

after saving, verify that the cron job has been added correctly by listing the current cron jobs:

```bash
crontab -l
```

example of cron job:

```bash
0 15 * * 2 rsync -av --delete ~/Music/chill_game_ost ~/backups >> ~/backups/chill_game_ost.log 2>&1
```

## chown

> changes the owner and/or group of a file or directory

- each file/directory has:
  - owner: single user who "owns" the file
  - group: group associated with the file

`ls -l` shows ownership:

```bash
-rw-r--r-- 1 alice developers file.txt
```

- `alice`: owner
- `developers`: group

- examples:
  - change owner only: `chown alice file.txt`
  - change group only: `chown :developers file.txt`
  - change owner and group: `chown alice:developers file.txt`

## chmod

> manage file permissions

- permission: control what users can do with files and directories
- every file or directory has 3 types of users:
  - owner: user who owns the file (normally the user that created the file)
    - controls permissions of the file
    - only the owner (or root) can change the file's permissions (using `chmod`)
    - the owner can transfer the ownership to someone else
  - group: groups of users who can share access
  - others: all other users
- every file or directory has 3 types of permissions:
  - read `r`: view the file's contents
  - write `w`: modify the file or directory
  - execute `x`: run the file as a program or access a directory

- permissions are represented like this: `rwxr-xr--`
  - owner: `rwx`
  - group: `r-x`
  - others: `r--`
- syntax: `chmod [options] mode file`
  - examples:
    - `chmod u+x file.txt`
    - `chmod 754 script.sh`
      - representation of mode with numbers:
        - add each permission to obtain the permissions for a type of user
        - `r = 4`
        - `w = 2`
        - `x = 1`
      - owner: `7` == `rwx`
      - group: `5` == `r-x`
      - others: `4` == `r--`

## systemctl

> manages systemd system and service manager

- service: wrapper or unit managed by init systems
- daemon: program that runs in the background without user interaction
  - names typically ends with 'd' (e.g. `sshd`, `dockerd`, `httpd`)
- systemd: init system and service manager used by most modern linux distributions
  - initializes the system (boot sequence)
  - manages services
  - logging (via `jornald`)

- check status of system or service: `sudo systemctl status`, `sudo systemctl status <service-name>`
- start a service: `sudo systemctl start <service-name>`
- stop a service: `sudo systemctl stop <service-name>`
- restart a service: `sudo systemctl restart <service-name>`
- start on boot: `systemctl enable <service>`
- don't start on boot: `systemctl disable <service>`
- view logs for a service: `journalctl -u <service>`
- list all failed services: `systemctl list-units --failed`

## openssl

> used for secure communication and cryptographic operations

- use cases:
  - ssl/tls protocol management
  - key generation
  - encryption/decryption

```bash
openssl rand -base64 64 -A
```

- `rand`: generate random data
- `-base64`: specifies output format of the generated random data, in this case, base64-encoded string
- `64`: length of the random data to be generated, in this case, 64 bytes

## nmap

```bash
nmap -sV -sC -p- -T4 -A
```

- `-u`: scan udp ports
- 

## sshfs

> mount remove directory over SSH, makint it appear as a local directory on your system

```bash
sshfs [user@]host:[remote_dir] mountpoint
sshfs user@remote_host:/home/user/shared /mnt/remote_share

# Unmount
fusermount -u ~/remote
```

## troubleshooting commands

- process status: `ps aux`, `top`, `htop`, `systemctl status`
- logs: `journalctl`, `docker logs`, `tail -f`, `less`, `grep`
- network: `curl`, `ping`, `ss -tuln`, `netstat`, `lsof -i`
- disk/memory: `df -h`, `du -sh *`, `free -m`
- permissions: `ls -l`, `chmod`, `chown`, `whoami`
- container status: `docker ps`, `docker-compose ps`, `docker inspect`

### process status

- `ps aux`: list all running processes with details like user, pid, cpu/mem usage and command
  - great with `grep`
  - resources used by a process: `ps -C flameshot -o %cpu,%mem,cmd`
- `top`: real-time view of processes using the most system resources
- `htop`: enhanced `top` with colors and better ui
- `systemctl status`: shows status of systemd service
  - e.g. `systemctl status apache`: shows if it's active, failed, etc
- `systemd-analyze`: shows an overall boot time summary
- `systemd-analyze blame`: show which units took the most time during boot

### logs

- `journalctl`: view logs collected from `systemd`
  - logs for services, kernel, boots, etc
  - `journalctl -u nginx`: shows logs for the nginx service
- `docker logs`: view logs from docker container
- `tail -f`: reads the end of a file as it grows
- `less`
- `grep`

### network

- `curl`
- `ping`: checks if a host is reachable via ICMP (Internet Control Message Protocol)
  - use cases:
    - test network connectivity
    - measure latency
    - detect packet loss
- `ss -tuln`: displays socket statistics (modern replacement for `netstat`)
  - use case: see which ports your system is listening on
  - `-t`: TCP connections
  - `-u`: UDP connections
  - `-l`: listening ports
  - `-n`: show numeric IPs and ports (don't resolve names)
- `netstat`
- `lsof -i`: shows which processes have network connections open
  - network connections are treated as files in linux

### disk/memory

- `df -h`: show disk space usage of mounted filesystems
- `du -sh *`: show size of each file/folder in current directory
- `free -m`: displays memory usage

### permissions

- `ls -l`: list files and directory in long format
- `chmod`: change file/folder permissions
- `chown`: change file/folder ownership
- `whoami`: print current user
