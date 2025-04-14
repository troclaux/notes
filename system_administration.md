
# system administration

> process of managing and maintaining computer systems, servers and networks to ensure they run reliably, securely and efficiently

- responsibilities:
  - installing and configuring operating systems and software
  - managing user accounts and permissions
  - monitoring performance and availability
  - performing backups and recovery
  - applying security patches and updates
  - troubleshooting system issues
  - automating tasks with scripts (like bash)

## environments

> separate setup where an application runs, used for different stages of development

- development: where features are developed and tested locally on dev servers
  - used by: developers
  - typical setup: local machines or remote dev servers
- testing/QA (Quality Assurance): run manual or automated tests on new code before it's promoted further
  - used by: QA testes, developers
  - typical setup: closely mirrors staging, often auto-deployed by CI/CD after commits
- staging : production-like environment for final testing before going live
  - used by: QA testes, developers, product owners
- homologation/UAT (User Acceptance Testing): clients/stakeholders test and approve features before production
  - used by: clients, stakeholders, QA
  - very close to production
  - goal: validate that the app behaves correctly in real-world scenarios
- production: live system accessible to customers
  - used by: end users
  - rules: strict deployment processes, rollback plans, alerting

- cloning environment: cloning an environment means replicating the full stack — application code, docker setup, environment variables, and database — into a new isolated space, usually for debugging, migration testing or demos
  - normally done with a shell script that copies the project directory, adjusts ports in .env and docker-compose.yml and spins up containers with docker compose up
  - for the database, you can snapshot or dump the data and restore it into a new db container or rds instance, so the clone doesn’t affect the original

## deployment strategies

> techniques to safely release new versions of an application, minimizing downtime and risk

- recreate: stop old version, start the new one
  - downtime during switch
  - simple, but not ideal for real time systems
- rolling update: replace instances gradually, one by one
  - no downtime
  - requires kubernetes or docker swarms or similar
- blue-green deployment: run two environments, switch traffic when ready
  - blue: current version
  - green: new version
- canary deployment: release to a small percentage of users, monitor, then expand
  - safer, controlled exposure
- A/B testing: route users to different app versions to compare
  - analyze which version performs better (e.g. UX, business metrics)

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

## troubleshooting scenarios

- laravel app isn't accessible to other services, how to restore accessibility?
  - check if container/aws resource is running
  - check security groups
  - check nginx/apache config
  - check network ACLs (?)
  - check if bound to correct port/IP
  - check env vars

- how do you check if a container is reachable?
  - [curl](/cli_tools.md#curl)
  - `ping`: checks if a host is reachable
  - `nc`: checks if a port is open
    - scan port 80: `nc -zv localhost 80`
  - `ss`
  - `netstat`
