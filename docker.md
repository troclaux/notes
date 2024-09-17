
# docker

## Image

## Volume

## Container

```shell
docker container run -d -p 8080:80 --name webhost nginx
```
- -d: detached mode
- -p: port mapping
  - format: host-port:container-port
    - my host machine that is running docker: 8080
    - container: 80
- --name: name of the container
- nginx: image to use
  - pulls the image from docker hub if it is not available locally

## Exec
