---
title: "Set Bearer Auth Token"
description: "This article will introduce how to Set Bearer Auth Token."
---

You can set custom `Set Bearer Auth Token` inside of request.

```yaml
request:
  headers:
    bearerAuthToken: token
```

> This will set `Authorization: Bearer NGU1ZWYwZDJhNmZhZmJhODhmMjQ3ZDc4`

## Example

```yaml
name: Set Bearer Auth Token
stages:
  - name: request with bearer auth token
    request:
      url: "https://httpbin.org/post"
      method: POST
      headers:
        accept: application/json
        bearerAuthToken: ASDF1234567890
```
