
## apache web server

> free and open-source web server

- multiplatform: linux, windows, Unix, FreeBSD, etc
- uses daemon `httpd`
- `apachectl`: front end for apache server
  - allows you to:
    - change environment variables

- default directories:
  - Document Root: `"/var/www/html"`
  - config files:
    - `/etc/apache2/` or `/etc/httpd/`
    - `.htaccess`

- types of pages
  - root page: page that is displayed when user enters the website
    - DocumentRoot: configuration line that defines the localization of root page
      - `DocumentRoot "/var/www/html"`: apache will look for the root page in this path
  - subpages

example of directory structure if DocumentRoot is set to `"/var/www/html"`:

```
/var/
  └── www/
      └── html/
          ├── index.html
          ├── about.html
          └── images/
              └── logo.png
```

- visiting `http://yourserver/` will display `index.html`
- visiting `http://yourserver/about.html` will display `about.html`
- visiting `http://yourserver/images/logo.png` will serve `logo.png`

- command to start httpd in linux distributions (`start` can be replaced with `restart`, `stop`, etc):
  - red hat: `service httpd start;`
  - ubuntu: `/etc/init.d/apache2 start`

- start apache: `sudo systemctl start apache2`
- stop apache: `sudo systemctl stop apache2`
- start apache: `sudo systemctl restart apache2`

### parameters

to learn more about parameters, use the terminal command `$ man httpd`:
- -d serveroot: 
- -f config
- -k start|restart|graceful|stop|graceful-stop: signals httpd to start, restart or stop
- configtest:

### config files

- `httpd.conf` or `apache2.conf`: main config file
  - default path: `/etc/httpd/conf` or `/etc/apache2`
- `.htaccess`: files that define configurations for each directory
- `srm.conf`: contains settings that define the behaviour of SRM (Storage Resource Management) software
  - default path: `/etc/srm`
  - defines behaviour for storage systems, for example:
    - paths, capacity and usage policies
- `access.conf`: contains settings that control document access

example:

```apache
# Example Apache configuration snippet
Listen 80                   # Port Apache listens to (default: 80 for HTTP)
ServerName www.example.com  # Domain name

DocumentRoot "/var/www/html"
```

virtual hosts: allow apache to serve multiple websites/domains from one server

```apache
# ubuntu/debian: /etc/apache2/sites-available/example.com.conf
# centos/rhel: /etc/httpd/conf.d/example.com.conf

<VirtualHost *:80>
  ServerName example.com
  ServerAlias www.example.com
  DocumentRoot "/var/www/example"
</VirtualHost>
```

### modules

> elements that extend the functionality of Apache HTTP server, allowing it to handle tasks beyond basic web serving

- types of modules:
  - static
    - compiled into apache HTTP server binary build time
    - cannot be loaded or unloaded dynamically without recompiling the server
  - dynamic
    - loaded into apache HTTP server at runtime, without needing to recompile the server
    - can be loaded or unloaded using configuration directives

- core
  - manages server configuration
  - provides HTTP request and responde handling
- mpm_common: 
- worker: 
- event: 
- prefork: 
- mpm_netware: 
- mpmt_os2: 
- mpm_winnt: 

### sections and directives
### log files
### testing apache server
