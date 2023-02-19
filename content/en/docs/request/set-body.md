---
title: "Set Body"
description: "This article will introduce how to set body."
---

You can set raw `body` using `body` tag inside of request.

Default content-type header is `text/plain; charset=utf-8`

```yaml
request:
  body: |
    {
      "id":1
      "name": "test-name"
    }
```

## Example

```yaml
name: Set Body
stages:
  - name: request with body
    request:
      url: "https://httpbin.org/post"
      method: POST
      body: |
        {
          "id":1
          "name": "test-name"
        }
      headers:
        content-type: application/json
```

To set the correct content type make to to update the headers as well.

**For JSON**

```yaml
      headers:
        content-type: application/json
```

**For XML**

```yaml
      headers:
        content-type: text/xml; charset=utf-8
```
