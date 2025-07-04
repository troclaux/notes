
# CSS (Cascading Style Sheets)

> used to style html documents, it controls the appearance of web pages

- inline css: `<p style="color: red;">Text in red</p>`
- internal css: used inside `<style>` tag
- external css: separate `.css` file linked to html: `<link rel="stylesheet" href="style.css">`

- comment with `/* */`
- viewport: visible part of a web page in the browser window, it's what the users sees without scrolling

syntax of a css property:

```
selector {
  property1: value1;
  property2: value2;
}
```

- declaration: `property1: value1;`
- declaration block: `{ ... }`

example:

```css
h1.highlight {
    color: white;
    text-align: center;
}

p {
    background-color: lightblue;
}

div#profile {
    font-family: verdana;
    font-size: 20px;
}
```

## selectors

- element selector
  - `p { color: red; }`
    - `p` styles all `<p>` tags
- class selector
  - `.myClass { color: red; }`
- id selector
  - `#myID { color: red; }`
- attribute selector
  - `a[target="_blank"] { color: red; }`
    - will apply the style to all `<a>` elements with `target="_blank"`
- global selector (`*`)
  - selects every element in the web page
- select multiple elements: `div, p { ... }`

> [!IMPORTANT]
> if multiple rules apply to an element, the most specific rule wins
> inline > id > class > tag

## common css properties

- `color`: color of text
  - can be specified with `rgb(255, 99, 71)` or hex codes `#ff6347`
    - `background: rgba (0, 128, 0, 0.3)`
  - `border-color`
  - `background-color`
- `background-image`:
- `opacity`:
- `font-size`: size of text
- `font-family`: sets font type
- `font-weight`: normal, bold
- `font-style`: normal, italic
- `margin`: `margin: 25px 50px 75px 100px;` (sets, in order: top margin, right margin, bottom margin and left margin)
  - `margin-top`, `margin-right`, `margin-bottom`, `margin-left`
- `padding`:
- `border`:
  - `border-style`: can be `dotted`, `dashed`, `solid`
  - `border-radius`: round the corner of an element's border box
- `display`: can be `block`, `inline`, `flex`, `grid`
  - `inline-block`: no newline, allows settign width and height
  - `flex`: 
- `position`:
  - `relative`: positions the element relative to its normal position
  - `absolute`: positions the element relative to the nearest positioned ancestor
    - ancestor: any element that wraps or contains another
  - `fixed`: positions the element relative to the viewport, so it stays in place when scrolling
- `width`:
- `height`:
- `float`: position element to the left or right of its container
- `z-index`: element with higher value will be displayed on top when overlapping

### how to center an element vertically and horizontally

#### with flexbox

```html
<div class="container">
  <p>Sample text</p>
</div>
```

```css
.container {
  display: flex;
  justify-content: center; /* horizontal alignment */
  align-items: center;     /* vertical alignment */
}
```

- `flex-direction`: defines main-axis (default is horizontal)
- `justify-content`: positions content in the main axis
- `align-items`: items are aligned in the cross-axis (default is vertical)

#### with grid layout

```css
.parent {
  display: grid;
  place-content: center;
  place-items: center;
}
```

## types of units

> [!TIP]
> Use px when you want precise control (e.g. borders, icons).
> Use rem and em for text, so it scales nicely across devices.
> Use %, vw, and vh for layout that adapts to screen size.

### absolute units (fixed size)

| Unit | Meaning            | Example            |
| ---- | ------------------ | ------------------ |
| `px` | Pixels             | `font-size: 16px;` |
| `pt` | Points (1/72 inch) | `font-size: 12pt;` |
| `cm` | Centimeters        | `width: 10cm;`     |
| `mm` | Millimeters        | `width: 50mm;`     |
| `in` | Inches             | `width: 1in;`      |

### relative units (flexible, better for responsive design)

| Unit   | Meaning                                               | Example              |
| ------ | ----------------------------------------------------- | -------------------- |
| `%`    | Relative to parent element                            | `width: 50%;`        |
| `em`   | Relative to the font-size of the **parent element**   | `font-size: 2em;`    |
| `rem`  | Relative to the **root element** font-size (`<html>`) | `font-size: 1.5rem;` |
| `vw`   | 1% of viewport width                                  | `width: 50vw;`       |
| `vh`   | 1% of viewport height                                 | `height: 100vh;`     |
| `vmin` | 1% of smaller of `vw` or `vh`                         | `font-size: 5vmin;`  |
| `vmax` | 1% of larger of `vw` or `vh`                          | `font-size: 5vmax;`  |

## media queries

> allow css to apply rules based on screen size, resolution or device type

```css
@media (max-width: 600px) {
  body { font-size: 14px; }
}
```

## css box model

> describes how elements are rendered

- each element is made up of four layers, listed below from inner to outer:
  - content in the center, inner layer
  - padding: surrounds the content
  - border
  - margin: outer layer

## pseudo-classes

> keywords added to selectors that specify a special state of the selected elements

- `:hover`: when the user hovers over an element
- `:active`: while the element is being clicked
- `:focus`: when an element (like an input) gains focus

```css
button:hover {
  background-color: blue;
}
```

## pseudo-elements

> style specific parts of elements without needing additional markup

- `::before`: inserts content before an element
- `::after`: inserts content after an element
- `::first-letter`: styles the first letter of a block
- `::first-line`: styles the first line of text

```css
p::first-letter {
  font-size: 200%;
  color: red;
}
```

## specificity hierarchy

> defines priority of application of hierarchy in case of conflict

**HIGHEST PRIORITY**

- inline style: `<h1 style="color: pink;">`
- id selectors: `#navbar`
- classes and pseudo-classes: `.container`
- attributes
- elements and pseudo-elements: `h1`

**LOWEST PRIORITY**

## flexbox

[flexbox basics](https://css-tricks.com/snippets/css/a-guide-to-flexbox)

## gridbox

[gridbox](https://css-tricks.com/snippets/css/complete-guide-grid)

- `display: grid;`: turn a container into a grid

