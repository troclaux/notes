
# react.js

> javascript library for building user interfaces through reusable components

## basic concepts

- **reactive** nature: ability to efficiently update and render user interfaces in response to changes in data
- component: basic building block that can be reused
- state: data that can change and trigger re-renders
- props: data passed from parent to child components
- virtual DOM: javascript object that mirrors the structure of the real dom
  - improves performance by minimizing expensive dom operations
  - when the state of a component changes, react creates a new virtual dom tree
- hooks: enables features like state and side effects
  - `useState`: add state to functional components
  - `useEffect`: handle side effects (like fetching data, setting timers or updating the dom)
  - `useContext`: used to access context (shared state) in a functional component

## JSX (JavaScript XML)

> syntax extension for JavaScript that looks similar to HTML

- allows you to write HTML-like code in your JavaScript files

### JSX rules

- always return a single root element
  - wrap everything in a single parent element
    - e.g.: `<React.Fragment>...</React.Fragment>` or `<>...</>` or `<div>...</div>`
      - `<React.Fragment>...</React.Fragment>`: when you don't want to add an extra div to the DOM
      - `<>...</>`: shorthand for `<React.Fragment>`
        - empty tag is called a fragment
        - allows you to group things without adding extra nodes to the DOM
        - shorter, but you can't use key or other attributes that are available in `<React.Fragment>`
      - `<div>...</div>`: when you need a real HTML wrapper for styling purposes
- close all the tags
  - self-closing tags (e.g. `<br>`, `<input>`) must end with a slash
    - e.g.: `<img src="..." alt="...">` (HTML) => `<img src="..." alt="..." />` (JSX)
- use camelCase for most HTML attributes

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
## components

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

> [!IMPORTANT]
> the code above is a mix of JavaScript and JSX
> The function definition `export default function MyApp() {...}` is JavaScript
> The part inside the parentheses `return (...)` is JSX

### props

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

- `key`: used to identify elements in a list
  - used for rendering a list of items dynamically
  - `{users.map((user) => ( <li key={user.id}>{user.name}</li> ))}`
  - keys should be unique and stable (like an `id`)

## state management
## lifecycle methods
## hooks

## useful props

- `htmlFor` (used in `<label>`):
  - This is a React/JSX version of HTML's `for` attribute used in labels
  - It creates a connection between the label and the input field by matching the input's `id`
  - When a user clicks the label, it automatically focuses the associated input field
  - This improves accessibility and usability, especially for screen readers
- `id` (used in `<input>`): unique identifier for the input element
  - Must match the `htmlFor` value of its associated label
  - Helps establish the label-input relationship for accessibility
  - Should be unique across the entire page
-  `name` (used in `<input>`): used to identify the input field when the form is submitted
  - The value will be sent to the server as `name=value` pair
  - Important for form handling and data processing
  - Multiple inputs can share the same name (like in radio buttons)
- `type` (used in `<input>`): specifies what kind of input field it is
  - in this case, `type="email"` provides:
    - Email format validation
    - Appropriate keyboard on mobile devices
    - Special input handling in browsers
  - common types include: `text`, `password`, `email`, `number`, `checkbox`, `radio`, etc.
- `required` (used in `<input>`): A boolean attribute that makes the field mandatory
  - prevents form submission if the field is empty
  - browsers will show a validation message if you try to submit without filling required fields

Example of how these work together:
```html
<!-- When you click "Email", it focuses the input field -->
<label htmlFor="email">Email</label>
<input
  id="email"           <!-- Connects to label's htmlFor -->
  name="email"         <!-- Will be sent as {email: value} -->
  type="email"         <!-- Ensures valid email format -->
  required            <!-- Must be filled before submission -->
/>
```

When the form is submitted, the data would look something like:

```javascript
{
  email: "user@example.com"  // name becomes the key
}
```

## events

> events in react are similar to DOM events but with camelCase naming and passed as JSX attributes

- `onClick`: triggered when an element is clicked
- `onChange`: triggered when form input value changes
  - `React.ChangeEvent<HTMLInputElement>`: type of `event` triggered by a change in the value of an input element
- `onSubmit`: triggered when a form is submitted
- `onMouseEnter`: triggered when mouse enters element area
- `onMouseLeave`: triggered when mouse leaves element area

- `e.target`: DOM element that triggered the event
  - `e.target.name`
  - `e.target.value`

---

## study plan for react.js

To prepare effectively for a job interview as a beginner React.js developer, here's a step-by-step learning plan:

#### 1. Understand the Basics of React
- Resources:
  - Official React documentation ([React Docs](https://reactjs.org/docs/getting-started.html))
  - FreeCodeCamp's React section ([FreeCodeCamp](https://www.freecodecamp.org/learn/front-end-development-libraries/))

#### 2. Learn JSX and Components
- Activities:
  - Create simple components
  - Understand props and state
- Resources:
  - Codecademy's React course ([Codecademy React](https://www.codecademy.com/learn/react-101))

#### 3. State Management and Lifecycle Methods
- Activities:
  - Implement state management
  - Use lifecycle methods in class components
- Resources:
  - React documentation on State and Lifecycle ([React Lifecycle](https://reactjs.org/docs/state-and-lifecycle.html))

#### 4. Hooks
- Activities:
  - Convert class components to function components using hooks
  - Use useState, useEffect, and custom hooks
- Resources:
  - "React Hooks" on FreeCodeCamp ([React Hooks](https://www.freecodecamp.org/news/an-introduction-to-react-hooks-12843fcd2fd9/))

#### 5. Routing and Context API
- Activities:
  - Implement routing with React Router
  - Use the Context API for global state management
- Resources:
  - React Router documentation ([React Router](https://reactrouter.com/))
  - React Context API tutorial ([Context API](https://reactjs.org/docs/context.html))

#### 6. Building and Deploying a Project
- Activities:
  - Create a small project (e.g., a todo list or a blog)
  - Deploy your project using services like Netlify or Vercel
- Resources:
  - Tutorial on React project deployment ([Netlify](https://www.netlify.com/blog/2016/07/20/introducing-deploy-previews-in-netlify/))

#### 7. Mock Interviews and Practice
- Activities:
  - Practice common interview questions on platforms like LeetCode or Interviewing.io
  - Participate in mock interviews (peers or platforms like Pramp)
- Resources:
  - Interview questions for React developers ([Interview Questions](https://www.interviewbit.com/react-interview-questions/))

#### Additional Tips
- Engage with Communities: Join forums or groups like Stack Overflow, Reddit’s r/reactjs, or Discord servers for real-time help and learning.
- Keep a Study Journal: Write down what you learn each day to solidify knowledge and track progress.

With a dedicated plan and consistent effort, you should be able to build a foundational understanding of React and prepare effectively for your job interview. Good luck, Arthur!

---

- DOM (Document Object Model): representation of the HTML elements of a webpage as a tree of nodes and objects
