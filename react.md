
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
  - `useContext`: used to access context (shared state) in a functional component without passing props manually at every level

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

> allow the use of state and side effects in functional components

### useState

```tsx
import { useState } from "react";

function SimpleForm() {
  const [formData, setFormData] = useState({
    name: "",
    age: "",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value, // Updates only the changed field
    });
  };

  return (
    <div>
      <input
        type="text"
        name="name"
        placeholder="Your Name"
        onChange={handleChange}
      />

      <input
        type="number"
        name="age"
        placeholder="Your Age"
        onChange={handleChange}
      />

      <p>Name: {formData.name}</p>
      <p>Age: {formData.age}</p>
    </div>
  );
}

export default SimpleForm;
```

- `useState`: hook to add state to functional components
  - can accept any type of value as initial state (number, string, boolean, array, object)
    - uses an object when grouping related data together
  - returns an array with two elements: the state and the function to update the state
  - the initial state is passed as an argument to `useState`
- `formData`: object that holds the form data
- `setFormData`: function to update the form data
  - it takes the new state as an argument
  - it can be a new object or a function that returns a new object
- `...formData`: spread operator to copy the existing form data
  - it ensures that the existing data is not lost when updating the state
- `name`: used to identify the input field when the form is submitted
  - the value will be sent to the server as `name=value` pair
- `[e.target.name]`: computed property name to update the specific field
  - it uses the `name` attribute of the input field to update the corresponding field in the state
- `e.target.value`: value of the input field that triggered the change event
  - it updates the value of the specific field in the state
- `onChange`: triggered when the value of an input element changes
  - `React.ChangeEvent<HTMLInputElement>`: type of `event` triggered by a change in the value of an input element

### useEffect

> enables side effects in functional components

- side effects:
  - data fetching
  - updating the DOM
  - subscribing to events
  - setting up timers
- runs after every render by default

```tsx
import { useState, useEffect } from "react";

function App() {
  const [count, setCount] = useState(0);

  useEffect(() => {
    console.log("Component re-rendered! Count:", count);
  });

  return (
    <div>
      <p>Count: {count}</p>
      <button onClick={() => setCount(count + 1)}>Increment</button>
    </div>
  );
}
```

### useContext

> access context (shared state) in a functional component without passing props manually at every level

intead of passing, for example, user authentication data through props, you can use `useContext` to access it directly

without `useContext`:

```jsx
<App user={user}>
  <Navbar user={user} />
  <Dashboard user={user}>
    <Profile user={user} />
    <Settings user={user} />
  </Dashboard>
</App>
```

with `useContext`:

```jsx
const AuthContext = createContext();

function App() {
  const user = { name: 'Alice' };
  return (
    <AuthContext.Provider value={user}>
      <Navbar />
      <Dashboard />
    </AuthContext.Provider>
  );
}

function Navbar() {
  const user = useContext(AuthContext);
  return <div>Welcome, {user.name}!</div>;
}

function Dashboard() {
  return (
    <>
      <Profile />
      <Settings />
    </>
  );
}

function Profile() {
  const user = useContext(AuthContext);
  return <div>Profile of {user.name}</div>;
}
```

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

