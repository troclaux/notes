
- CSS = Cascading Style Sheets


syntax:

```
selector {
    property: value;
}
```

example:

```css
body {
    background-color: lightblue;
}

h1 {
    color: white;
    text-align: center;
}

p {
    font-family: verdana;
    font-size: 20px;
}
```
## selector

types:
- element selector
  - `p { color: red; }`
- class selector
  - `.myClass { color: red; }`
- id selector
  - `#myID { color: red; }`
- attribute selector
  - `a[target="_blank"] { color: red; }`
    - will apply the style to all `<a>` elements with `target="_blank"`
- global selector
  - selects all elements

