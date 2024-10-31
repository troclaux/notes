
# CLI Tools

> [!IMPORTANT]
> path behaviour changes depending on whether there is a trailing slash
> with `/`: refers to the contents of the directory
> without `/`: refers to the directory itself
> behaviour present in `cp`, `mv`, `scp`, `rsync`, `ln`, `tar`, `zip`, `find`, `curl`, `du`

## CLI tools with its own markdown notes

- [basic shell commands](./shell.md)
- [git](./git.md)
- [docker](./docker.md)
- [psql](./postgresql.md)

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

## jq

## curl

> used to transfer data using various protocols

- downloads files
- interacts with APIs
- tests network performance
- can use many protocols:
  - HTTP
  - HTTPS
  - FTP
  - SCP
  - SMTP
  - POP3
  - etc

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
- `-x`: extracts the audio from the video file.
- `--audio-format mp3`: converts the extracted audio to mp3 format.

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
