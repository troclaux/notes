
# basic concepts

- HTML = HyperText Markup Language
- describes the structure of a webpage
- HTML5 is the most recent version
  - compatible with previous versions
  - new features:
    - semantic elements
      - `<header></header>`
      - `<footer></footer>`
      - `<article></article>`
    - audio and video support
    - canvas element
      - `<element></element>`
    - offline storage
    - improvements for form controls


# basic structure

```html
<!DOCTYPE html>
<html>
    <head>
        <title>Page Title</title>
    </head>
    <body>
        <h1>My First Heading</h1>
        <p>My first paragraph.</p>
    </body>
</html>
```

Explanation:

- `<!DOCTYPE html>`: This is the document type declaration and it tells the browser that this is an HTML5 document
- `<html>`: This is the root element of an HTML page
- `<head>`: This contains meta-information about the HTML document, including the title of the document, character encoding, CSS, scripts etc
- `<title>`: This specifies a title for the HTML document, which is shown in the browser's title bar or tab
- `<body>`: This contains the content of the document, such as text, images, hyperlinks, tables, lists, etc
- `<h1>`: This is a heading. HTML headings are defined with the `<h1>` to `<h6>` tags
- `<p>`: This is a paragraph. It defines a block of text

# elements

- block elements
  - `<br>`: Inserts a single line break
  - `<a>`: Defines a hyperlink
    - OBS: `<a>` => Anchor
  - `<p>`: Defines a paragraph
  - `<div>`: Defines a section in a document
  - `<img>`: Defines an image
  - `<ul>`: Defines an unordered list
  - `<ol>`: Defines an ordered list
  - `<li>`: Defines a list item
  - `<input>`: Defines an input control
  - `<table>`: Defines a table
  - `<tr>` stands for "table row". Defines a row in the table
  - `<th>` stands for "table header". Defines a header cell in the table
  - `<td>` stands for "table data". Defines a cell in the table
- inline elements
  - `<strong>`: bold text
  - `<em>`: italic text
  - `<br>`: line break
  - `<button>`: Defines a clickable button
- form elements
  - `<form>`: Defines an HTML form for user input
- semantic elements
  - `<header>`: Represents a container for introductory content or a set of navigational links
  - `<nav>`: Defines a set of navigation links
  - `<section>`: Defines a section in a document, such as chapters, tabs, etc
  - `<footer>`: Represents a container for the footer of a document or a section
- other elements
  - `<script>`: embed javascript code within html document


```html
<form action="/submit_form" method="post">
  <label for="fname">First name:</label><br>
  <input type="text" id="fname" name="fname"><br>
  <label for="lname">Last name:</label><br>
  <input type="text" id="lname" name="lname"><br><br>
  <input type="submit" value="Submit">
</form>
```

In this example:
- `action` attribute of the `<form>` tag specifies the URL where the form data should be sent when the form is submitted
- `method` attribute specifies the HTTP method (get or post) to be used when sending the form data
- `<label>` tag defines a label for an `<input>` element
- `<input>` tag is used to create input controls
  - `type` attribute determines what kind of input control to create (text field, checkbox, radio button, submit button, etc.)
  - `id` attribute is used to uniquely identify the input element for styling and scripting purposes
  - `name` attribute is used to reference the form data after the form is submitted
- `<input type="submit">` creates a submit button
  - When the user clicks this button, the form data is sent to the URL specified in the `action` attribute of the `<form>` tag

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

## inline elements

> elements that don't occupy full width and also don't start a newline

- `<a>`
- `<img>`
- `<strong>`
- `<em>`
- `<button>`

# attributes

```html
<a href="https://www.example.com">Visit Example</a>
<div id="header">Welcome</div>
<div class="content">Hello World</div>
<!-- <tagname style="property:value;"> -->
<p style="color: blue; font-size: 18px;">Styled Text</p>
<img src="image.jpg" alt="Description of image">
<img src="https://media.giphy.com/media/l0HlBO7eyXzSZkJri/giphy.gif" alt="What time is it?">
```

---

- `<iframe>`: embed another document within the current HTML document
  - It creates an inline frame, which can contain an entire webpage within it

```html
<iframe src="https://www.example.com" title="Description of the iframe content"></iframe>
```

- `<meta>`: provides metadata about HTML document

```html
<meta charset="UTF-8">
<meta name="description" content="A brief description of the webpage">
```

- `getElementById`: select and manipulate HTML elements

if you have the following tag:
```html
<p id="myParagraph">Hello, World!</p>
```

you can print the contents with `getElementById`:
```javascript
var paragraph = document.getElementById("myParagraph");
console.log(paragraph.textContent); // Output: Hello, World!
```

