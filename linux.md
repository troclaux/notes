
# linux

> open-source operating system

- highly flexible: users can customize everything
- secure: less prone to malware

- distributions: different packaged versions of linux
  - 3 main distributions:
    - debian
      - ubuntu
    - fedora
    - arch
  - each distribution includes various software components:
    - system libraries
    - applications
    - package manager (e.g. `apt`, `dnf`, `pacman`)

- kernel
  - the core of the OS
  - one of the first programs loaded during startup
  - acts as a bridge between hardware and user-level applications, controlling the allocation of system resources
  - responsible for:
    - hardware abstraction
    - process scheduling and management
    - memory allocation and management
    - file system operations (read/write)
    - device I/O operations
    - system calls handling
    - interrupt handling
    - inter-process communication (IPC)
- GNU system utilities
  - cp, mv, rm, mkdir, grep, sed, awk, sort, uniq, chmod, ssh, etc
- GNU libraries
  - glibc (GNU Clibrary)
    - string manipulation
    - memory management (`malloc`, `free`)
    - math operations
    - file I/O
    - networking
  - libstdc++ (GNU standard C++ library)
  - libm (GNU math library)
  - libutil (GNU utility library)
  - librt (GNU real-time library)
  - etc

## linux directory structure

- `/bin`: binary or executable programs
- `/boot`: boot loader files, Linux kernel
- `/dev`: device files
- `/etc`: system configuration files
- `/home`: user home directories
- `/lib`: system libraries
- `/media`: mount point for removable media
- `/mnt`: mount point for temporary filesystems
- `/opt`: optional application software packages
- `/proc`: kernel and process information
- `/root`: root user home directory
- `/sbin`: system administration binaries
- `/srv`: service data
- `/tmp`: temporary files
- `/usr`: user programs and data
- `/var`: variable files (logs, mail, etc.)
