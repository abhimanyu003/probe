---
title: "Upload"
description: "This article will introduce how to Upload."
---

You can send request using basic auth as well, basic auth need two things.
`username` and `password`

```yaml
request:
  url: "https://example.com/upload"
  method: POST
  files:
    - name: pic.jpg
      path: "/full/path/to/image/150.png"
    - name: pic-2.jpg
      path: "/full/path/to/image/150.png"
```

## Example

```yaml
name: Basic Auth
stages:
  - name: file upload
    request:
      url: "https://example.com/upload"
      method: POST
      files:
        - name: pic.jpg
          path: "/full/path/to/image/150.png"
        - name: pic-2.jpg
          path: "/full/path/to/image/150.png"
```
