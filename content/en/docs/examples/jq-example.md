---
title: "JQ example"
description: "JQ example"
date: 2020-11-12T13:26:54+01:00
lastmod: 2020-11-12T13:26:54+01:00
draft: false
toc: true
weight: 200
---

Simple GET request that will show various JQ operations.

```yaml
name: JQ Example
stages:
  - name: get products request
    request:
      url: "https://dummyjson.com/products"
      method: GET
    assert:
      status: 200
      body:
        - select: .products | length
          want: 50

        - select: .products[0].id | isnormal
          want: true

        - select: .products[0].images | length
          want: 5

        - select: .products[] | select(.title == "iPhone 9") | .id
          want: 1
```
