
## JSX (JavaScript XML)

> allows you to write HTML-like code in your JavaScript files

> syntax extension for JavaScript that looks similar to HTML

JSX rules:
1. Always return a single root element
  - wrap everything in a single parent element
    - e.g.: `<React.Fragment>...</React.Fragment>` or `<>...</>` or `<div>...</div>`
      - `<React.Fragment>...</React.Fragment>`: when you don't want to add an extra div to the DOM
      - `<>...</>`: shorthand for `<React.Fragment>`
        - empty tag is called a fragment
        - allows you to group things without adding extra nodes to the DOM
        - shorter, but you can't use key or other attributes that are available in `<React.Fragment>`
      - `<div>...</div>`: when you need a real HTML wrapper for styling purposes
2. Close all the tags
  - self-closing tags must end with a slash
    -                   HTML            =>             JSX
    - e.g.: `<img src="..." alt="...">` => `<img src="..." alt="..." />`
3. Use camelCase for most HTML attributes

- expressions
  - `{}`: used to embed JavaScript expressions in JSX
    - e.g.: `<h1>{name}</h1>`
  - ternary operator
    - e.g.: `{isTrue ? 'Yes' : 'No'}`
- function components

```jsx
function Car() {
  return <h2>Hi, I am a Car!</h2>;
}
```
# components

> piece of the UI that has its own logic and appearance

```jsx

import React from 'react';

function Hello() {
  return <h1>Hello, World!</h1>;
}

export default Hello;

```

using the previous component:

```jsx

export default function MyApp() {
  return (
    <>
      <h1>Welcome to my app</h1>
      <MyButton />
    </>
  );
}

```

JavaScript vs JSX
- IMPORTANT: the code above is a mix of JavaScript and JSX
  - The function definition `export default function MyApp() {...}` is JavaScript
  - The part inside the parentheses `return (...)` is JSX


## props

- properties that you can pass to a component
- similar to function arguments, that you can send into the component as attributes

- can be used to:
  - pass data from parent to child components
  - pass functions from parent to child components
  - pass children elements

```jsx

function Car(props) {
  return <h2>I am a {props.color} Car!</h2>;
}

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(<Car color="red"/>);

```


---

# study plan for react.js

To prepare effectively for a job interview as a beginner React.js developer, here's a step-by-step learning plan:

### 1. Understand the Basics of React
- Resources:
  - Official React documentation ([React Docs](https://reactjs.org/docs/getting-started.html))
  - FreeCodeCamp's React section ([FreeCodeCamp](https://www.freecodecamp.org/learn/front-end-development-libraries/))

### 2. Learn JSX and Components
- Activities:
  - Create simple components
  - Understand props and state
- Resources:
  - Codecademy's React course ([Codecademy React](https://www.codecademy.com/learn/react-101))

### 3. State Management and Lifecycle Methods
- Activities:
  - Implement state management
  - Use lifecycle methods in class components
- Resources:
  - React documentation on State and Lifecycle ([React Lifecycle](https://reactjs.org/docs/state-and-lifecycle.html))

### 4. Hooks
- Activities:
  - Convert class components to function components using hooks
  - Use useState, useEffect, and custom hooks
- Resources:
  - "React Hooks" on FreeCodeCamp ([React Hooks](https://www.freecodecamp.org/news/an-introduction-to-react-hooks-12843fcd2fd9/))

### 5. Routing and Context API
- Activities:
  - Implement routing with React Router
  - Use the Context API for global state management
- Resources:
  - React Router documentation ([React Router](https://reactrouter.com/))
  - React Context API tutorial ([Context API](https://reactjs.org/docs/context.html))

### 6. Building and Deploying a Project
- Activities:
  - Create a small project (e.g., a todo list or a blog)
  - Deploy your project using services like Netlify or Vercel
- Resources:
  - Tutorial on React project deployment ([Netlify](https://www.netlify.com/blog/2016/07/20/introducing-deploy-previews-in-netlify/))

### 7. Mock Interviews and Practice
- Activities:
  - Practice common interview questions on platforms like LeetCode or Interviewing.io
  - Participate in mock interviews (peers or platforms like Pramp)
- Resources:
  - Interview questions for React developers ([Interview Questions](https://www.interviewbit.com/react-interview-questions/))

### Additional Tips
- Engage with Communities: Join forums or groups like Stack Overflow, Reddit’s r/reactjs, or Discord servers for real-time help and learning.
- Keep a Study Journal: Write down what you learn each day to solidify knowledge and track progress.

With a dedicated plan and consistent effort, you should be able to build a foundational understanding of React and prepare effectively for your job interview. Good luck, Arthur!

---

- DOM (Document Object Model): representation of the HTML elements of a webpage as a tree of nodes and objects
