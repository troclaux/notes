
# node.js

> runtime environment that allows you to run javascript on the server side

- built on chrome's V8 engine

- event-driven
- asynchronous: node.js doesn't wait for operations to finish before moving on to the next operation
  - operation: reading a file, making a HTTP request or querying a database
  - asynchronous code is handled using callbacks, promises or async/await
- non-blocking I/O
- single-threaded
- cross-platform: windows, linux, macOS

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

