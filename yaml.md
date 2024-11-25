
# YAML

- key-value pairs
  - quotes not required (unlike json)
- normally 2 spaces for indentation
- comments begin with `#`
- strings with special characters should be enclosed in quotes (`'` or `"`)

## lists

```yaml
fruits:
  - Apple
  - Banana
  - Cherry
```

## multiline strings

- use either `|` or `>`
  - `|`: preserves newlines
  - `>`: folding (newlines are replaced by spaces)

```yaml
description: |
  This is a multiline string.
  It will keep newlines intact.

description: >
  This is a folded string.
  It will replace newlines with spaces.
```

## yaml vs json

```yaml
employees:
  - name: John
    position: Manager
    department: HR
  - name: Sarah
    position: Developer
    department: IT
```

```json
{
  "employees": [
    {
      "name": "John",
      "position": "Manager",
      "department": "HR"
    },
    {
      "name": "Sarah",
      "position": "Developer",
      "department": "IT"
    }
  ]
}
```
