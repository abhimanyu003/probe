---
title: "Setup and Teardown"
description: "This article will introduce how to Setup and Teardown."
weight: 5
---

## Overview

In your tests, you may want to run some code before and after each test or file.
In this section, we'll discuss the globally available functions that allow you to do that.

* beforeAll `run before all stages`
* beforeEach `run after each stage`
* afterEach `run after each stage`
* afterAll `run after all stages`

```yaml
name: test setup and teardown

# run before all stages
beforeAll:
  - path/to/first-test.yaml
  - path/to/first-second.yaml

# run after each stage
beforeEach:
  - path/to/first-test.yaml
  - path/to/first-second.yaml

# run after each stage
afterEach:
  - path/to/first-test.yaml
  - path/to/first-second.yaml

# run after all stages
afterAll:
  - path/to/first-test.yaml
  - path/to/first-second.yaml

stages:
  - name: send request
    request:
      url: "https://httpbin.org/get"
      method: GET
    assert:
      status: 200
```
