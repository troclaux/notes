
# node.js

> open-source javascript runtime environment

- built on chrome's V8 engine
- cross-platform
- JavaScript Runtime Environment (JSRE): environment which provides all the necessary components in order to use and run a javascript program
  - JSRE = javascript engine + extra features (API, libraries that javascript alone doesn't have)
  - javascript engine: program that reads, interprets and executes javascript code
    - e.g. v8 (chrome, node.js), SpiderMonkey (firefox), JavaScriptCore (safari)
    - v8 engine is written in C++

- event-driven
- single-threaded: one thread of execution
  - this thread is responsible for executing:
    - event loop: runs continuously, checking for events and executing callbacks or event handlers as needed
    - v8 javascript engine
  - examples of events:
    - HTTP server receiving a request
    - file system completing a read operation
    - database query finishing execution
    - timer expiring (setTimeout/setInterval)
- asynchronous: node.js doesn't wait for operations to finish before moving on to the next operation
  - operation: reading a file, answering a HTTP request or querying a database
  - asynchronous code is handled using callbacks, promises or async/await
- non-blocking I/O: node.js can handle multiple operations simultaneously without waiting for each one to complete
  - when an I/O operation starts (like reading a file), node.js continues executing other code
  - when the I/O operation completes, a callback function is triggered to handle the result
  - example: while reading a large file, node.js can still process incoming HTTP requests
- cross-platform: runs on multiple operating systems, including windows, linux, macOS

- `npm` (node package manager): default package manager for node
  - manages dependencies with `package.json` and `package-lock.json`
  - `package.json`: lists project metadata and dependencies
  - `package-lock.json`: locks exact versions of packages to ensure consistent installs

- javascript in the browser vs node.js:
  - javascript is the same language everywhere, but what's available depends on the runtime
  - browser gives browser-related tools
  - node.js gives backend/server-related tools

| Feature               | **Browser**                               | **Node.js**                                |
|-----------------------|-------------------------------------------|--------------------------------------------|
| **Purpose**           | Run JavaScript in the browser (frontend)  | Run JavaScript on the server (backend)     |
| **Environment**       | Client-side                               | Server-side                                |
| **JavaScript Engine** | Uses engines like **V8** (Chrome), **SpiderMonkey** (Firefox) | Uses **V8** (same as Chrome) |
| **APIs available**    | Web APIs (e.g., `DOM`, `fetch`, `localStorage`, `alert`) | Node APIs (e.g., `fs`, `http`, `process`) |
| **Access to DOM**     | ✅ Yes                                    | ❌ No (there is no DOM)                     |
| **File System Access**| ❌ No                                     | ✅ Yes (`fs` module)                        |
| **Modules**           | Uses ES Modules (`import/export`) or `<script>` | Uses CommonJS (`require`) or ES Modules    |
| **Async operations**  | Via browser APIs like `fetch`, `setTimeout` | Via Node APIs like `fs.readFile`, `http`  |
| **Usage Example**     | Interactive websites                      | Web servers, APIs, CLI tools                |

code that only works on the **browser**:

```javascript
document.body.style.background = "blue";
alert("Hello from the browser!");
```

- the code above works only in the browser because:
  - `document` is part of the DOM API
  - `alert()` is part of the window API
- in node.js, this crashes because document and alert don’t exist

code that only works in **node.js**:

```javascript
const fs = require("fs");

fs.readFile("file.txt", "utf8", (err, data) => {
  console.log(data);
});
```

- the code above only in node because:
  - `fs` (file system module) is a node API
  - `require()` is node's module system
- in the browser this crashes because there’s no access to your computer’s file system, for security

## module systems: CommonJS vs ES Modules

> lets you import code from one file and export code to another

> way to split and reuse javascript code into separate files

- module: javascript file that exports something or is imported in another file

- CommonJS (old, not recommended)
  - introduced in node.js
  - modules are loaded with `require()` and exported with `module.exports`
  - loads modules synchronously (one after the other, like blocking)

```javascript
// math.js
const add = (a, b) => a + b;
module.exports = { add };

// app.js
const math = require('./math');
console.log(math.add(1, 2));
```

- ES Modules (recommended): official standard in modern javascript
  - introduced in ECMAScript 2015 (ES6)
  - uses `import` and `export`
  - loads modules asynchronously (non-blocking)

```javascript
// math.js
export const add = (a, b) => a + b;

// app.js
import { add } from './math.js';
console.log(add(1, 2));
```

### types of modules

- local modules: modules that we create in our application
- built-in modules: modules that node.js ships with out of the box
- third-party modules: modules written by other developers that we can use in our application

- module wrapper function: every module is wrapped in a function behind the scenes before it's executed

```javascript
(function (exports, require, module, __filename, __dirname) {
  // Your module code here
});
```

- `exports`: object used to export module contents
- `require`: function to import other modules
- `module`: object representing the current module
- `__filename`: full path to the current file
- `__dirname`: directory path of the current module

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

- run javascript code: `node code.js`
- initialized new node.js project: `npm init -y`
  - creates `package.json` file that stores metadata about the project (name, version, dependencies, etc)
  - `-y`: accepts all defaults, skipping interactive prompts
- install all dependencies listed in `package.json`: `npm install`
- install a package locally to current project directory: `npm install express`
  - package reference will be added to `package.json` if it exists
  - package won't be available as a system-wide cli command
  - uninstall local package: `npm uninstall express`
    - removes package from `/node_modules` and `package.json`
  - install a package as development dependency: `npm install jest --save-dev`
- install package globally on your system: `npm install -g typescript`
  - package will be placed in a global location (e.g. `/usr/lib/node_modules` or `~/.volta/tools/...`)
  - package will be accessible anywhere in the terminal
  - doesn't affect `package.json`
  - uninstall global package: `npm uninstall -g typescript`
- updates packages to their latest versions, according to `package.json`: `npm update`
- list all installed packages and their dependencies: `npm list`
- list node version: `node -v`
- list npm version: `npm -v`
- executes package directly from npm registry without installing it globally: `npx create-react-app my-app`
  - useful for trying out tools or running one-off commands

- `NODE_ENV`: special environment variable that:
  - is not set in your `.env` file
  - Is automatically set by Next.js/Node.js
  - indicates which environment your application is running in
    - has three common values: 'development', 'production', or 'test'
    - the value dependes on the command you run
      - `npm run dev` => sets `NODE_ENV='development'`
      - `npm run build` or `npm run start` => sets `NODE_ENV='production'`
      - `npm run test` => sets `NODE_ENV='test'`

## file system module

## event-driven architecture

