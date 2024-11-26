
# docker

> platform and toolset for developing, shipping, and running applications inside containers

- containers: lightweight, portable, and isolated environments that package an application and all its dependencies so that it can run consistently across different computing environments
  - dependencies: libraries, configurations, and runtime

- docker uses namespaces and cgroups to provide a lightweight/secure/efficient containerization platform
  - namespace: linux feature that isolates resources of a process, to simulate that a container runs in a separate system
  - cgroups: allocate/limit resources (CPU, memory, I/O, etc) for processes

- docker engine: core component that runs and manages containers
- docker daemon (dockerd): background process that manages docker containers, images and networks
- docker CLI: allows users to interact with docker with shell commands

## images

> lightweight standalone package that contains everything necessary to create containers

- read-only snapshot of the code, libraries, dependencies and configuration to run containers
- template to create multiple containers
- blueprint for building a container

- `Dockerfile`: text file containing instructions for building a Docker image

Dockerfile example:

```
# Use an official Ubuntu image as the base
FROM ubuntu:latest

# Set the working directory to /app
WORKDIR /app

# Copy the web server configuration file
COPY nginx.conf /etc/nginx/nginx.conf

# Copy the web server files
COPY . /app

# Install the necessary dependencies
RUN apt-get update && apt-get install -y nginx

# Expose the port
EXPOSE 80

# Run the command to start the web server
CMD ["nginx", "-g", "daemon off;"]
```

better keywords:

- `ENV MY_VAR="Hello World"`: define container's environment variables
- `EXPOSE 80`: define which ports docker container will listen to
- `COPY . /app`: copy files from host machine's path `.` to container's `/app` directory
- `ENTRYPOINT ["node", "app.js"]`: run command at container's startup
- `CMD ["nginx", "-g", "daemon off;"]`: run command at container's startup if there isn't `ENTRYPOINT`
- `RUN`: triggers while docker builds the image
  - for multiple line commands, use `\` and `&&` to guarantee order of execution

```
RUN <bash_command1> \
&& <bash_command2> \
&& <bash_command3>
```

> [!NOTE]
> `ENTRYPOINT` vs `CMD`
> `ENTRYPOINT`: command that will run at container's startup, cannot be easily overridden by CLI arguments (`docker run`)
> `CMD`: command to run if no `ENTRYPOINT` is defined, can be easily overridden by CLI arguments

- Dockerfile layer optimization best practices:
  - put the stuff that changes less on top of the docker file to speed up the build command
  - following the same logic, put the stuff that changes most on the bottom of the dockerfile

### CLI

- build image: `docker build -t my-image .`
  - `-t my-image`: tag image with a name
  - `.`: Dockerfile's path
- list all images stored locally: `docker images`
- remove docker image from your local machine: `docker rmi my-image`
  - `-f`: force remove if image is being used by stopped containers
- remove unused images: `docker image prune`
- clean up everything: `docker system prune`
  - stopped containers, unused networks, dangling images, etc

## containers

> self-contained environment that runs an application and its dependencies in isolation from the host machine

- runtime instance of a docker image
- has its own file system, network stack and process space

```bash
docker container run -d -p 8080:80 --name webhost nginx
```

- `-d`: run container in detached mode
- `-p`: port mapping
  - `8080:80` == `<host-port>:<container-port>`
    - my host machine that is running docker: 8080
    - container: 80
- `--name`: name of the container
- `-e MY_ENV_VAR=true`: define environment variable
- `nginx`: image to use
  - pulls the image from docker hub if it is not available locally

```bash
docker run --rm -it -v ~/dotfiles:/root/dotfiles my-ubuntu-image bash
```

- `--rm`: remove container when it stops
- `-it`: attach a terminal to the container
- `-v ~/dotfiles:/root/dotfiles`: mounts local directory from host machine that's running docker to container's filesystem
  - `~/dotfiles`: source directory's path for the [bind mount](#bind-mounts)
  - `/root/dotfiles`: directory inside container where source directory will be mounted
- `my-ubuntu-image`: name of the image being used to create the container
- `bash`: command that will run inside the container
- run command inside running container: `docker exec -it my-container bash`
- stop running container: `docker stop my-container`
- start stopped container: `docker start my-container`
- remove container: `docker rm my-container`
- print container logs: `docker logs my-container`

> [!IMPORTANT]
> docker allows you to reorder the options as long as the last arguments are the image name (`my-ubuntu-image`) and the command (`bash`)
> `-v`, `-it` and `--rm` can change order without changing the resulting behavior as long as they are before the image name

### container lifecycle

- possible states in docker container lifecycle:
  - created: container has been created from an image but not started
  - running: container is running with all its processes
  - paused: container whose processes have been temporarily paused without stopping it completely
    - resources (CPU and memory) used by container are still in use
  - stopped: container is shut down
    - frees resources used by container
  - deleted: container is removed from the system

### healthchecks

TODO

- 3 possible states after healthcheck:
  - starting
  - healthy
  - unhealthy

## exec

> run a command in a running container

```bash
docker exec -it my-container bash
```

- `-i`: keep STDIN open even if not attached
  - forces docker to keep bash running
  - without it, when container would run `bash` and immediately close because there's no input stream connected
- `-t`: emulate a terminal
  - without a pseudo-TTY (Terminal TeleType `-t`), you would get a more basic, non-interactive interface that might not properly display command prompts
- `my-container`: name or ID of the container
- `bash`: command to execute inside container

## persistent data

> data that persists even after the container is deleted or recreated

Two main ways to persist data:

### Volumes

- Managed by Docker
- Stored in `/var/lib/docker/volumes/` on host
- Better portability and backup
- Can be shared between containers
- Volume drivers allow storing on remote hosts/cloud providers
- Easier to backup and migrate

```bash
# Create volume
docker volume create myvolume

# Mount volume
docker run -v myvolume:/app myimage
```

### Bind Mounts

- Store anywhere on host system
- Less functionality than volumes
- Good for development to mount source code
- Host machine path must exist

```bash
# Mount host directory
docker run -v /host/path:/container/path myimage
```
- print docker volumes: `docker volume ls`
- inspect volume: `docker volume inspect myvolume`
- remove volume: `docker volume rm myvolume`
- remove unused volumes: `docker volume prune`

## networks

> allow multiple Docker containers to communicate and share resources with each other

- create docker network: `docker network create mynetwork`
- print docker networks: `docker network ls`
  - `-a`: show stopped containers
- remove docker network: `docker network rm mynetwork`
- remove unused docker networks: `docker network prune`
- inspect docker network: `docker network inspect mynetwork`

## docker-compose

- start services defined in file with name `docker-compose.yml`: `docker-compose up`
- stop services defined in file with name `docker-compose.yml`: `docker-compose down`

