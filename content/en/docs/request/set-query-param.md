---
title: "Set Query Param"
description: "This article will introduce how to set form data."
---

Set a URL query parameter with a key-value pair for requests.

```yaml
request:
  queryParams:
    key: value
```

You can set multiple values

```yaml
request:
  queryParams:
    name: Abhimanyu
    username: abhimanyu003
```

## Example

```yaml
name: Set Query Param
stages:
  - name: request with query param
    request:
      url: "https://httpbin.org/post"
      method: POST
      queryParams:
        name: Abhimanyu
        username: abhimanyu003
```
