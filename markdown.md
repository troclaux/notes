
# markdown

[markdown best practices](https://learn.microsoft.com/powershell/scripting/community/contributing/general-markdown)
[markdown syntax](https://www.markdownguide.org/basic-syntax/)
[markdown supported languages](https://github.com/highlightjs/highlight.js/blob/main/SUPPORTED_LANGUAGES.md)

## basic syntax

- `#` for headers
  - only have 1 H1 per markdown file
  - limit depth to H3 or H4
- `-` for unordered lists
  - lists should be surrounded by 1 blank line
  - don't end list items with a period
- `1.` for numbered lists
  - all items in a list should use the number 1. rather than incrementing each item
    - markdown will automatically number the items
    - this makes it easier to reorder items
- `---` for horizontal rules
- `| header1 | header2 |` for tables
- `| --- | --- |` for table headers
- `| data1 | data2 |` for table data
- `_` single underscore for italic
- `**` for bold
- `***` for bold and italic
- `>` for blockquotes
- [link text](url) for links
- ![alt text](url) for images
- `inline code` for inline code (also called code span)
- ```code block``` for code blocks
  - add the language after the first triple backticks for syntax highlighting
    - you can use `output` tag
- `[<description>][<reference>]` for reference links
- [reference]: www.url-for-reference-links.com "title for reference links"
- reference a header using cross-reference links (also called anchors)
  - format: `[label](/filepath.md#heading-name)`
    - replace consecutive non-letters from heading with `-`
    - only use 1 `#` character between full path and the heading
      - examples:
        - `#### JSON (JavaScript Object Notation)` => `[JSON](/javascript.md#json-javascript-object-notation)`
        - [basic syntax](#basic-syntax)
        - [git anchor](git.md#git-bisect)

## best practices

- there should be a single blank line between markdown blocks of different types
  - e.g. between a list and a code block
  - don't use more than 1 blank line

- remove trailing whitespace
- use spaces instead of tabs

## escaping characters

characters that can be escaped with a backslash:
- backslash `\`
- backtick
  - also escape backtick by wrapping it in double backticks
    - e.g. ``sentence that uses a backtick (`) that should be escaped``
- `*` asterisk
- `_` underscore
- `#` hash
- `+` plus
- `-` minus
- `.` dot
- `!` exclamation mark
- `|` pipe
- `{}` curly braces
- `[]` brackets
- `()` parentheses
- `<>` angle brackets

## alert boxes

```markdown
> [!NOTE]
> Information the user should notice even if skimming.

> [!TIP]
> Optional information to help a user be more successful.

> [!IMPORTANT]
> Essential information required for user success.

> [!CAUTION]
> Negative potential consequences of an action.

> [!WARNING]
> Dangerous certain consequences of an action.
```
