---
title: "Set Allow Insecure"
description: "This article will introduce how to set allow insecure."
---

Many time in development env, there are no valid ssl crets.
Enable send https without verifying the server's certificates (disabled by default).

```yaml
request:
  allowInsecure: true
```

## Example

```yaml
name: Set Allow Insecure
stages:
  - name: this will allow insecure request
    request:
      url: "https://httpbin.org/get"
      method: GET
      allowInsecure: true
```
