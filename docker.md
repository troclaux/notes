
# docker

> platform and toolset for developing, shipping, and running applications inside containers

- containers: lightweight, portable, and isolated environments that package an application and all its dependencies so that it can run consistently across different computing environments
  - dependencies: libraries, configurations, and runtime

- benefits:
  - quickly setup developer environment
  - application isolation
  - reproducibility
  - security: isolation between containers and host system reduces the risk for system-wide compromises
  - IaC (Infrastructure as Code)
    - easy to reproduce environment
    - allows version control for environments
      - if new environment has issues, you can rollback to previous version
  - CI/CD
    - faster build setup
    - faster testing
    - faster deployment

- docker uses namespaces and cgroups to provide a lightweight/secure/efficient containerization platform
  - namespace: linux feature that isolates resources of a process, to simulate that a container runs in a separate system
  - cgroups: allocate/limit resources (CPU, memory, I/O, etc) for processes

- docker engine: core component that runs and manages containers
- docker daemon (dockerd): background process that listen to requests and manages docker containers, images and networks
- docker CLI: allows users to interact with docker with shell commands
- dockerhub: official cloud service for storing and sharing docker images

- [containers vs virtual machines](https://www.atlassian.com/microservices/cloud-computing/containers-vs-vms)
  - virtual machines virtualize an entire machine down to the hardware layers
  - containers only virtualize software layers above the operating system level

## CLI commands

- pull specific docker image: `docker pull node:18-alpine`
- list available images: `docker images`
- remove an image: `docker rmi nginx`
- remove unused images: `docker image prune`
- run container: `docker run -d -p 8080:80 nginx`
  - `-d`: run container with detached shell
  - `-p`: expose port 80 in the container on port 8080 on your machine
- stop container: `docker stop <container_id>`
- restart container: `docker restart <container_id>`
- run container with interactive shell: `docker run -it ubuntu bash`
- list running containers: `docker ps`
- list all containers: `docker ps -a`
- get logs of a running container: `docker logs <container_id>`
  - `-f`: shows live logs (useful for debugging)
- remove container: `docker rm <container_id>`
- remove all stopped containers: `docker container prune`
- create volume: `docker volume create myvolume`
- list volumes: `docker volume ls`
- remove volume: `docker volume rm myvolume`
- remove all unused volumes: `docker volume prune`
- list docker networks: `docker network ls`
- create new docker network: `docker network create mynetwork`
- remove docker network: `docker network rm mynetwork`

- start services from `docker-compose.yml`: `docker-compose up -d`
- stop services: `docker-compose down`
- restart services: `docker-compose restart`
- show real-time logs: `docker-compose logs -f`
- check logs of a specific service: `docker-compose logs nginx`

## images

> lightweight standalone package that contains everything necessary to create containers

- read-only snapshot of the code, libraries, dependencies and configuration to run containers
- template to create multiple containers
- blueprint for building a container

- `Dockerfile`: text file containing instructions for building a Docker image

Dockerfile example:

```dockerfile
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

- `ENV MY_VAR="Hello World"`: define container's environment variables
- `EXPOSE 80`: define which ports docker container will listen to
  - doesn't actually publish or open the ports to the host system
  - serves only as documentation
  - the port won't be accessible from the outside unless explicitly published with the `-p` flag when running the shell command
- `COPY . /app`: copy files from host machine's path `.` to container's `/app` directory
- `ENTRYPOINT ["/bin/bash"]`: default command to run when container is started
- `CMD ["nginx", "-g", "daemon off;"]`: default command-line arguments to pass to the `ENTRYPOINT` command
  - can be overridden when running a CLI command
  - e.g. if you run `docker run -it my-image bash`, it will override the `CMD` above and the container **won't** run nginx
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

- build image: `docker build --tag imagename:imagetag .`
  - `--tag imagename:imagetag`: tags image with the name `imagename` and `imagetag` tag
    - `docker tag <image-name> <new-tag>`
      - `<image-name>`: existing image that will be tagged
      - `<new-tag>`: new tag that will reference image
  - `.`: Dockerfile's path
- list all images stored locally: `docker images`
- remove 1 or more image: `docker rmi my-image`
  - `-f`: force remove if image is being used by stopped containers
- remove unused images: `docker image prune`
- clean up everything: `docker system prune`
  - stopped containers, unused networks, dangling images, etc

- to see containers' ID, name, CPU, Memory usage, NET I/O, BLOCK I/O and PIDS: `docker stats`

### dockerignore

> specifies which files and directories Docker should ignore when building images

- similar to `.gitignore`
- prevents unnecessary files from being copied into the image during build
- helps reduce image size and build time
- improves security by excluding sensitive files

Example `.dockerignore`:

```dockerignore
# Version control
.git
.gitignore

# Dependencies
node_modules
vendor

# Build files
dist
build

# Logs and temp files
*.log
.tmp

# Development files
.env
.env.local
*.md
```

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

> [!NOTE]
> `--name` vs `--tag`
> `--name` is used for containers
> `--tag` is used for images

```bash
docker run --rm -it --network my-network -v ~/dotfiles:/root/dotfiles my-ubuntu-image bash
```

- `--rm`: remove container when it stops
- `-it`: attach a terminal to the container
- `--network my-network`: connects container to `my-network` network
- `-v ~/dotfiles:/root/dotfiles`: mounts local directory from host machine that's running docker to container's filesystem
  - `~/dotfiles`: source directory's path for the [bind mount](#bind-mounts)
  - `/root/dotfiles`: directory inside container where source directory will be mounted
- `my-ubuntu-image`: name of the image being used to create the container
- `bash`: command that will run inside the container

- list all containers (including stopped ones): `docker ps -a`
- run command inside running container: `docker exec -it my-container bash`
- start stopped container: `docker start my-container`
- stop running container: `docker stop my-container`
- remove 1 or more stopped container: `docker rm my-container`
- remove 1 or more running containers: `docker rm -f my-container`
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

> [!IMPORTANT]
> docker cannot remove paused containers
> containers have to be stopped can be removed

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

> [!IMPORTANT]
> `-v myvolume:/app` behavior:
> if `myvolume` exists: binds existing volume to `/app` in container
> if `myvolume` doesn't exist: creates new volume named `myvolume` and binds it to `/app`
> this is different from bind mounts where the host path must exist beforehand

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

## docker hub

> docker's official cloud-based registry for storing and sharing docker images

- requires free account at hub.docker.com
- image naming convention: `username/image-name:tag`
  - example: `troclaux/my-go-server:1.0`
  - if no tag is specified, defaults to `latest`
  - use [semantic versioning](https://semver.org/) and also push to the "latest" tag
    - `docker build -t username/imagename:0.0.0 -t username/imagename:latest .`
    - `docker push username/imagename --all-tags`
      - `--all-tags`: ensures all tags of the image are pushed, not just the default tag (usually latest)

common commands:

```bash
# Login to Docker Hub
docker login

# Tag local image for Docker Hub
docker tag local-image:tag username/repository:tag

# Push image to Docker Hub
docker push username/repository:tag

# Pull image from Docker Hub
docker pull username/repository:tag
```

best practices:
- always tag your images with specific versions
- test your image locally before pushing
- use meaningful repository names
- add a README.md to your Docker Hub repository
- scan your images for security vulnerabilities before pushing

> [!NOTE]
> Public repositories are free
> Private repositories may require paid subscription

## docker-compose

- start services defined in file with name `docker-compose.yml`: `docker-compose up`
  - `--build`: forces docker to rebuild images before starting containers
  - `-d`: runs the containers in detached mode (in the background)
- stop services defined in file with name `docker-compose.yml`: `docker-compose down`
- runs `certbot` container once and then removes it after execution: `docker-compose run --rm certbot`
  - `--rm`: removes container after execution
  - requests SSL certificate from Let's Encrypt
- restart container: `docker-compose restart nginx`
- renew SSL certificate: `docker-compose run --rm certbot renew`
  - certificates from Let's Encrypt expire every 90 days


```yaml
services:
  nextjs-service:
    image: 072216710152.dkr.ecr.sa-east-1.amazonaws.com/peso-repo:nextjs-latest
    container_name: nextjs_app
    restart: always
    env_file:
      - .env.local
    environment:
      - NODE_ENV=production
    expose:
      - "3000"
    networks:
      - peso_network
```

- `nextjs-service`: service name
  - use this name within `docker-compose.yml` file to reference this service (like in `depends_on` or networks)
- `nextjs_app`: container name
  - this is the name that appears when you run docker CLI commands (e.g. `docker ps`, `docker logs`)
