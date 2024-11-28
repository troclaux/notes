
# ansible

- playbook header: the playbook starts with a header that specifies the target hosts and any additional settings
- plays: a play is a set of tasks that will be executed on a set of hosts
  - a playbook can contain one or more plays
- tasks: a task is a single unit of work to be executed on the target hosts
  - tasks are written in a sequence and are executed one after the other
- modules: each task uses a specific ansible module to perform actions like installing packages, copying files, or starting services

```yml
---
- name: Ensure a web server is installed and running
  hosts: webservers  # This is the group of hosts to target (defined in your inventory)
  become: true  # Run tasks as root user (sudo)

  tasks:
    - name: Install Apache web server
      apt:  # Ansible module for managing packages on Debian-based systems
        name: apache2
        state: present  # Ensure the package is installed

    - name: Start Apache service
      service:  # Ansible module for managing services
        name: apache2
        state: started  # Ensure the service is running

    - name: Ensure the web server is enabled to start on boot
      service:
        name: apache2
        enabled: true  # Ensure the service starts on boot
```

- playbook header `---`: this marks the beginning of the yaml file
  - not strictly necessary but it's a good practice
- `name`: describes playbook in human-readable form
- `hosts`: specifies target machines
- `become`: if set to `true`, runs tasks with elevated privileges (like `sudo`)
- `tasks`: contains list of tasks to be executed on the hosts
- `modules`: pre-built units of code that perform specific tasks

modules examples:

- `package/apt/dnf`: package management

```yaml
- name: Install more neovim requirements for Fedora
  dnf:
    name:
      - gcc
      - gcc-c++
      - lua
  tags: nvim
  when: ansible_pkg_mgr == 'dnf'
```

- `copy`: copy files

```yaml
- copy: src=/local/file dest=/remote/path
```

- `file`: manage files/directories

```yaml
- name: Remove .config/nvim folder
  ansible.builtin.file:
    path: "{{ lookup('env', 'HOME') }}/.config/nvim"
    state: absent
  tags: nvim
```

- variables: store and reuse values throughout your playbook

```yaml
vars:
  app_port: 8080
  app_path: /var/www/app

tasks:
  - name: Copy app config
    template:
      src: config.j2
      dest: "{{ app_path }}/config.conf"
```

- tags: label tasks to run or skip specific parts of a playbook

```yaml
tasks:
  - name: Install stow
    package:
      name: stow
      state: present
    tags:
      - stow
      - gitconfig
      - zshrc
      - tmux

  - name: Install Docker
    apt: name=docker state=present
    tags: docker
```

- conditionals: execute tasks based on specific conditions

```yaml
tasks:
  - name: Install Apache on Debian
    apt: name=apache2 state=present
    when: ansible_os_family == "Debian"

  - name: Install Apache on RedHat
    dnf: name=httpd state=present
    when: ansible_os_family == "RedHat"
```

- run ansible-playbook: `ansible-playbook -K --tags "zsh,kitty" --skip-tags "sway" playbook.yml`
  - `-K`: ask for privilege escalation password
    - required to run on machines with login password
  - `--tags`: runs only the tasks with the specified tags
  - `--skip-tags`: skip tasks with the following tags
