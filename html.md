
# HTML

> defines the structure and content of web pages

## basic concepts

- HTML = HyperText Markup Language
- describes the structure of a webpage
- HTML5 is the most recent version
  - compatible with previous versions
  - new features:
    - semantic elements: convey meaning to browsers and search engines
      - improve SEO (Seach Engine Optimization) because:
        - crawlers can understand what each part of the page represents
          - crawler: programs used by search engines to automatically browse and index web pages
          - better accessibility for screen readers
      - e.g. `<header></header>`, `<footer></footer>`, `<article></article>`, `<nav></nav>`
    - audio and video support
    - canvas element
      - `<element></element>`
    - offline storage
    - improvements for form controls

- html element: starting tag (e.g. `<p>`) + content + closing tag (e.g. `</p>`)

## basic structure

```html
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>Page Title</title>
    </head>
    <body>
        <h1>My First Heading</h1>
        <p>My first paragraph.</p>
    </body>
</html>
```

- `<!DOCTYPE html>`: document type declaration that tells the browser that this is an HTML5 document
- `<html>`: root element of an HTML page
- `<head>`: contains metadata about the HTML document, e.g. title, character encoding, etc
  - `<meta>`: provides metadata about HTML document
    - `<meta name="viewport" content="width=device-width, initial-scale=1.0">`
- `<title>`: specifies a title for the HTML document, which is shown in the browser's title bar or tab
- `<body>`: contains the content of the document, such as text, images, hyperlinks, tables, lists, etc
- `<h1>`: heading
- `<p>`: paragraph, defines a block of text

## elements

- block elements: start a new line, take up 100% of the width of their container by default
  - `<br>` (self-closing): Inserts a single line break
  - `<a>`: defines the URL a link goes to in an `<a>` tag
    - e.g. `<a href="https://example.com">Visit Site</a>`
    - OBS: `<a>` => Anchor, `href` => Hypertext REFerence
  - `<p>`: defines a paragraph
  - `<div>`: defines a section in a document
  - `<span>`: wrap small chunks of text for styling
  - `<img src="image.jpg" alt="A picture">` (self-closing): defines an image
    - `src`: define image path
    - `alt`: shows when image fails to load and when screen reader is used (accessibility)
  - `<ul>`: defines an unordered list
  - `<ol>`: defines an ordered list
  - `<li>`: defines a list item
  - `<input>`: defines an input control
  - `<table>`: defines a table
  - `<tr>` stands for "table row". Defines a row in the table
  - `<th>` stands for "table header". Defines a header cell in the table
  - `<td>` stands for "table data". Defines a cell in the table
- inline elements: do not start a new line
  - `<strong>`: bold text
  - `<em>`: italic text
  - `<br>`: line break
  - `<button>`: Defines a clickable button
- semantic elements
  - `<header>`: Represents a container for introductory content or a set of navigational links
  - `<nav>`: Defines a set of navigation links
  - `<section>`: Defines a section in a document, such as chapters, tabs, etc
  - `<footer>`: Represents a container for the footer of a document or a section
- other elements
  - `<script>`: embed javascript code within html document

unordered list example:

```html
<ol>
  <li>Wake up</li>
  <li>Brush teeth</li>
  <li>Have breakfast</li>
</ol>
```

ordered list example:

```html
<ul>
  <li>Apples</li>
  <li>Oranges</li>
  <li>Bananas</li>
</ul>
```

form example:

```html
<form action="/submit_form" method="post">
  <label for="fname">First name:</label><br>
  <input type="text" id="fname" name="fname"><br>
  <label for="lname">Last name:</label><br>
  <input type="text" id="lname" name="lname"><br><br>
  <input type="submit" value="Submit">
</form>
```

- `<form>`: Defines an HTML form for user input
  - `action` attribute of the `<form>` tag specifies the URL where the form data should be sent when the form is submitted
  - `method` attribute specifies the HTTP method (get or post) to be used when sending the form data
- `<label>` tag defines a label for an `<input>` element
- `<input>` tag is used to create input controls
  - `type` attribute determines what kind of input control to create (text field, checkbox, radio button, submit button, etc.)
  - `id` attribute is used to uniquely identify the input element for styling and scripting purposes
  - `name` attribute is used to reference the form data after the form is submitted
- `<input type="submit">` creates a submit button
  - when the user clicks this button, the form data is sent to the URL specified in the `action` attribute of the `<form>` tag

```html
<table>
    <tr>
        <th>Firstname</th>
        <th>Lastname</th>
        <th>Email</th>
    </tr>
    <tr>
        <td>John</td>
        <td>Doe</td>
        <td>john@example.com</td>
    </tr>
    <tr>
        <td>Jane</td>
        <td>Doe</td>
        <td>jane@example.com</td>
    </tr>
</table>
```

### inline elements

> elements that don't occupy full width and also don't start a newline

- `<a>`
- `<img>`
- `<strong>`
- `<em>`
- `<button>`

## attributes

- global
  - `id`: unique identifier for an element
  - `class`: class name for css/js targeting
  - `style`: inline css styling
- `<input>`
  - `type`: input type (`text`, `email`, `checkbox`, etc)
  - `name`: name of the input
  - `value`: default input value
  - `placeholder`: hint text inside the input
  - `checked`: sets default state for checkboxes/radios
  - `required`: makes the field mandatory

```html
<a href="https://www.example.com">Visit Example</a>
<div id="header">Welcome</div>
<div class="content">Hello World</div>
<!-- <tagname style="property:value;"> -->
<p style="color: blue; font-size: 18px;">Styled Text</p>
<img src="image.jpg" alt="Description of image">
<img src="https://media.giphy.com/media/l0HlBO7eyXzSZkJri/giphy.gif" alt="What time is it?">
```

## script

> embed javascript code within html document

- `getElementById("myid")`: select and manipulate HTML elements

```html
<!DOCTYPE html>
<html>
<body>
  <p id="demo">Hello</p>

  <script>
    document.getElementById("demo").textContent = "Changed with JavaScript!";
  </script>
</body>
</html>
```

if you have the following tag: `<p id="myParagraph">Hello, World!</p>`

you can print the contents with `getElementById`:

```javascript
var paragraph = document.getElementById("myParagraph");
console.log(paragraph.textContent); // Output: Hello, World!
```

- changing content: `paragraph.textContent = "Hi there!;`
- changing css:
  - `paragraph.style.color = "blue";`
  - `paragraph.style.fontSize = "20px";`

---

- `<iframe>`: embed another document within the current HTML document
  - it creates an inline frame, which can contain an entire webpage within it
  - `<iframe src="https://www.example.com" title="Description of the iframe content"></iframe>`

