
## Properties

![Forward Proxy vs Reverse Proxy](./images/proxy.jpg)

- mainly used as reverse proxy
  - sits in front of one or more web servers and acts on behalf of these servers when responding to client requests
  - has many use cases, such as:
    1. load balancer
      - distributes client requests across multiple servers
    2. cache
      - caches the server's responses to future clients
    3. extra layer of security
      - conceals internal infrastructure
      - traffic filtering
      - SSL/TLS authentication
    4. routing
      - route requests to different servers based on:
        - request's domain
        - request's URL
        - request's HTTP header
- can also be web server

- asynchronous
- non-blocking
- event-driven
  - producers and consumers of events

- high performance
- high concurrency
- low memory usage
- small footprint
- high scalability

- doesn't depend on threads to handle requests
  - uses a single-threaded event-driven architecture instead

- composed by:
  - main process
    - reads configuration files
    - manages worker processes
  - worker processes
    - handle requests
    - each worker process can handle thousands of connections

- configuration file
  - `/etc/nginx/nginx.conf`


## Installation

```shell
sudo apt update && sudo apt upgrade -y
```

```shell
sudo apt install nginx -y
```

To check if the installation was successful, run the following command:

```shell
sudo systemctl status nginx

# OUTPUT:

# ● nginx.service - A high performance web server and a reverse proxy server
#      Loaded: loaded (/lib/systemd/system/nginx.service; enabled; vendor preset: enabled)
#      Active: active (running)
```

If the status says `running`, then you're good to go. Otherwise you may start the service by executing this command:

```shell
sudo systemctl start nginx
```

## NGINX's Configuration Files

- NGINX's configuration files end with `.conf`
- located at
  - `/etc/nginx`
  - `/usr/local/nginx/conf`
  - `/usr/local/etc/nginx`

- `nginx.conf`
  - main configuration file
  - there are multiple configuration files
  - can be modified to suit your needs

```shell
cd /etc/nginx

ls -lh

# OUTPUT:

# drwxr-xr-x 2 root root 4.0K Apr 21  2020 conf.d
# -rw-r--r-- 1 root root 1.1K Feb  4  2019 fastcgi.conf
# -rw-r--r-- 1 root root 1007 Feb  4  2019 fastcgi_params
# -rw-r--r-- 1 root root 2.8K Feb  4  2019 koi-utf
# -rw-r--r-- 1 root root 2.2K Feb  4  2019 koi-win
# -rw-r--r-- 1 root root 3.9K Feb  4  2019 mime.types
# drwxr-xr-x 2 root root 4.0K Apr 21  2020 modules-available
# drwxr-xr-x 2 root root 4.0K Apr 17 14:42 modules-enabled
# -rw-r--r-- 1 root root 1.5K Feb  4  2019 nginx.conf
# -rw-r--r-- 1 root root  180 Feb  4  2019 proxy_params
# -rw-r--r-- 1 root root  636 Feb  4  2019 scgi_params
# drwxr-xr-x 2 root root 4.0K Apr 17 14:42 sites-available
# drwxr-xr-x 2 root root 4.0K Apr 17 14:42 sites-enabled
# drwxr-xr-x 2 root root 4.0K Apr 17 14:42 snippets
# -rw-r--r-- 1 root root  664 Feb  4  2019 uwsgi_params
# -rw-r--r-- 1 root root 3.0K Feb  4  2019 win-utf
```

run the following command to restart the NGINX service:

```shell
sudo systemctl restart nginx
```

to send a signal to the NGINX master process, run the following command:
- in the example below, the signal that is being sent is `reload`
```shell
sudo nginx -s reload
```

`-s` flag is used to send various signals to a master process, such as:
- `stop`
  - fast shutdown
- `quit`
  - graceful shutdown
- `reload`
  - reloading the configuration file
- `reopen`
  - reopening the log files


## Directives

```conf
events {

}

http {
    server {
        listen 80;
        server_name nginx-handbook.test;
        return 200 "Bonjour, mon ami!\n";
    }
}
```

## Configuration file structure

composed by modules
- each module is controlled directives
- directives specifications are defined by nginx.conf

There are two types of directives in NGINX:
1. simple directives (terminated by a semicolon)
  - directive name
  - space delimited parameters, such as:
    - `listen`
    - `return`
2. block directives (terminated by curly braces)
  - Context: A block directive capable of containing other directives inside it
    - `events`
      - global configuration that defines how NGINX should handle requests
      - only one `events` block is allowed in the configuration file
    - `http`
      - contains all the configuration related to HTTP or HTTPS requests
      - contains `server` blocks
      - only one `http` block is allowed in the configuration file
      - `upstream`
        - defines a group of servers that NGINX can send proxy requests to
    - `server`
      - defines the configuration for a virtual server
      - contained within the `http` block
      - There can be multiple server contexts in a valid configuration file
      - `location`
        - defines how NGINX should handle requests for a specific URL
        - contained within the `server` block
    - `main`
      - contains the main configuration settings for NGINX
      - only one `main` block is allowed in the configuration file
      - contains `events` and `http`

Example of a `upstream` block:
```conf
http {
  upstream test {
    # uses round-robin by default

    server test1.exemple.com;
    server test2.exemple.com;
    server test3.exemple.com;
  }
}
```

Example of a `proxy_pass` block:
```conf
http {
  # ...

  server {
    listen 80;
    location / {
      proxy_pass http://test;
    }
  }

  # ...
}
```
in the example above:
- NGINX listens to port 80
- forwards requests to the `test` upstream block
- proxy_pass directive is used to specify the URL to which the requests should be forwarded to


