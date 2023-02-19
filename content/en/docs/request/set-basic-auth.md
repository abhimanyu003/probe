---
title: "Set Basic Auth"
description: "This article will introduce how to set basic auth."
---

You can send request using basic auth as well, basic auth need two things.
`username` and `password`

```yaml
request:
  basicAuth:
    username: me
    password: mypassword
```

## Example

```yaml
name: Basic Auth
stages:
  - name: request with basic auth
    request:
      url: "https://httpbin.org/get"
      method: GET
      basicAuth:
        username: me
        password: mypassword
```
