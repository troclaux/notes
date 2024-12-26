
# node.js

> runtime environment that allows you to run javascript on the server side

- built on chrome's V8 engine

- event-driven
  - examples:
    - HTTP server receiving a request
    - file system completing a read operation
    - user clicking a button in a web application
    - database query finishing execution
    - timer expiring (setTimeout/setInterval)
- asynchronous: node.js doesn't wait for operations to finish before moving on to the next operation
  - operation: reading a file, answering a HTTP request or querying a database
  - asynchronous code is handled using callbacks, promises or async/await
- non-blocking I/O: node.js can handle multiple operations simultaneously without waiting for each one to complete
  - when an I/O operation starts (like reading a file), node.js continues executing other code
  - when the I/O operation completes, a callback function is triggered to handle the result
  - example: while reading a large file, node.js can still process incoming HTTP requests
- single-threaded: one thread of execution
  - this thread is responsible for executing:
    - event loop: runs continuously, checking for events and executing callbacks or event handlers as needed
    - v8 javascript engine
- cross-platform: runs on multiple operating systems, including windows, linux, macOS

## installation

- you can install node.js:
  - from the official [website](https://nodejs.org/en/)
  - or with a node version manager (recommended):
    - most popular: [nvm](https://github.com/nvm-sh/nvm)
    - recommended:[volta](https://volta.sh)

```bash
curl https://get.volta.sh | bash -s -- --skip-setup
```

after installing volta, install node.js with:

```bash
volta install node
```

## basic commands

- to run javascript code: `node code.js`
- to install a package: `node install express`
- 

## file system module

## event-driven architecture

