---
title: "Form Data"
description: "Form Data"
date: 2020-11-12T13:26:54+01:00
lastmod: 2020-11-12T13:26:54+01:00
draft: false
toc: true
weight: 200
---

Simple GET request that will show various JQ operations.

```yaml
name: Form Data
stages:
  - name: add product
    request:
      url: "https://dummyjson.com/products/add"
      method: POST
      headers:
        Content-Type: 'application/json'
      formData:
        title: 'BMW Pencil'
    assert:
      status: 200
      body:
        - select: .id
          want: 101
```
