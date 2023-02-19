---
title: "Set User-Agent"
description: "This article will introduce how to set User-Agent."
---

Set the "User-Agent" header for requests.

Default value is `probe ( https://github.com/abhimanyu003/probe )`

```yaml
request:
  userAgent: my-custom-user-agent
```

## Example

```yaml
name: Set User Agent
stages:
  - name: request with custom user agent
    request:
      url: "https://httpbin.org/get"
      method: GET
      userAgent: "AppleTV11,1/11.1"
```
