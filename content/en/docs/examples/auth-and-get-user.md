---
title: "Auth and Get User"
description: "Auth and Get User."
date: 2020-10-06T08:49:31+00:00
lastmod: 2020-10-06T08:49:31+00:00
draft: false
weight: 100
---


> Based on
> * Auth [https://dummyjson.com/docs/auth](https://dummyjson.com/docs/auth)
> * Single User [https://dummyjson.com/docs/users](https://dummyjson.com/docs/users)

```yaml
name: auth and get user
stages:
  - name: auth
    request:
      url: "https://dummyjson.com/auth/login"
      method: POST
      formData:
        username: kminchelle
        password: 0lelplR
    assert:
      status: 200
    export:
      body:
        - select: .id
          as: userId # setting id as userId variable
        - select: .token
          as: token # setting token id as variable

  - name: Get Single User
    request:
      # using userId variable from previous stage
      url: "https://dummyjson.com/users/${userId}"
      method: GET
      headers:
        # using token from previous stage
        token: ${token}
    assert:
      status: 200
      body:
        - select: .username
          want: "kminchelle"
        - select: .id
          want: 15
```
