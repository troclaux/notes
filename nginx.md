
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

- can decrypt https requests and forward them to backend
- to enable https to a website, it's required an ssl certificate
  - how to configure ssl certificate:
    - getting a certificate, issued by Let's Encrypt (Certbot) or other trusted authority
    - installing it on your web server (e.g. nginx, apache)
    - configuring the web server (with nginx.conf) to use the certificate for secure connections

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
  - [example](./code/nginx.conf)

## basic concepts

- master process:
  - manages the server
  - reads the configuration
  - spawns worker processes
  - doesn't handle requests itself
- worker processes:
  - handle client connections and requests
  - each worker can manage thousands of connections

- directives: key-value pairs or standalone commands that define how nginx behaves
- blocks: group of directives
  - provides scope to directives
  - e.g. `events { ... }`, `http { ... }`, `server { ... }`, `location { ... }`
  - can be nested (e.g. `location` inside `server` block)
- variables
  - e.g. `$uri` is a variable:`try_files $uri /index.php;`

## Installation

```bash
sudo apt update && sudo apt upgrade -y
```

```bash
sudo apt install nginx -y
```

To check if the installation was successful, run the following command:

```bash
sudo systemctl status nginx

# OUTPUT:

# ● nginx.service - A high performance web server and a reverse proxy server
#      Loaded: loaded (/lib/systemd/system/nginx.service; enabled; vendor preset: enabled)
#      Active: active (running)
```

If the status says `running`, then you're good to go. Otherwise you may start the service by executing this command:

```bash
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

```bash
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

```bash
sudo systemctl restart nginx
```

to send a signal to the NGINX master process, run the following command:
- in the example below, the signal that is being sent is `reload`
```bash
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

> key-value pairs or standalone commands that define how nginx behaves

example:

```nginx
events {
    worker_connections 1024;
}

http {
    server {
        listen 80;
        server_name example.com;
        root /var/www/html;
        index index.html;
    }
}
```

- `listen 80;`: listen to port 80
- `root /var/www/html;`: serves files from `/var/www/html`
- `index index.html;`: returns `index.html` for `/`

another example:

```nginx
worker_processes auto;

events {
    worker_connections 1024;
}

http {
    server {
        listen 80;
        location / {
            proxy_pass http://nextjs-service:3000;
            proxy_http_version 1.1;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection "upgrade";
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }
    }
}
```

- `worker_processes auto;`: automatically determine the number of worker processes based on available cpu cores
- `events`: controls how nginx handles connections
  - `worker_connections 1024;`: each worker process can handle up to 1024 simultaneous connections
- `http`: 
  - `server`: 
  - `location /`: configuration for all requests
    - `proxy_pass http://nextjs-service:3000;`: forward all incoming requests to the backend service running in `http://nextjs-service:3000;`
    - `proxy_http_version 1.1;`: uses http version 1.1
    - `proxy_set_header Upgrade $http_upgrade;`: 
    - `proxy_set_header Connection "upgrade";`: 
    - `proxy_set_header host $host;`: this forwards the original host header from the client to the backend, preserving the intended hostname
    - `proxy_set_header X-Real-IP $remote_addr;`: forwards the real ip address of the client to the backend
      - useful for logging and application logic on the backend
    - `proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;`: 
    - `proxy_set_header X-Forwarded-Proto $scheme;`: 
      - `$scheme`: built-in variable that represents the protocol (http or https) of the original request

### directives for begginers

- `listen`: port and protocol (e.g. `listen 80` or `listen 443 ssl`)
- `server_name`: domains handled by this server
- `root`: directory for static files (e.g. `/var/www/html`)
- `location`: rules for specific url paths
- `try_files`: attempts to serve files or fall back (e.g. `try_files $uri /index.php`)
- `proxy_pass`: forwards requests to a backend (e.g., `proxy_pass http://127.0.0.1:8000;`)
- `fastcgi_pass`: sends php requests to php-fpm (e.g., `fastcgi_pass unix:/var/run/php-fpm.sock;`)

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


