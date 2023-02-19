---
title: "Set Sleep"
description: "This article will introduce how to set basic auth."
---

You can sleep after each request in `milliseconds`

```yaml
request:
  sleepAfter: 2000 # in milliseconds this is = 2 second
```

## sleepAfter stage

 `sleepAfter` will sleep after stage is completed


```yaml
name: Set Headers
stages:
  - name: Request with Set Sleep
    sleepAfter: 2000 # sleep after each stage in milliseconds
```

Note: if you are doing multiple request using `times`, then sleep will be applied to each request

```yaml
request:
  times: 10 # sleep will be applied to after each request
  sleepAfter: 2000 # in milliseconds this is = 2 second
```

## Example

**sleepAfter example**

```yaml
name: Set Sleep
stages:
  - name: Request with Set Sleep
    request:
      url: "https://httpbin.org/get"
      method: GET
      sleepAfter: 2000
```

**sleepAfter stage is completed example**


```yaml
name: Set Sleep
stages:
  - name: Request with Set Sleep
    sleepAfter: 2000
    request:
      url: "https://httpbin.org/get"
      method: GET
```

**Sleep after each request for 10 times example**

> This will sleep for 2000 millisecond after each request

```yaml
name: Set Sleep
stages:
  - name: Request with Set Sleep
    request:
      url: "https://httpbin.org/get"
      method: GET
      times: 10
      sleepAfter: 2000
```
