---
title: "Set Headers"
description: "This article will introduce how to set headers."
---

You can set custom `headers` inside of request.

```yaml
request:
  headers:
    key: value
```

You can set multiple headers, all of these will be sent in request.

```yaml
request:
  headers:
    accept: application/json
    x-auth-token: 123adsf
    content-type: text/xml; charset=utf-8
```

## Example

```yaml
name: Set Headers
stages:
  - name: request with custom headers
    request:
      url: "https://httpbin.org/post"
      method: POST
      headers:
        accept: application/json
        x-auth-token: custom-token
        content-type: text/xml; charset=utf-8
```
