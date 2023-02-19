---
title: "Set Json"
description: "This article will introduce how to Set Json."
---

You can set `json` using `bodyJson` tag inside of request.

This will auto add header

```
Content-Type: application/json; charset=utf-8
```

```yaml
request:
  bodyJson: |
    {
      "id":1
      "name": "test-name"
    }
```

## Example

```yaml
name: Set Json
stages:
  - name: request with body
    request:
      url: "https://httpbin.org/post"
      method: POST
      bodyJson: |
        {
          "id":1
          "name": "test-name"
        }
```
