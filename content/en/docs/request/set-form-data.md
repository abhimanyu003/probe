---
title: "Set Form Data"
description: "This article will introduce how to set form data."
---

> Note: the form data of GET, HEAD, and OPTIONS requests will be ignored by default.

```yaml
request:
  formData:
    key: value
```

You can set multiple values

```yaml
request:
  formData:
    name: Abhimanyu
    username: abhimanyu003
```

## Example

```yaml
name: Set Form Data
stages:
  - name: request with form data
    request:
      url: "https://httpbin.org/post"
      method: POST
      formData:
        name: abhimanyu003
        url: https://github.com/abhimanyu003
```
