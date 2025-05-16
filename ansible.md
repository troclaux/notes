
# ansible

## architecture

- control node: executes playbooks and connects to the managed nodes
  - where ansible is installed and run
- managed nodes (hosts): servers that are configured by ansible
  - don't need ansible installed, only control node requires ansible
- inventory: list of managed nodes
  - static: plain text file (`ini`, `yaml` or `json`) with hardcoded IPs/hostnames
  - dynamic: used when infrastructure is dynamic (e.g. AWS, azure, GCP or kubernetes)
    - instead of writing hostnames manually, ansible pulls them from APIs using plugins or scripts
- playbooks: set of instructions that define how a managed node will be configured/managed
- modules: built-in instructions that perform tasks
- roles: structured, reusable unit of ansible content
- collection: packaged directory of ansible content, can contain (roles, modules, plugins, playbooks, documentation, tests)
- ansible galaxy: hub for finding, reusing and sharing ansible content, such as roles, collections and playbooks
  - like npm for node.js or pip for python, but for ansible
  - command example: `ansible-galaxy install geerlingguy.nginx`

## cli

- run ansible-playbook: `ansible-playbook -K --tags "zsh,kitty" --skip-tags "sway" playbook.yml`
  - `-K`: ask for privilege escalation password
    - required to run on machines with login password
  - `--tags`: runs only the tasks with the specified tags
  - `--skip-tags`: skip tasks with the following tags

## playbook

- playbook header: the playbook starts with a header that specifies the target hosts and any additional settings
- plays: a play is a set of tasks that will be executed on a set of hosts
  - a playbook can contain one or more plays
- tasks: a task is a single unit of work to be executed on the target hosts
  - tasks are written in a sequence and are executed one after the other
- modules: each task uses a specific ansible module to perform actions like installing packages, copying files, or starting services
- agentless: no need to install any software on the target machines that ansible manages

```yml
---
- name: Ensure a web server is installed and running
  hosts: webservers  # Play applies to this group of hosts (defined in your inventory)
  become: true  # Run tasks as root user (sudo)

  tasks:          # List of tasks
    - name: Install Apache web server
      apt:  # Ansible module for managing packages on Debian-based systems
        name: apache2
        state: present  # Ensure the package is installed

    - name: Start Apache service
      service:  # Ansible module for managing services
        name: apache2
        state: started  # Ensure the service is running now

    - name: Ensure the web server is enabled to start on boot
      service:
        name: apache2
        enabled: true  # Ensure the service will automatically start on boot
```

- playbook header `---` (optional): this marks the beginning of the YAML file
  - it's optional but recommended for clarity and compatibility
  - not strictly necessary but it's a good practice
- `name`: describes playbook in human-readable form
- `hosts`: specifies target machines
- `become`: if set to `true`, runs tasks with elevated privileges (like `sudo`)
- `tasks`: contains list of tasks to be executed on the hosts
- `modules`: pre-built units of code that perform specific tasks
- `service`: ansible module that manages services
  - `state`: defines desired condition of a resource, when used with `service`, these are the possible values:
    - `started`: start the service if it's not running
    - `stopped`: stop the service if it's running
    - `restarted`: restart the service (regardless of whether it's running or not)
    - `reloaded`: reload the service (useful if config files changed but restart not needed)
    - `enabled`: ensure the service is enabled at boot
    - `disabled`: ensure the service is disabled at boot

## modules

> built-in instructions that perform tasks

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
- ansible.builtin.copy:
  src: /local/file
  dest: /remote/path
```

- `file`: manage files/directories

```yaml
- name: Remove .config/nvim folder
  ansible.builtin.file:
    path: "{{ lookup('env', 'HOME') }}/.config/nvim"
    state: absent
  tags: nvim
```

## tags

> run only specific parts of a playbook

- running tagged tasks only: `ansible-playbook site.yml --tags "install, start"`
- skipping tags: `ansible-playbook site.yml --skip-tags "start"`
- list all tags in a playbook: `ansible-playbook playbook.yml --list-tags`

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
    apt:
      name: docker
      state: present
```

## variables

> store and reuse values throughout playbooks, tasks, roles, etc

```yaml
---
- name: Install Apache
  hosts: web
  vars:
    apache_package: apache2
  tasks:
    - name: Install Apache package
      apt:
        name: "{{ apache_package }}"
        state: present
```

here, `{{ apache_package }}` will be replaced with "apache2" when the playbook runs

## conditionals

> execute tasks based on specific conditions

- conditionals can use:
  - variables
  - registered results
  - logical expressions (`and`, `or`, `not`)
  - `in`, `is defined`, etc

```yaml
- name: Only run if OS is Ubuntu
  apt:
    name: nginx
    state: present
  when: ansible_facts['os_family'] == "Debian"
```

```yaml
- name: Check if a file exists
  stat:
    path: /etc/myconfig.conf
  register: file_status

- name: Print message if file exists
  debug:
    msg: "Config file is present"
  when: file_status.stat.exists
```

## common tasks

remove a directory:

```yaml
- name: Deletes ~/.config/nvim
  ansible.builtin.file:
    path: "{{ lookup('env', 'HOME') }}/.config/nvim"
    state: absent
  tags: nvim
```

ensure a directory exists:

```yaml
- name: Creates ~/.config/wofi
  ansible.builtin.file:
    path: "{{ lookup('env', 'HOME') }}/.config/wofi"
    state: directory
  tags: sway
```
