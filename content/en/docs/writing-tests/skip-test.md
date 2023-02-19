---
title: "Skip Test"
description: "This article will introduce how to Skip Test."
weight: 6
---

## Overview

During development, you may want to temporarily turn off a test.
Rather than commenting it out, you can use the skip method.

It's possible to skip a test as well as individual stage

```yaml
name: setting test variables
skip: true # skip whole test
stages:
  - name: send request
    skip: true # skip stage
    request:
      url: "https://httpbin.org/get"
      method: GET
    assert:
      status: 200
```
