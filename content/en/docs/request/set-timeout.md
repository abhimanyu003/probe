---
title: "Set Timeout"
description: "This article will introduce how to Set Timeout."
---

There are some api's that take long to response, you can specify timeout to
each request to wait for response.

```yaml
request:
  timeout: 2000 # in milliseconds this is = 2 second
```

## Example

```yaml
name: Set Timeout
stages:
  - name: this will Set Timeout to Request
    request:
      url: "https://httpbin.org/get"
      method: GET
      timeout: 2000
```
